package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var count = 0

func main() {
	//1. listen on :8080 but accept only one call
	// s, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	fmt.Printf("fail to set up server: %v", err)
	// 	return
	// }
	// s.Close()
	// l, err := s.Accept()
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	// for {
	// 	data, err := bufio.NewReader(l).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Printf("%v", err)
	// 		return
	// 	}
	// 	if strings.TrimSpace(data) == "exit" {
	// 		fmt.Println("TCP server exiting...")
	// 		return
	// 	}
	// 	t := time.Now().Format("2006-01-02 15:04:05")
	// 	l.Write([]byte(t + ":" + data + "\n"))
	// }

	// 2. listen on :8080 and build concurrent tcp server
	s, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("fail to set up server: %v", err)
		return
	}
	defer s.Close()
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		go handler(c)
		count++
	}
}

func handler(c net.Conn) {
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		if strings.TrimSpace(data) == "exit" {
			fmt.Println("break")
			break
		}
		counter := strconv.Itoa(count)
		t := time.Now().Format("2006-01-02 15:04:05")
		c.Write([]byte(counter + ":" + t + ":" + data + "\n"))
	}
	c.Close()
}
