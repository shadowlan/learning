package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	server := args[1]
	c, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		// option 1: write and read data through bufio
		fmt.Println("sending message option 1...")
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("1->: " + message)
		//option 2: using connection methods
		fmt.Println("sending message option 2...")
		c.Write([]byte(text))
		buf := make([]byte,256)
		_,err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("2->: " + string(buf))
		if strings.TrimSpace(string(text)) == "exit" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
