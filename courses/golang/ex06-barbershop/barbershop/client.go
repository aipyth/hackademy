package barbershop

import (
	"log"
	"math/rand"
	"time"
)

type Client struct {
	id uint
}

func (c *Client) CheckWaitingRoom(wr *WaitingRoom) WaitingRoomStatus {
	if len(wr.seats) == 0 {
		log.Printf("Client %d found empty waiting room\n", c.id)
		return WaitingRoomEmpty
	} else if len(wr.seats) == cap(wr.seats) {
		log.Printf("Client %d found full waiting room\n", c.id)
		return WaitingRoomFull
	} else {
		log.Printf("Client %d sees %d free seats out of %d\n", c.id, cap(wr.seats)-len(wr.seats), cap(wr.seats))
		return WaitingRoomFreeSeats
	}
}

func (c *Client) CheckBarber(barber *Barber, wr *WaitingRoom) BarberStatus {
	select {
	case wr.checkingBarbers <- c:
		defer func() {
			<-wr.checkingBarbers
		}()
		log.Printf("Client %d checks barber's room\n", c.id)
		if barber.sleeping {
			log.Printf("Client %d sees barber's sleeping\n", c.id)
			return BarberSleeping
		} else if barber.busy {
			log.Printf("Client %d sees barber's busy\n", c.id)
			return BarberBusy
		} else {
			log.Printf("Client %d sees that barber neither sleeping not busy. Weird stuff...\n", c.id)
			return BarberWTF
		}
	default:
		log.Printf("Client %d sees somebody is checking barber's room\n", c.id)
		return BarberBusy
	}
}

func (c *Client) TakeAWaitingRoomSeat(wr *WaitingRoom) bool {
	select {
	case wr.seats <- c:
		log.Printf("Client %d takes a seat in the waiting room\n", c.id)
		return true
	default:
		log.Printf("Somebody's took a seat before client did. Client %d is angry and walks away!\n", c.id)
		return false
	}
}

func (c *Client) WakeUpBarber(b *Barber) {
	log.Printf("Client %d wakes up barber and takes a seat for barber to cut his hair\n", c.id)
	b.clientNow = c
	b.WakeUp()
}

func (c *Client) WalkInto(wr *WaitingRoom, b *Barber) bool {
	switch c.CheckWaitingRoom(wr) {
	case WaitingRoomEmpty:

		switch c.CheckBarber(b, wr) {
		case BarberSleeping:
			c.WakeUpBarber(b)
			return true
		case BarberBusy:
			return c.TakeAWaitingRoomSeat(wr)
		case BarberWTF:
			return false
		}
	case WaitingRoomFull:
		return false
	case WaitingRoomFreeSeats:
		return c.TakeAWaitingRoomSeat(wr)
	}
	return false
}

func (c *Client) GoToBarber(wr *WaitingRoom, b *Barber) {
	roadTime := time.Duration(rand.Intn(10000)) * time.Millisecond
	time.Sleep(roadTime)
	log.Printf("Client %d arrived\n", c.id)
	for {
		if c.WalkInto(wr, b) {
			break
		}
		time.Sleep(roadTime)
		log.Printf("Client %d arrived again\n", c.id)
	}
}
