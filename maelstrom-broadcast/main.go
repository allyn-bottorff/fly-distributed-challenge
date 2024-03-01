package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	n := maelstrom.NewNode()

	ch := make(chan []int)

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		body := make(map[string]any)
		messageContents := make(map[string]any)
		body["type"] = "broadcast_ok"
		if err := json.Unmarshal(msg.Body, &messageContents); err != nil {
			return err
		}
		seenMessages := <-ch
		seenMessages = append(seenMessages, messageContents["message"].(int))
		ch <- seenMessages
		return n.Reply(msg, body)
	})
	n.Handle("read", func(msg maelstrom.Message) error {
		body := make(map[string]any)
		body["type"] = "read_ok"
		messages := <-ch
        ch <- messages
		body["messages"] = messages
		return n.Reply(msg, body)
	})
	n.Handle("topology", func(msg maelstrom.Message) error {
		body := make(map[string]any)
		body["type"] = "topology_ok"
		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
