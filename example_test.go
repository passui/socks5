package socks5_test

import (
	"log"

	"github.com/passui/socks5"
)

func ExampleServer() {
	s, err := socks5.NewClassicServer("127.0.0.1:1080", "127.0.0.1", 0, 60)
	if err != nil {
		log.Println(err)
		return
	}
	// You can pass in custom Handler
	s.ListenAndServe(nil)
	// #Output:
}
