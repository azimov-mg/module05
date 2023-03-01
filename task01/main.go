package main

import "fmt"

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	panic("Panic in foo()")
	return err
}

func main() {
	if err := foo(); err != nil {
		// Обрабатываем ошибку
		fmt.Println(err)
	}
}
