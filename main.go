package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

func timeout() (int, error) {
	var res int
	done := make(chan struct{})

	go func() {
		res = doWork(1)
		close(done)
	}()

	select {
	case <-done:
		return res, nil
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}

}

func doWork(t int) int {
	time.Sleep(time.Duration(t) * time.Second)

	return rand.Intn(10)
}

func main() {
	res, err := timeout()
	if err != nil {
		log.Print(err)

		return
	}

	log.Print("work ok, result: ", res)

}
