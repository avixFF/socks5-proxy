package main

import (
	"fmt"
	"os"
	"strconv"

	socks5 "github.com/armon/go-socks5"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	address := "0.0.0.0"

	port, _ := strconv.ParseUint(getEnv("PORT", "1080"), 10, 32)

	fmt.Printf("Running SOCKS5 proxy on %s:%d\n", address, port)

	err = server.ListenAndServe("tcp", fmt.Sprintf("%s:%d", address, port))

	if err != nil {
		panic(err)
	}
}
