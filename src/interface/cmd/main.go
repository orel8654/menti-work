package main

import (
    "log"
    "menti/src/interface/internal/delivery"
)

func main() {
	if err := run(":8005"); err != nil {
		log.Fatal(err)
	}
}

func run(host string) error {
	handler := delivery.NewHandlers()
	return handler.Listen(host)
}

