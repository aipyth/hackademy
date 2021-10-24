package main

import (
	"encoding/json"
	"log"
	"math/big"
	"net"
)

type Params struct {
	Number uint32 `json:"number"`
}

func CachedComputeFib(func(uint32) *big.Int) func(uint32) *big.Int {
	cache := make(map[uint32]*big.Int)
	return func(n uint32) *big.Int {
		res, ok := cache[n]
		if ok {
			return res
		}
		cache[n] = ComputeFib(n)
		return cache[n]
	}
}

func handleRequest(conn net.Conn) {
	defer func() {
		log.Println("Closed", conn.RemoteAddr())
		conn.Close()
	}()

	cComputeFib := CachedComputeFib(ComputeFib)

	for {
		params := &Params{}
		err := json.NewDecoder(conn).Decode(params)
		if err != nil {
			conn.Write([]byte(err.Error()))
		}

		res := cComputeFib(params.Number).String()

		log.Printf("Computed %d-th fibonacci number\n", params.Number)

		_, err = conn.Write([]byte(string(res + "\n")))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func ComputeFib(n uint32) *big.Int {
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)

	if n == 0 {
		return f0
	} else if n == 1 {
		return f1
	}

	for n > 1 {
		fn := big.NewInt(0)
		fn.Add(f0, f1)
		f0 = f1
		f1 = fn

		n--
	}
	return f1
}

func main() {
	listener, err := net.Listen("tcp", ":8012")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Println("Connected", conn.RemoteAddr())

		go handleRequest(conn)
	}
}
