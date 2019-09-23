package main

import (
	"bufio"
	"fmt"
	"net"
)

var routinID int = 0

func main() {

	listener, _ := net.Listen("tcp", ":8081")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Connection Error: ", err)
			conn.Close()
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			routinID = routinID + 1
			id := routinID

			fmt.Println("Connected ", conn.RemoteAddr(), " in routin=", id)

			bufReader := bufio.NewReader(conn)
			fmt.Println("Start reading...")

			for {
				rbyte, err := bufReader.ReadByte()
				if err != nil {
					fmt.Println("Read Error: ", err, conn.RemoteAddr())
					break
				}

				switch rbyte {
				case '\r':
				case '\n':
				default:
					fmt.Println("routinID=", id, " - char(", string(rbyte), ") key=", rbyte)
				}
			}
		}(conn)

	}

}
