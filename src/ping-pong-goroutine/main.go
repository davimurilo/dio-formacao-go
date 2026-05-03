package main

import (
	"fmt"
	"time"
)

func Ping(canalPing chan string, pingPong chan struct{}, sinalizarPong chan struct{}) {

	for {
		<-pingPong
		time.Sleep(time.Second)
		canalPing <- "Ping"
		sinalizarPong <- struct{}{}
	}
}

func Pong(canalPong chan string, pingPong chan struct{}, aguardarPing chan struct{}) {
	for {
		<-aguardarPing
		time.Sleep(time.Millisecond * 500)
		canalPong <- "Pong"
		pingPong <- struct{}{}
	}
}

func main() {

	canalPing := make(chan string)
	canalPong := make(chan string)
	pingPong := make(chan struct{})
	sinalizacao := make(chan struct{})

	go Ping(canalPing, pingPong, sinalizacao)
	pingPong <- struct{}{}
	go Pong(canalPong, pingPong, sinalizacao)

	go func() {
		for {
			select {
			case msg := <-canalPing:
				fmt.Println(msg)
			case msg := <-canalPong:
				fmt.Println(msg)
			}
		}
	}()

	fmt.Scanln()
}
