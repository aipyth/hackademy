package barbershop

import (
	"log"
	"time"
)

type Barber struct {
	sleeping      bool
	busy          bool
	clientNow     *Client
	clientsToCome uint
	cutClients    uint
}

type BarberStatus int

const (
	BarberSleeping BarberStatus = iota
	BarberBusy
	BarberWTF
)

func NewBarber(clientsToCome uint) *Barber {
	return &Barber{
		sleeping:      true,
		busy:          false,
		clientsToCome: clientsToCome,
	}
}

func (b *Barber) InviteClient(wr *WaitingRoom) bool {
	select {
	case client := <-wr.seats:
		log.Printf("Barber invites next client %d\n", client.id)
		b.clientNow = client
		b.busy = true
		return true
	default:
		log.Printf("Barber unsuccessfully tried to invite a client\n")
		return false
	}
}

func (b *Barber) Work() {
	b.busy = true
	log.Printf("Barber works on client %d\n", b.clientNow.id)
	time.Sleep(1 * time.Second)
	b.EndUpWithClient()
}

func (b *Barber) WakeUp() {
	log.Printf("Barber wakes up")
	b.sleeping = false
}

func (b *Barber) Sleep() {
	log.Printf("Barber takes a nap\n")
	b.sleeping = true
}

func (b *Barber) EndUpWithClient() {
	log.Printf("Barber finished cutting hair and punches the client %d away\n", b.clientNow.id)
	b.cutClients++
	b.clientNow = nil
	b.busy = false
}

func (b *Barber) CheckWaitingRoom(wr *WaitingRoom) WaitingRoomStatus {
	if len(wr.seats) == 0 {
		log.Println("Barber finds waiting room empty")
		return WaitingRoomEmpty
	} else if len(wr.seats) == cap(wr.seats) {
		log.Println("Barber takes a look at waiting room and shouts 'Such a surprise to see full waiting room!'")
		return WaitingRoomFull
	} else {
		log.Println("Barber sees some clients in waiting room")
		return WaitingRoomFreeSeats
	}
}

func (b *Barber) AreClientsDone() bool {
	return b.cutClients == b.clientsToCome
}

func (b *Barber) Run(wr *WaitingRoom) {
out:
	for {
		if (b.clientNow != nil) && !b.sleeping {
			b.Work()
		}
		if (b.clientNow == nil) && !b.sleeping {
			switch b.CheckWaitingRoom(wr) {
			case WaitingRoomFull, WaitingRoomFreeSeats:
				if !b.InviteClient(wr) {
					b.Sleep()
				}
			case WaitingRoomEmpty:
				if b.AreClientsDone() {
					log.Printf("Barber checks his records and found that he already cut all clients' hair!!!\n")
					break out
				} else {
					log.Printf("Barber checks his records and found that there are clients to come\n")
					b.Sleep()
				}
			}
		}
	}
}
