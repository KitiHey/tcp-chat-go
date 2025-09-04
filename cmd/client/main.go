package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"net"
	"os"
	"bufio"
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
	conn, err := net.Dial("tcp", "localhost"+port)
	if err != nil {
		slog.Error("Error connecting!: %e", err)
		return
	}

	var buffer = make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

	go func() {
		for {
			var buffer = make([]byte, 1024)
			conn.Read(buffer)
			fmt.Println(string(buffer))
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(text[:len(text)-1]))
		if err != nil {
			slog.Error("Couldn't send message!: %e", err)
			continue
		}
	}
}
