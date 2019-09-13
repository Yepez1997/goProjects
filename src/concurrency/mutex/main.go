package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

// func init() {
// 	balance = 1000
// }

// deposit - put money into balance
func deposit(value int, done chan bool) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	done <- true
}

// withdraw - take out money from balance
func withdraw(value int, done chan bool) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from acocunt with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	done <- true
}

// use mutex to prevent race conditions (overriding happens)
// use go programs safely with mutex
// mutex (mutual exclusion) prevent concurrent process from entering a critical section
func main() {
	init := func() {
		balance = 1000
	}
	init()
	fmt.Println("Go Mutex")
	done := make(chan bool)
	go withdraw(700, done)
	go deposit(500, done)
	// alternative is to use a waitgroup
	<-done
	<-done
	fmt.Printf("New Balance %d\n", balance)
}
