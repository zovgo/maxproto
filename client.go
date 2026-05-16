package maxproto

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zovgo/maxproto/internal"
	"github.com/zovgo/maxproto/packet"
	"github.com/zovgo/maxproto/protocol"
)

type Client struct {
	conn *websocket.Conn

	writeMu sync.Mutex
	stateMu sync.Mutex

	closed bool

	ctx    context.Context
	cancel context.CancelFunc

	data internal.ValueWithMutex[struct {
		contacts map[int64]protocol.Contact
		chats    map[int64]protocol.Chat
		profile  *protocol.Profile
	}]
	seq atomic.Int64

	wg sync.WaitGroup
}

func (c *Client) Conn() *websocket.Conn {
	return c.conn
}

var ErrClientClosed = errors.New("client closed")

func (c *Client) WritePacket(pk packet.Packet) error {
	if c.Closed() {
		return ErrClientClosed
	}
	c.writeMu.Lock()
	defer c.writeMu.Unlock()

	if c.Closed() {
		return ErrClientClosed
	}
	data, err := pk.MarshalPacket(int(c.seq.Add(1)))
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	c.setContextualDeadline(true)
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Client) ReadPacket() (packet.Packet, error) {
	if c.Closed() {
		return nil, ErrClientClosed
	}
	c.setContextualDeadline(false)

	_, data, err := c.conn.ReadMessage()
	if err != nil {
		if c.Closed() {
			return nil, ErrClientClosed
		}
		return nil, err
	}
	pk, err := packet.UnmarshalResponse(data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return pk, nil
}

func (c *Client) keepAlive() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			if err := c.WritePacket(&packet.KeepAliveRequest{}); err != nil {
				//fixme: log
				return
			}
		}
	}
}

func (c *Client) WaitForMessages(handler func(*packet.ReceiveMessage)) error {
	c.wg.Add(1)
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			if errors.Is(c.ctx.Err(), context.Canceled) {
				return nil
			}
			return c.ctx.Err()
		default:
		}
		pk, err := c.ReadPacket()
		if err != nil {
			if errors.Is(err, packet.ErrUnhandledOpcode) {
				continue
			}
			if errors.Is(err, ErrClientClosed) {
				return nil
			}
			return err
		}
		switch pk := pk.(type) {
		case *packet.ReceiveMessage:
			c.handleMessage(handler, pk)
		case *packet.ProfileChange:
			c.handleProfileChange(pk)
		case *packet.ContactChange:
			c.handleContactChange(pk)
		}
	}
}

func (c *Client) handleContactChange(pk *packet.ContactChange) {
	c.data.Lock()
	defer c.data.Unlock()

	if _, ok := c.data.V.contacts[pk.Contact.ID]; !ok {
		c.data.V.contacts[pk.Contact.ID] = pk.Contact
		return
	}
	c.data.V.contacts[pk.Contact.ID] = pk.Contact
}

func (c *Client) handleProfileChange(pk *packet.ProfileChange) {
	c.data.Lock()
	defer c.data.Unlock()

	if pk.Profile.Contact.ID != c.data.V.profile.Contact.ID {
		return
	}
	*c.data.V.profile = pk.Profile
}

func (c *Client) handleMessage(h func(*packet.ReceiveMessage), pk *packet.ReceiveMessage) {
	defer h(pk)

	if len(pk.Message.Attaches) == 0 {
		return
	}
	a := pk.Message.Attaches[0]
	if a.Type != "CONTROL" || a.Event != "title" {
		return
	}
	c.data.Lock()
	defer c.data.Unlock()

	ch, ok := c.data.V.chats[pk.ChatID]
	if !ok {
		return
	}
	ch.Title = a.Title
	c.data.V.chats[pk.ChatID] = ch
}

func (c *Client) setContextualDeadline(write bool) {
	deadline, ok := c.ctx.Deadline()
	if !ok {
		return
	}
	if write {
		_ = c.conn.SetWriteDeadline(deadline)
		return
	}
	_ = c.conn.SetReadDeadline(deadline)
}

var ErrAlreadyClosed = errors.New("already closed")

func (c *Client) Close() error {
	c.stateMu.Lock()
	if c.closed {
		c.stateMu.Unlock()
		return ErrAlreadyClosed
	}
	c.closed = true
	c.stateMu.Unlock()

	c.cancel()
	_ = c.conn.Close()
	c.wg.Wait()
	return nil
}

func (c *Client) Closed() bool {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	if c.closed {
		return true
	}
	select {
	case <-c.ctx.Done():
		c.closed = true
		_ = c.conn.Close()
		return true
	default:
	}
	return false
}

func (c *Client) Contact(id int64) (protocol.Contact, bool) {
	c.data.Lock()
	defer c.data.Unlock()

	if p := c.data.V.profile; p.Contact.ID == id {
		// self
		return p.Contact, true
	}
	co, ok := c.data.V.contacts[id]
	return co, ok
}

func (c *Client) Contacts() iter.Seq[protocol.Contact] {
	return func(yield func(protocol.Contact) bool) {
		c.data.Lock()
		defer c.data.Unlock()

		for _, co := range c.data.V.contacts {
			if !yield(co) {
				return
			}
		}
	}
}

func (c *Client) Chat(id int64) (protocol.Chat, bool) {
	c.data.Lock()
	defer c.data.Unlock()

	ch, ok := c.data.V.chats[id]
	return ch, ok
}

func (c *Client) Chats() iter.Seq[protocol.Chat] {
	return func(yield func(protocol.Chat) bool) {
		c.data.Lock()
		defer c.data.Unlock()

		for _, ch := range c.data.V.chats {
			if !yield(ch) {
				return
			}
		}
	}
}

func (c *Client) Profile() protocol.Profile {
	c.data.Lock()
	defer c.data.Unlock()
	return *c.data.V.profile
}
