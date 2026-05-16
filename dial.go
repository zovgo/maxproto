package maxproto

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zovgo/maxproto/packet"
	"github.com/zovgo/maxproto/protocol"
)

const Hostname = "wss://ws-api.oneme.ru/websocket"

type DialConfig struct {
	Dialer    *websocket.Dialer
	Token     string
	DeviceID  uuid.UUID
	UserAgent *protocol.UserAgent
	ChatCount int
}

func (conf DialConfig) DialContext(ctx context.Context, parentTemporary bool) (*Client, error) {
	return dial(ctx, parentTemporary, conf)
}

func (conf DialConfig) DialTimeout(timeout time.Duration) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return dial(ctx, true, conf)
}

func (conf DialConfig) Dial() (*Client, error) {
	return conf.DialTimeout(time.Second * 30)
}

var ErrEmptyToken = errors.New("empty token")

var pingPeriod = 20 * time.Second

func dial(parent context.Context, parentTemp bool, conf DialConfig) (*Client, error) {
	if err := parent.Err(); err != nil {
		return nil, err
	}
	if err := fillConf(&conf); err != nil {
		return nil, err
	}
	conn, _, err := conf.Dialer.DialContext(parent, Hostname, NewHeader(conf.UserAgent.Header))
	if err != nil {
		return nil, fmt.Errorf("ws handshake: %w", err)
	}
	conn.SetPongHandler(func(string) error {
		_ = conn.SetReadDeadline(time.Now().Add(pingPeriod * 2))
		return nil
	})
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	if parentTemp {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithCancel(parent)
	}
	cl := &Client{
		conn:   conn,
		ctx:    ctx,
		cancel: cancel,
	}
	if err := login(cl, conf); err != nil {
		_ = cl.Close()
		return nil, fmt.Errorf("login: %w", err)
	}
	cl.wg.Go(func() {
		select {
		case <-parent.Done():
			_ = cl.Close()
		case <-ctx.Done():
		}
	})
	cl.wg.Go(cl.keepAlive)
	return cl, nil
}

func login(cl *Client, conf DialConfig) error {
	if err := handshake(cl, conf); err != nil {
		return fmt.Errorf("send handshake: %w", err)
	}
	if err := loginRequest(cl, conf); err != nil {
		return fmt.Errorf("send login request: %w", err)
	}
	if err := expectBanners(cl); err != nil {
		return fmt.Errorf("expect banners: %w", err)
	}
	if err := chatsGetRequest(cl); err != nil {
		return fmt.Errorf("chats get: %w", err)
	}
	return nil
}

func handshake(cl *Client, conf DialConfig) error {
	if err := cl.WritePacket(&packet.HandshakeRequest{
		UserAgent: conf.UserAgent,
		DeviceID:  conf.DeviceID,
	}); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	pk, err := cl.ReadPacket()
	if err != nil {
		return fmt.Errorf("read resp: %w", err)
	}
	if pk.Opcode() != packet.HandshakeOpcode {
		return fmt.Errorf("unexpected opcode: %v", pk.Opcode())
	}
	return nil
}

func loginRequest(cl *Client, conf DialConfig) error {
	if err := cl.WritePacket(&packet.LoginRequest{
		Token:        conf.Token,
		ChatCount:    conf.ChatCount,
		Interactive:  true,
		ChatSync:     0,
		ContactsSync: -1,
		PresenceSync: -1,
		DraftsSync:   -1,
	}); err != nil {
		return fmt.Errorf("write req: %w", err)
	}
	pk, err := cl.ReadPacket()
	if err != nil {
		return fmt.Errorf("read resp: %w", err)
	}
	if pk.Opcode() != packet.LoginOpcode {
		return fmt.Errorf("unexpected opcode: %v", pk.Opcode())
	}
	initSession(cl, pk.(*packet.LoginResponse))
	return nil
}

func expectBanners(cl *Client) error {
	pk, err := cl.ReadPacket()
	if err != nil {
		return fmt.Errorf("read resp: %w", err)
	}
	if pk.Opcode() != packet.BannersConfigOpcode {
		return fmt.Errorf("unexpected opcode: %v", pk.Opcode())
	}
	return nil
}

func chatsGetRequest(cl *Client) error {
	if err := cl.WritePacket(&packet.ChatsGetRequest{ChatIDs: []int64{0}}); err != nil {
		return fmt.Errorf("write req: %w", err)
	}
	if _, err := cl.ReadPacket(); err != nil {
		return fmt.Errorf("read resp: %w", err)
	}
	return nil
}

func initSession(cl *Client, pk *packet.LoginResponse) {
	cl.contacts = make(map[int64]protocol.Contact) //TODO: max also sends packet if chat,profile,contact changes
	for _, c := range pk.Contacts {
		cl.contacts[c.ID] = c
	}
	cl.chats = make(map[int64]protocol.Chat)
	for _, c := range pk.Chats {
		cl.chats[c.ID] = c
	}
	cl.profile = &pk.Profile
}

func fillConf(conf *DialConfig) error {
	if conf.Token == "" {
		return ErrEmptyToken
	}
	if conf.DeviceID == uuid.Nil {
		conf.DeviceID = uuid.New()
	}
	if conf.UserAgent == nil {
		conf.UserAgent = protocol.DefaultUserAgent
	}
	if conf.Dialer == nil {
		conf.Dialer = defaultDialer()
	}
	if conf.ChatCount <= 0 {
		conf.ChatCount = 100
	}
	return nil
}

func defaultDialer() *websocket.Dialer {
	return &websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  15 * time.Second,
		ReadBufferSize:    1024 * 4,
		WriteBufferSize:   1024 * 4,
		EnableCompression: true,
	}
}

func NewHeader(ua string) http.Header {
	headers := make(http.Header)
	headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	headers.Set("Accept-Language", "en-US,en;q=0.5")
	headers.Set("Cache-Control", "no-cache")
	headers.Set("Origin", "https://web.max.ru")
	headers.Set("Pragma", "no-cache")
	headers.Set("Sec-Fetch-Dest", "empty")
	headers.Set("Sec-Fetch-Mode", "websocket")
	headers.Set("Sec-Fetch-Site", "cross-site")
	headers.Set("Sec-GPC", "1")
	headers.Set("User-Agent", ua)
	return headers
}
