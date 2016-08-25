package main

import (
	"fmt"
	"log"
	"time"
)

func wait1(c chan string) {
	time.Sleep(1 * time.Second)
	log.Print("waited 1 sec")
	c <- "wait1 finished\n"

}

func wait2(c chan string) {
	time.Sleep(2 * time.Second)
	log.Print("waited 2 sec")
	c <- "wait2 finished\n"
}

func wait3(c chan string) {
	time.Sleep(3 * time.Second)
	log.Print("waited 3 sec")
	c <- "wait3 finished\n"
}

func main() {
	c := make(chan string)
	log.Print("started")
	go wait1(c)
	go wait2(c)
	go wait3(c)
	w1, w2, w3 := <-c, <-c, <-c
	log.Print("finished")
	fmt.Println(w1, w2, w3)
}
