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
// Int64
//-----------------------------------------------------------------------------

// Int64 stores an int64 value and allows it to be read and modified atomically.
type Int64 struct {
	val int64
}

func NewInt64(initialValue int64) *Int64 {
	return &Int64{initialValue}
}

// Add atomically adds the given delta to the stored value.
func (ai *Int64) Add(delta int64) int64 {
	return atomic.AddInt64(&ai.val, delta)
}

// Val safely returns the stored value.
func (ai *Int64) Val() int64 {
	return atomic.LoadInt64(&ai.val)
}

// Set safely sets the stored value.
func (ai *Int64) Set(newValue int64) {
	atomic.StoreInt64(&ai.val, newValue)
}

// Swap safely swaps the stored value and returns the old value
func (ai *Int64) Swap(newValue int64) int64 {
	return atomic.SwapInt64(&ai.val, newValue)
}

//-----------------------------------------------------------------------------
// String
//-----------------------------------------------------------------------------

// String stores a string and allows it to be read and modified atomically.
type String struct {
	val string
	mu  sync.RWMutex
}

func NewString(initialValue string) *String {
	return &String{val: initialValue}
}

// Set safely sets the stored value.
func (at *String) Set(newValue string) {
	at.mu.Lock()
	defer at.mu.Unlock()

	at.val = newValue
}

// Val safely returns the stored value.
func (at *String) Val() string {
	at.mu.RLock()
	defer at.mu.RUnlock()

	return at.val
}

// Swap safely swaps the stored value and returns the old value
func (at *String) Swap(newValue string) string {
	at.mu.Lock()
	defer at.mu.Unlock()

	oldVal := at.val
	at.val = newValue

	return oldVal
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

// Swap safely swaps the stored value and returns the old value
func (at *Time) Swap(newValue time.Time) time.Time {
	at.mu.Lock()
	defer at.mu.Unlock()

	oldVal := at.val
	at.val = newValue

	return oldVal
}

//-----------------------------------------------------------------------------
// Duration
//-----------------------------------------------------------------------------

// Duration stores a time.Duration and allows it to be read and modified atomically.
type Duration struct {
	val time.Duration
	mu  sync.RWMutex
}

func NewDuration(initialValue time.Duration) *Duration {
	return &Duration{val: initialValue}
}

// Set safely sets the stored value.
func (at *Duration) Set(newValue time.Duration) {
	at.mu.Lock()
	defer at.mu.Unlock()

	at.val = newValue
}

// Val safely returns the stored value.
func (at *Duration) Val() time.Duration {
	at.mu.RLock()
	defer at.mu.RUnlock()

	return at.val
}

// Swap safely swaps the stored value and returns the old value
func (at *Duration) Swap(newValue time.Duration) time.Duration {
	at.mu.Lock()
	defer at.mu.Unlock()

	oldVal := at.val
	at.val = newValue

	return oldVal
}

//-----------------------------------------------------------------------------
// Bool
//-----------------------------------------------------------------------------

// Bool stores a boolean that can be changed atomically, and accessed in a threadsafe way.
type Bool struct {
	val bool
	mu  sync.RWMutex
}

func NewBool(initialValue bool) *Bool {
	return &Bool{val: initialValue}
}

// Set safely sets the stored value.
func (ab *Bool) Set(newValue bool) {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	ab.val = newValue
}

// SetAtomically holds the lock while the callback determines the new value to set.
func (ab *Bool) SetAtomically(callback func(oldValue bool) (newValue bool)) {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	newValue := callback(ab.val)

	ab.val = newValue
}

// Val safely returns the stored value.
func (ab *Bool) Val() bool {
	ab.mu.RLock()
	defer ab.mu.RUnlock()

	return ab.val
}

// ValWithCallback holds the readlock while the callback executes.
func (ab *Bool) ValWithCallback(callback func(curVal bool) error) error {
	ab.mu.RLock()
	defer ab.mu.RUnlock()

	return callback(ab.val)
}

// Swap safely swaps the stored value and returns the old value
func (ab *Bool) Swap(newValue bool) bool {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	oldVal := ab.val
	ab.val = newValue

	return oldVal
}

// Allow a read-only Bool to exist
type BoolReadonly interface {
	Val() bool
}

var _ BoolReadonly = &Bool{}

//-----------------------------------------------------------------------------
// Map
//-----------------------------------------------------------------------------

// Map stores a map that can be changed atomically, and accessed in a threadsafe way.
type Map struct {
	mu   sync.RWMutex
	data map[interface{}]interface{}
}

func NewMap() *Map {
	return &Map{data: map[interface{}]interface{}{}}
}

// Get safely returns the stored value.
func (am *Map) Get(key interface{}) interface{} {
	am.mu.RLock()
	defer am.mu.RUnlock()

	return am.data[key]
}

func (am *Map) Len() int {
	am.mu.RLock()
	defer am.mu.RUnlock()

	return len(am.data)
}

func (am *Map) Values() []interface{} {
	am.mu.RLock()
	defer am.mu.RUnlock()

	vals := make([]interface{}, 0, len(am.data))
	for _, v := range am.data {
		vals = append(vals, v)
	}
	return vals
}

// Set safely sets the stored value.
func (am *Map) Set(key interface{}, val interface{}) {
	am.mu.Lock()
	defer am.mu.Unlock()

	am.data[key] = val
}

// Set safely sets the stored value.
func (am *Map) Del(key interface{}) {
	am.mu.Lock()
	defer am.mu.Unlock()

	delete(am.data, key)
}
