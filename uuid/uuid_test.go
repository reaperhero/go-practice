package uuid

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func Test_uuid_01(t *testing.T) {
	uuid := uuid.NewV4().String()
	fmt.Println(uuid)
}
