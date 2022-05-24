package fsm

import (
	"fmt"
	"github.com/looplab/fsm"
	"testing"
)

func enterState(e *fsm.Event) {
	fmt.Printf("event: %s, from:%s to %s\n", e.Event, e.Src, e.Dst)
}

func TestFsmName(t *testing.T) {
	f := fsm.NewFSM(
		"sleeping",
		fsm.Events{
			{Name: "eat", Src: []string{"sleeping"}, Dst: "eating"},
			{Name: "work", Src: []string{"eating"}, Dst: "working"},
			{Name: "sleep", Src: []string{"working"}, Dst: "sleeping"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { enterState(e) },
		},
	)

	f.Event("eat")

	f.Event("work")

	f.Event("sleep")

}
