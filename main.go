package main

import (
	"fmt"
	"log"
	"os"
	"TCPChat/server"
)

// Owner: Member A
//
// main.go only STARTS the program. Keep it short.
//
// TODO (Member A):
//   - Read the port from the command line (none -> 8989, one -> that port,
//     anything else -> print the usage message exactly as in the subject).
//   - Call the server's start function (see CHECKLIST.txt, Section 3).
//   - Print the error and exit if the server cannot start.

func main() {
	port := "8989"

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	if !server.ValidPort(port) {
		fmt.Println("Invalid Port: Port must be in this range 1024 - 65535")
		return
	}
	
	err := server.Start(port)
	if err != nil {
		log.Fatal(err)
	}
}
