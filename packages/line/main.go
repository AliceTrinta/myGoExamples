package main

import "fmt"

func main()  {
	p1 := Point{2.0, 2.2}
	p2 := Point{2.0, 4.0}
	fmt.Println(cathetus(p1, p2))
	fmt.Println(Distance(p1, p2))
}
