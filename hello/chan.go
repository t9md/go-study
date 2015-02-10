package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- (strconv.Itoa(i) + ":ping")
	}
}
func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- (strconv.Itoa(i) + ":pong")
	}
}
func printer(c <-chan string, quit <-chan bool) {
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
			time.Sleep(time.Millisecond * 500)
		case <-quit:
			return
		}
	}
}

func mySleep(sec int) {
	fmt.Printf("Now start to sleep for %s sec\n", sec)
	for {
		select {
		case <-time.After(time.Second * time.Duration(sec)):
			return
		}
	}
}
func main() {
	var c chan string = make(chan string)
	quit := make(chan bool)

	mySleep(5)
	fmt.Println("Waked Up!!")
	go pinger(c)
	go ponger(c)
	go printer(c, quit)
	time.Sleep(time.Second * 10)
	quit <- true
	fmt.Println("Finish")

	fmt.Println("# Next Chapter")
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg := <-c1:
				fmt.Println(msg)
			case msg := <-c2:
				fmt.Println(msg)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
