package atomic

import (
	"sync"
	"sync/atomic"
	"time"
)

//-----------------------------------------------------------------------------
// Int32
//-----------------------------------------------------------------------------

// Int32 stores an int32 value and allows it to be read and modified atomically.
type Int32 struct {
	val int32
}

func NewInt32(initialValue int32) *Int32 {
	return &Int32{initialValue}
}

// Add atomically adds the given delta to the stored value.
func (ai *Int32) Add(delta int32) int32 {
	return atomic.AddInt32(&ai.val, delta)
}

// Val safely returns the stored value.
func (ai *Int32) Val() int32 {
	return atomic.LoadInt32(&ai.val)
}

// Set safely sets the stored value.
func (ai *Int32) Set(newValue int32) {
	atomic.StoreInt32(&ai.val, newValue)
}

// Swap safely swaps the stored value and returns the old value
func (ai *Int32) Swap(newValue int32) int32 {
	return atomic.SwapInt32(&ai.val, newValue)
}

//-----------------------------------------------------------------------------
// Time
//-----------------------------------------------------------------------------

// Time stores a time.Time and allows it to be read and modified atomically.
type Time struct {
	val time.Time
	mu  sync.RWMutex
}

func NewTime(initialValue time.Time) *Time {
	return &Time{val: initialValue}
}

// Alter modifies the current value with the given function, atomically.
func (at *Time) Alter(alterFn func(time.Time) time.Time) time.Time {
	at.mu.Lock()
	defer at.mu.Unlock()

	at.val = alterFn(at.val)

	return at.val
}

// Val safely returns the stored value.
func (at *Time) Val() time.Time {
	at.mu.RLock()
	defer at.mu.RUnlock()

	return at.val
}

// Set safely sets the stored value.
func (at *Time) Set(newValue time.Time) {
	at.Alter(func(_ time.Time) time.Time { return newValue })
}
