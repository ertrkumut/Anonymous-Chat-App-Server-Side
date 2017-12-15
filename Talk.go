package main

type Talk struct {
	id           int
	playerOneId  string
	playerTwoId  string
	recieverName string
	messages     []*Message
}
