package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
	extLib "sampleApp/testlib"
	lib "sampleApp/testapp/lib"
)


func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(extLib.GetExternalValue())
	fmt.Println(lib.GetLocalValue())

	i := 0

	fmt.Println("Hello world", i)

	fmt.Println("Time is", time.Now())

	fmt.Println("Rand number", rand.Intn(10))

	fmt.Printf("Now you have %g problems.", math.Sqrt(7))

	log.Print("Test Logs")

	fmt.Println(add(42, 13))
}
