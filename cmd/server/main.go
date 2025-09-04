package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"net"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error(".env couldn't be loaded!: %e", err)
		return
	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		slog.Error(".env doesn't have a 'PORT'!")
		return
	}

	fmt.Printf("Initializing server...\n")
	ln, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error("Error listening!: %e", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("Connection failed!: %e", err)
		}
		conn.Write([]byte("Welcome to TCP Chat!"))
		var buffer = make([]byte, 1024)
		conn.Read(buffer)
		fmt.Println(string(buffer))
	}
}
