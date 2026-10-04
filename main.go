package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
		fmt.Println("[USAGE]: go run main.go [port]")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	address := ":" + port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close() //Ensure the listener closes when main() exits
	fmt.Printf("Listening on the port %s\n", address)

	// wait until client connects
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer conn.Close()

	// read what the client sent and print it (just for testing)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Failed to read message: %v", err)
		return
	}
	fmt.Printf("Recieved: %s", buf[:n])
}
