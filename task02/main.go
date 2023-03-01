package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("Goroutie %d is started to acting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Goroutie %d is finished to acting\n", id)
}

func execute() error {
	start := time.Now()

	for i := 1; i <= 4; i++ {
		go doWork(i)
	}

	time.Sleep(3 * time.Second)

	elapsed := time.Since(start)
	if elapsed > 2*time.Second {
		return fmt.Errorf("the time of acting is exceeded 2 sec")
	}

	fmt.Println("All goroutines finished acting")
	return nil
}

func main() {
	if err := execute(); err != nil {
		// Обрабатываем ошибку
		fmt.Println(err)
	}
}
