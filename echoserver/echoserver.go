package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func echoServer(c net.Conn) {
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		input := scanner.Bytes()
		fmt.Println("Server got:", string(input)) // Println will add back the final '\n'

		var m Message
		json.Unmarshal(input, &m)
		println("Name:", m.Name)
		println("Body:", m.Body)
		println("Time:", m.Time)

		_, err := c.Write(input)
		_, err = c.Write([]byte("\n"))
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func main() {
	socketfile := "/tmp/echo.sock"
	if _, err := os.Stat(socketfile); err == nil {
		os.Remove(socketfile)
	}
	l, err := net.Listen("unix", socketfile)
	if err != nil {
		println("listen error", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			println("accept error", err)
			return
		}

		go echoServer(c)
	}
}
