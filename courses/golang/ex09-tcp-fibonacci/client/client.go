package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"fmt"
	"time"
)

type askFibParams struct {
	Number uint64 `json:"number"`
}

func wrapInput(input string) []byte {
	s := strings.Replace(input, "\n", "", -1)
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	params := &askFibParams{
		Number: uint64(number),
	}
	data, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	return data
}

func main() {
	conn, err := net.Dial("tcp", ":8012")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)
runloop:
	for {
		select {
		case <-sigc:
			break runloop
		default:
			break
		}

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		input := wrapInput(text)

		start := time.Now()
		conn.Write(input)
		message, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		totalTime := []byte(time.Now().Sub(start).String())

		outputMessage := make([]byte, 0)
		outputMessage = append(outputMessage, totalTime...)
		outputMessage = append(outputMessage, ' ')
		outputMessage = append(outputMessage, message...)

		os.Stdout.Write(outputMessage)
	}
}
