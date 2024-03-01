package main

import (
	"fmt"
	"time"
)

type theatre struct {
	rows            int
	seats           int
	seatingPosition [][]int
}

type user struct {
	name                   string
	num_of_seats_requested int
	ticket_no              int
}

type ticket struct {
	price       int
	id          int
	total_seats int
	start_seat  int
	end_seat    int
}

func (t theatre) print_receipt(u *user, tic *ticket) {
	fmt.Println("Ticket Details ")
	fmt.Println(" User Name : ", u.name)
	fmt.Println(" Ticket Number : ", tic.id)
	fmt.Println(" Num of seats : ", tic.total_seats)
	fmt.Println(" Price : ", tic.price)
	fmt.Println(" Start Seat : ", tic.start_seat)
	fmt.Println(" End seat : ", tic.end_seat)
}

func (u *user) book_ticket(t *theatre, n int, start int, end int, coupon string) ticket {
	tic := ticket{
		price:       n * 100,
		total_seats: n,
		start_seat:  start + end,
	}
	if coupon == TRYNEW {
		tic.price -= 200
		if tic.price < 0 {
			tic.price = 0
		}
	} else if coupon == PLUS1 {
		n++
		tic.total_seats = n
	}
	for n > 0 {
		if t.seatingPosition[start][end] != 1 {
			t.seatingPosition[start][end] = 1
			end++
		} else {
			fmt.Println(" Seat Already booked ")
			break
		}
		n--
	}

	u.ticket_no = time.Now().Day() + time.Now().Minute()
	u.num_of_seats_requested = n
	fmt.Println(" Ticket booked for ", start+1, "th row ", end, " seat")
	tic.id = u.ticket_no
	tic.end_seat = end
	return tic
}

const TRYNEW = "200 OFF"
const PLUS1 = "One Extra Ticket"

func main() {

	t := theatre{
		rows:  2,
		seats: 10,
		seatingPosition: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	x := user{
		name: "XX",
	}
	tic := x.book_ticket(&t, 2, 0, 2, PLUS1)
	t.print_receipt(&x, &tic)
	fmt.Println(t.seatingPosition)
	y := user{
		name: "YY",
	}
	tic1 := y.book_ticket(&t, 1, 0, 5, TRYNEW)
	t.print_receipt(&y, &tic1)
	fmt.Println(t.seatingPosition)
}
