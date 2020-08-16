package main

import (
	"fmt"
	"time"
)

func talk(person string) <-chan string{
	c := make(chan string)
	go func() {
		for i := 0; i <3; i++{
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s talking: %d", person, i)
		}
	}()
	return c
}

func add(input1, input2 <-chan string) <-chan string  {
	c := make(chan string)
	go func(){
		for{
			select {
			case s := <- input1:
				c <- s
			case s := <- input2:
				c <- s
			}
		}
	}()
	return c
}

func main()  {
	c := add(talk("Mary"), talk("Jhon"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
