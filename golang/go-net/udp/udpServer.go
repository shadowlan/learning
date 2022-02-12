package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	// // option 1,using net.ListenUDP directly
	// addr, err := net.ResolveUDPAddr("udp", ":8088")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// connection, err := net.ListenUDP("udp", addr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// option 2, using net.ListenPacket()
	// net.ListenPacket("udp", ":0"), 如果使用":0"作为地址，那么将随机自动分配一个端口
	connection, err := net.ListenPacket("udp", ":0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		t := time.Now().Format("2006-01-02 15:04:05")
		// option 1
		//n, addr, err := connection.ReadFromUDP(buffer)
		// option 2
		n, addr, err := connection.ReadFrom(buffer)
		fmt.Println("-> "+t, string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "exit" {
			fmt.Println("UCP server exiting...")
			return
		}
		data := []byte(t + "->" + string(buffer[0:n-1]))
		// option 1
		//_, err = connection.WriteToUDP(data, addr)
		// option 2
		_, err = connection.WriteTo(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
