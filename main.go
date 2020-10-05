package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleConnections(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))

		if temp == "STOP" {
			break
		}

		price := rand.Int()
		result := strconv.Itoa(price) + "\n"

		conn.Write([]byte(string(result)))

	}
}

func main() {
	PORT := os.Getenv("ENV_PORT")
	fmt.Printf("Server listen on %s\n", PORT)

	l, err := net.Listen("tcp4", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go handleConnections(c)
	}

}