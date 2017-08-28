// Building upon the code from the previous exercise:
//
// Use bufio.NewScanner() to read from the connection.
// After all of the reading, include these lines of code:
//
// fmt.Println("Code got here.")
// io.WriteString(c, "I see you connected.")
//
// Launch your TCP server.
// In your **web browser,** visit localhost:8080.
// Now go back and look at your terminal.
// Can you answer the question as to why "I see you connected." is never written?

// Answer to question above:  for s.Scan() never breaks, so never get to io.WriteString call

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := li.Accept()
		defer conn.Close()

		if err != nil {
			log.Fatalln(err)
		}

		s := bufio.NewScanner(conn)

		for s.Scan() {
			ln := s.Text()
			fmt.Fprint(conn, ln)
			fmt.Println(ln)
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected\n")
	}
}
