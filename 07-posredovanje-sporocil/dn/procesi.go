package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

type message struct {
	data   []byte
	length int
}

var N int
var id int
var rootPort int

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func spreadMessage(conn *net.UDPConn) {
	msg := []byte("Root message")

	for process := 1; process < N; process++ {
		sent := false
		targetAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", rootPort+process))
		checkError(err)

		for attempt := 1; attempt <= 5; attempt++ {
			fmt.Println("[Root] Sending message to", process, "attempt #", attempt)
			_, err := conn.WriteToUDP(msg, targetAddr)
			checkError(err)

			conn.SetDeadline(time.Now().Add(500 * time.Millisecond))

			buffer := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buffer)
			conn.SetDeadline(time.Time{})
			if err == nil && string(buffer[:n]) == "ACK" {
				fmt.Println("[Root] Received ACK from", addr)
				sent = true
				break
			}

			time.Sleep(500 * time.Millisecond)
		}
		if !sent {
			fmt.Println("[Root] Failed to send message to process", process)
		}
	}
}

func receiveMessage(conn *net.UDPConn) {
	received := false
	buffer := make([]byte, 1024)

	// Simuliranje zamude pri začetku
	time.Sleep(time.Duration(id) * 1 * time.Second)

	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("[Process", id, "] Timeout waiting for message.")
			break
		}

		if !received {
			fmt.Println("[Process", id, "] Received message:", "'"+string(buffer[:n])+"'", "from", addr)
			_, err := conn.WriteToUDP([]byte("ACK"), addr)
			checkError(err)
			received = true
		}
	}
}

func main() {
	idPtr := flag.Int("id", 0, "ID procesa")
	NPtr := flag.Int("N", 5, "Stevilo procesov v sistemu")
	rootPtr := flag.Int("root", 9000, "Osnovna vrata za root proces")
	flag.Parse()

	id = *idPtr
	N = *NPtr
	rootPort = *rootPtr

	localAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", rootPort+id))
	checkError(err)

	conn, err := net.ListenUDP("udp", localAddr)
	checkError(err)
	defer conn.Close()

	if id == 0 {
		fmt.Println("Root started.")
		spreadMessage(conn)
	} else {
		fmt.Println("Process", id, "waiting for message.")
		receiveMessage(conn)
	}
}
