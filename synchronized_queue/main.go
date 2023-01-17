package main

import (
	"fmt"
	"time"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    int
	playTicket  int
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)
}

func (Queue *Queue) startTicketIssue() {
	Queue.message <- messageTicketStart
	<-Queue.queueTicket
}

func (Queue *Queue) EndTicketIssue() {
	Queue.message <- messageTicketEnd
}

func ticketIssue(queue *Queue) {
	for {
		time.Sleep(time.Second * 10)
		queue.startTicketIssue()
		fmt.Println("Ticket issue starts")
		time.Sleep(time.Second * 2)
		fmt.Println("Ticket issue ends")
		queue.EndTicketIssue()
	}
}

func (Queue *Queue) startPass() {
	Queue.message <- messagePassStart
	<-Queue.message
}

func (Queue *Queue) endPass() {
	Queue.message <- messagePassEnd
}

func passenger(Queue *Queue) {
	for {
	}
}

func main() {}
