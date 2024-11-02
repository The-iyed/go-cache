package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	serverAddress = "localhost:6379"
)

func main() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to ICache CLI! Type your commands below:")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("Failed to send command: %v\n", err)
			continue
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read response: %v\n", err)
			continue
		}

		fmt.Println(response)
	}
}
