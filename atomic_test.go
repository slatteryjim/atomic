package atomic

import (
	"testing"
	"time"

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

func TestTime(t *testing.T) {
	RegisterTestingT(t) // for Gomega matchers

	now := time.Now()
	earlier := now.Add(-5 * time.Minute)
	Ω(now).ShouldNot(Equal(earlier))

	// create new Time instance with initial value
	at := NewTime(now)
	Ω(at.Val()).Should(Equal(now))

	// Set a different value
	at.Set(earlier)
	Ω(at.Val()).Should(Equal(earlier))

	// atomically add 3 seconds to the time
	at.Alter(func(old time.Time) time.Time { return old.Add(3 * time.Second) })
	Ω(at.Val().Sub(earlier)).Should(Equal(3 * time.Second))

	// zero value supported
	var x Time
	Ω(x.Val().Nanosecond()).Should(Equal(0))
}

func TestBool(t *testing.T) {
	RegisterTestingT(t) // for Gomega matchers

	// Create with initial value
	Ω(NewBool(true).Val()).Should(Equal(true))
	Ω(NewBool(false).Val()).Should(Equal(false))

	// Set different value
	b := NewBool(true)
	b.Set(false)
	Ω(b.Val()).Should(Equal(false))

	b.Set(true)
	Ω(b.Val()).Should(Equal(true))

	// zero value supported
	var x Bool
	Ω(x.Val()).Should(Equal(false))
}

func TestMap(t *testing.T) {
	RegisterTestingT(t) // for Gomega matchers

	b := NewMap()
	b.Set("test", true)
	Ω(b.Get("test")).Should(Equal(true))

	b.Set(123, false)
	Ω(b.Get(123)).Should(Equal(false))

	b.Del(123)
	Ω(b.Get(123)).Should(BeNil())

	Ω(b.Get(456)).Should(BeNil())
}
