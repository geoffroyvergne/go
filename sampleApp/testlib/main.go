package main

import (
	"fmt"
	"log"
	lib "sampleApp/testapp/lib"
)

func main() {
	fmt.Println(lib.GetLocalValue())
	log.Println(lib.GetLocalValue())
}