package sync

import (
	"fmt"
	"time"
)

func ExampleSemaphore() {
	sem := New(2)
	go func() {
		sem.Acquire()
		fmt.Println("go 1 run")
		time.Sleep(1 * time.Second)
		sem.Release()
	}()

	go func() {
		sem.Acquire()
		fmt.Println("go 2 run")
		time.Sleep(1 * time.Second)
		sem.Release()
	}()

	go func() {
		sem.Acquire()
		fmt.Println("go 3 run")
		time.Sleep(1 * time.Second)
		sem.Release()
	}()

	go func() {
		sem.Acquire()
		fmt.Println("go 4 run")
		time.Sleep(1 * time.Second)
		sem.Release()
	}()

	//bufio.NewScanner(os.Stdin).Scan()
	time.Sleep(3 * time.Second)
}
