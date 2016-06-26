package atomic

import (
	"testing"
	. "github.com/onsi/gomega"
)

func TestInt32(t *testing.T) {
	RegisterTestingT(t) // for Gomega matchers

	// create new Int32 instance with initial value
	ai := NewInt32(int32(1))
	Ω(ai.Val()).Should(Equal(int32(1)))

	// atomically add 1
	ai.Add(int32(1))
	Ω(ai.Val()).Should(Equal(int32(2)))

	// atomically set value back to 1
	ai.Set(int32(1))
	Ω(ai.Val()).Should(Equal(int32(1)))

	// atomically swap value to 2
	oldValue := ai.Swap(int32(2))
	Ω(oldValue).Should(Equal(int32(1)))
	Ω(ai.Val()).Should(Equal(int32(2)))

	// zero value supported
	var x Int32
	Ω(x.Val()).Should(Equal(int32(0)))

}