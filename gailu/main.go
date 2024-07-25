package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ss := "0.0045"
	float, _ := strconv.ParseFloat(ss, 64)
	floatI := int64(float * float64(time.Second))
	floatN := time.Nanosecond * time.Duration(floatI)

	fmt.Println(floatN.Seconds())
}
