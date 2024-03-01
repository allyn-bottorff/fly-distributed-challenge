package main

import (
	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return err
		}

		body := make(map[string]any)
		body["type"] = "generate_ok"
		body["id"] = newUUID.String()

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
