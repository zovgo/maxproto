package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/zovgo/maxproto"
	"github.com/zovgo/maxproto/protocol"
)

func main() {
	_ = godotenv.Load(".env")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	conf := maxproto.DialConfig{
		Token:     os.Getenv("TOKEN"),
		DeviceID:  uuid.MustParse("49951d64-8a7a-4887-ac80-42cdf8e3a9c9"),
		ChatCount: 40,
	}
	fmt.Println("dialing...")
	client, err := conf.DialContext(ctx, false)
	if err != nil {
		panic(fmt.Errorf("dial: %w", err))
	}
	defer client.Close() //nolint:errcheck
	fmt.Println("successfully connected and logged in:", client.Conn().RemoteAddr().String())
	fmt.Println("waiting for messages...")
	wait(client)
	fmt.Println("stopped waiting for messages")
	//TODO: max website ITSELF drops connection sometimes so we can infinitely dial until ctrl+c
}

func wait(cl *maxproto.Client) {
	err := cl.WaitForMessages(handleMessage(cl))
	if err != nil && !errors.Is(err, maxproto.ErrClientClosed) {
		panic(fmt.Errorf("wait: %w", err))
	}
}

func handleMessage(cl *maxproto.Client) func(protocol.Message) {
	return func(message protocol.Message) {
		if message.Type != "USER" {
			fmt.Println("not a user message:", message.Type)
			return
		}
		c, ok := cl.GetContact(message.Sender)
		if !ok {
			fmt.Println("contact not found:", message.Sender)
			return
		}
		printName(c)

		ok = true
		for _, a := range message.Attaches {
			ok = false
			fmt.Printf("(%s)", strings.ToLower(a.Type))
		}
		if !ok {
			return
		}
		fmt.Println(message.Text)
	}
}

func printName(c protocol.Contact) {
	fmt.Println(name(c))
}

func name(c protocol.Contact) string {
	fall := "undefined"
	for _, n := range c.Names {
		if n.Type != "ONEME" {
			fall = cleanName(n)
			return fall
		}
		return cleanName(n)
	}
	return fall
}

func cleanName(n protocol.ContactName) string {
	if n.LastName != "" {
		return n.FirstName + " " + n.LastName
	}
	return n.FirstName
}
