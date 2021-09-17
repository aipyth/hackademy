package barbershop

import "errors"

type WaitingRoom struct {
	seats           chan *Client
	checkingBarbers chan *Client
}

type WaitingRoomStatus int

const (
	WaitingRoomEmpty WaitingRoomStatus = iota
	WaitingRoomFull
	WaitingRoomFreeSeats
)

func NewWaitingRoom(seatsNumber uint) (*WaitingRoom, error) {
	if seatsNumber == 0 {
		return nil, errors.New("seatsNumber cannot be 0")
	}
	return &WaitingRoom{
		seats:           make(chan *Client, seatsNumber),
		checkingBarbers: make(chan *Client, 1),
	}, nil
}
