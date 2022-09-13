package main

import (
	"fmt"
	sample_nodeps "github.com/brybacki/go-examples/sample-nodeps"
	sample_withdeps "github.com/brybacki/go-examples/sample-withdeps"
)

func main() {
	fmt.Println("yolo")
	result := sample_nodeps.Add(1, 2)
	fmt.Println(result)

	sample_withdeps.PrintHello()
	sample_withdeps.Print("aqq")
}
