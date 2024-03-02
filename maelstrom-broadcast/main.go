package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastMessage struct {
	MType string `json:"type"`
	Message  int    `json:"message"`
}

func main() {

	n := maelstrom.NewNode()

	state := make([]int, 0)

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		body := make(map[string]any)
		msgBody := new(BroadcastMessage)
		body["type"] = "broadcast_ok"
		if err := json.Unmarshal(msg.Body, &msgBody); err != nil {
			return err
		}

		state = append(state, msgBody.Message)
		return n.Reply(msg, body)
	})
	n.Handle("read", func(msg maelstrom.Message) error {
		body := make(map[string]any)
		body["type"] = "read_ok"
		body["messages"] = state
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
