package cmp

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGoogleCmp(t *testing.T) {

	type Contact struct {
		Phone string
		Email string
	}

	type User struct {
		Name    string
		Age     int
		Contact *Contact
	}

	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 18}

	c1 := &Contact{Phone: "123456789", Email: "dj@example.com"}
	c2 := &Contact{Phone: "123456789", Email: "dj@example.com"}

	u1.Contact = c1
	u2.Contact = c1
	assert.Equal(t, true, cmp.Equal(u1, u2))

	u2.Contact = c2
	assert.Equal(t, true, cmp.Equal(u1, u2)) // 比较的是值

	type FloatPair struct {
		X float64
		Y float64
	}
	p1 := FloatPair{X: math.NaN()}
	p2 := FloatPair{X: math.NaN()}
	assert.Equal(t, true, cmp.Equal(p1, p2, cmpopts.EquateNaNs()))

	f1 := 0.1
	f2 := 0.2
	f3 := 0.3
	p3 := FloatPair{X: f1 + f2}
	p4 := FloatPair{X: f3}
	assert.Equal(t, true, cmp.Equal(p3, p4, cmpopts.EquateApprox(0.1, 0.001))) // 两个数的差的绝对值小于这个数即|x-y| ≤ max(fraction*min(|x|, |y|), margin)，则认为它们相等

	var s1 []int
	var s2 = make([]int, 0)

	var m1 map[int]int
	var m2 = make(map[int]int)

	assert.Equal(t, true, cmp.Equal(s1, s2, cmpopts.EquateEmpty()))
	assert.Equal(t, true, cmp.Equal(m1, m2, cmpopts.EquateEmpty()))

	// 比较无序slice
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{4, 3, 2, 1}
	assert.Equal(t, true, cmp.Equal(sl1, sl2, cmpopts.SortSlices(func(i, j int) bool { return i < j })))


	// 比较无序map
	mp1 := map[int]int{1: 10, 2: 20, 3: 30}
	mp2 := map[int]int{1: 10, 2: 20, 3: 30}
	assert.Equal(t, true, cmp.Equal(mp1, mp2, cmpopts.SortMaps(func(i, j int) bool { return i < j })))
}
