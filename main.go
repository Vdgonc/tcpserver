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

func Timer(ch <-chan bool) {
	started := time.Now()
	for {
		stop := <-ch
		if stop {
			stopped := time.Now()
			diff := stopped.Sub(started)
			log.Printf("First connection %.0fs", diff.Seconds())
			break
		}
	}
}

func HandleConnections(conn net.Conn) {
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

		if temp == "CHECK" {
			ok := "ok\n"
			conn.Write([]byte(string(ok)))
		} else {

			number := rand.Int()
			log.Printf("Sending number: %v", number)
			result := strconv.Itoa(number) + "\n"

			conn.Write([]byte(string(result)))
		}

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

	ch := make(chan bool, 1)
	go Timer(ch)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		ch <- true
		go HandleConnections(c)
	}

}
