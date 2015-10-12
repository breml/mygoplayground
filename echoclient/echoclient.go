package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func reader(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Client got:", input) // Println will add back the final '\n'

		var m Message
		json.Unmarshal([]byte(input), &m)
		println("Name:", m.Name)
		println("Body:", m.Body)
		println("Time:", m.Time)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	go reader(c)
	m := Message{"Alice", "Hello0123456789_0123456789\n0123456789_0123456789", 1294706395881547000}
	for {
		b, err := json.Marshal(m)
		_, err = c.Write(b)
		_, err = c.Write([]byte("\n"))
		if err != nil {
			println(err)
			break
		}
		time.Sleep(1e9)
	}
}
