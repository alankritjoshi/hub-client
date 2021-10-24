package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func outgoingRoutine(outC chan<- string) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	for {
		o, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Outgoing error: %v", err)
			return
		}
		outC <- o
	}
}

func incomingRoutine(inC chan<- string, c net.Conn) {
	connReader := bufio.NewReader(c)
	for {
		i, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Incoming error: %v", err)
			return
		}
		inC <- i
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	conn, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	outgoing := make(chan string)
	go outgoingRoutine(outgoing)

	incoming := make(chan string)
	go incomingRoutine(incoming, conn)

	for {
		select {
		case out := <-outgoing:
			fmt.Print(">> ")
			fmt.Fprintf(conn, out+"\n")
		case in := <-incoming:
			fmt.Print("\n<- " + in + ">> ")
		}
	}

}
