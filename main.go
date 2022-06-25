package main

import (
	"fmt"

	print "github.com/baixinke/InferContinue/formatter"
	"github.com/baixinke/InferContinue/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
