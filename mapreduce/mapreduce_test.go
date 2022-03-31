package mapreduce

import (
	"fmt"
	"github.com/kevwan/mapreduce"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	val, err := mapreduce.MapReduce(
		func(source chan<- interface{}) {
			// generator
			for i := 0; i < 10; i++ {
				source <- i
			}
		},
		func(item interface{}, writer mapreduce.Writer, cancel func(error)) {
			// mapper
			i := item.(int)
			writer.Write(i * i)
		},
		func(pipe <-chan interface{}, writer mapreduce.Writer, cancel func(error)) {
			// reducer
			var sum int
			for i := range pipe {
				sum += i.(int)
			}
			writer.Write(sum)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
}

