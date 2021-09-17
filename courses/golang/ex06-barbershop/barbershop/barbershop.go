package barbershop

import (
	"log"
)

type Barbershop struct {
	barber      *Barber
	clients     func() *Client
	waitingRoom *WaitingRoom
}

func clientGenerator(num uint) func() *Client {
	i := uint(0)
	log.Printf("[ClientGenerator] today %d clients are going to come\n", num)
	return func() *Client {
		defer func() {
			i++
		}()
		if i >= num {
			return nil
		}
		return &Client{
			id: i,
		}
	}
}

func NewBarbershop() *Barbershop {
	waitingRoom, err := NewWaitingRoom(3)
	if err != nil {
		log.Println(err)
		return nil
	}
	clientsToCome := uint(cap(waitingRoom.seats)) * 10
	return &Barbershop{
		waitingRoom: waitingRoom,
		barber:      NewBarber(clientsToCome),
		clients:     clientGenerator(clientsToCome),
	}
}

func (bs *Barbershop) RunBarbershop() {
	go bs.directClients()
	bs.barber.Run(bs.waitingRoom)
}

func (bs *Barbershop) directClients() {
	var client *Client
	for {
		client = bs.clients()
		if client == nil {
			break
		}
		go client.GoToBarber(bs.waitingRoom, bs.barber)
	}
}
