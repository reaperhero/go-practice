package consistent

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"testing/quick"
)

func TestName(t *testing.T) {

	c := New()
	c.Add("cacheA")
	c.Add("cacheB")
	c.Add("cacheC")
	users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer"}
	fmt.Println("initial state [A, B, C]")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server)
	}
	c.Add("cacheD")
	c.Add("cacheE")
	fmt.Println("\nwith cacheD, cacheE [A, B, C, D, E]")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server)
	}
	// Output:
	// initial state [A, B, C]
	// user_mcnulty => cacheA
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheC
	// user_stringer => cacheC
	//
	// with cacheD, cacheE [A, B, C, D, E]
	// user_mcnulty => cacheE
	// user_bunk => cacheA
	// user_omar => cacheA
	// user_bunny => cacheE
	// user_stringer => cacheE
}

func TestGetNQuick(t *testing.T) {
	x := New()
	x.Add("abcdefg")
	x.Add("hijklmn")
	x.Add("opqrstu")

	f := func(s string) bool {
		members, err := x.GetN(s, 3)
		fmt.Println(members)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		if len(members) != 3 {
			t.Logf("expected 3 members instead of %d", len(members))
			return false
		}
		set := make(map[string]bool, 4)
		for _, member := range members {
			if set[member] {
				t.Logf("duplicate error")
				return false
			}
			set[member] = true
			if member != "abcdefg" && member != "hijklmn" && member != "opqrstu" {
				t.Logf("invalid member: %q", member)
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}
