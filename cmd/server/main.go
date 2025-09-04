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
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			slog.Error("Connection failed!: %e", err)
			continue
		}
		conn.Write([]byte("Welcome to TCP Chat!"))
		go func() {
			for {
				var buffer = make([]byte, 1024)
				_, err := conn.Read(buffer)
				if err != nil {
					break;
				}
				fmt.Println(string(buffer))
			}
		}()
	}
}
