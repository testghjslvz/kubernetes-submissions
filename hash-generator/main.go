package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	for {
		timestamp := time.Now().UTC().Format(time.RFC3339Nano)
		randomUUID := uuid.New()

		fmt.Printf("%s: %s\n", timestamp, randomUUID)

		time.Sleep(5 * time.Second)
	}
}
