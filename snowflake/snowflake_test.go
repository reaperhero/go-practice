package snowflake_test

import (
	"go-example-demo/snowflake"
	"log"
	"testing"
)

func TestSnowId(t *testing.T) {
	flake := snowflake.NewSnowflake()
	for i := 0; i < 100; i++ {
		log.Println(flake.NextVal())
	}
}
