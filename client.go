package maxproto

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"sync"
	"time"

	"github.com/gorilla/websocket"
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

	contacts map[int64]protocol.Contact
	profile  *protocol.Profile

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
	data, err := pk.MarshalPacket()
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

func (c *Client) keepAlive(seq func() int) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			if err := c.WritePacket(&packet.KeepAliveRequest{Seq: seq()}); err != nil {
				//fixme: log
				return
			}
		}
	}
}

func (c *Client) WaitForMessages(handler func(protocol.Message)) error {
	c.wg.Add(1)
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			return c.ctx.Err()
		default:
		}
		pk, err := c.ReadPacket()
		if err != nil {
			if errors.Is(err, packet.ErrUnhandledOpcode) {
				continue
			}
			return err
		}
		if pk.Opcode() != packet.ReceiveMessageOpcode {
			continue
		}
		handler(*pk.(*packet.ReceiveMessage).Message)
	}
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

func (c *Client) GetContact(id int64) (protocol.Contact, bool) {
	if c.profile.Contact.ID == id {
		// self
		return c.profile.Contact, true
	}
	co, ok := c.contacts[id]
	return co, ok
}

func (c *Client) Contacts() iter.Seq[protocol.Contact] {
	return func(yield func(protocol.Contact) bool) {
		for _, co := range c.contacts {
			if !yield(co) {
				return
			}
		}
	}
}
