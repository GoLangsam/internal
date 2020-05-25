package snip // adapted from "flag/flag.go"

import (
	"fmt"
	"strconv"
	"time"
)

// -- bool Value
type Bool bool

func NewBool(val bool, p *bool) *Bool {
	*p = val
	return (*Bool)(p)
}

func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = Bool(v)
	return err
}

func (b *Bool) Get() interface{} { return bool(*b) }

func (b *Bool) String() string { return fmt.Sprintf("%v", *b) }

// -- int Value
type Int int

func NewInt(val int, p *int) *Int {
	*p = val
	return (*Int)(p)
}

func (i *Int) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = Int(v)
	return err
}

func (i *Int) Get() interface{} { return int(*i) }

func (i *Int) String() string { return fmt.Sprintf("%v", *i) }

// -- int64 Value
type Int64 int64

func NewInt64(val int64, p *int64) *Int64 {
	*p = val
	return (*Int64)(p)
}

func (i *Int64) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = Int64(v)
	return err
}

func (i *Int64) Get() interface{} { return int64(*i) }

func (i *Int64) String() string { return fmt.Sprintf("%v", *i) }

// -- uint Value
type Uint uint

func NewUint(val uint, p *uint) *Uint {
	*p = val
	return (*Uint)(p)
}

func (i *Uint) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = Uint(v)
	return err
}

func (i *Uint) Get() interface{} { return uint(*i) }

func (i *Uint) String() string { return fmt.Sprintf("%v", *i) }

// -- uint64 Value
type Uint64 uint64

func NewUint64(val uint64, p *uint64) *Uint64 {
	*p = val
	return (*Uint64)(p)
}

func (i *Uint64) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = Uint64(v)
	return err
}

func (i *Uint64) Get() interface{} { return uint64(*i) }

func (i *Uint64) String() string { return fmt.Sprintf("%v", *i) }

// -- string Value
type String string

func NewString(val string, p *string) *String {
	*p = val
	return (*String)(p)
}

func (s *String) Set(val string) error {
	*s = String(val)
	return nil
}

func (s *String) Get() interface{} { return string(*s) }

func (s *String) String() string { return fmt.Sprintf("%s", *s) }

// -- float64 Value
type Float float64

func NewFloat(val float64, p *float64) *Float {
	*p = val
	return (*Float)(p)
}

func (f *Float) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = Float(v)
	return err
}

func (f *Float) Get() interface{} { return float64(*f) }

func (f *Float) String() string { return fmt.Sprintf("%v", *f) }

// -- time.Duration Value
type Duration time.Duration

func NewDuration(val time.Duration, p *time.Duration) *Duration {
	*p = val
	return (*Duration)(p)
}

func (d *Duration) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = Duration(v)
	return err
}

func (d *Duration) Get() interface{} { return time.Duration(*d) }

func (d *Duration) String() string { return (*time.Duration)(d).String() }

// Value is the interface to a dynamic value.
//
// Hint: All implentations provide a New function, which takes two arguments
// - an initial default content
// - a go pointer variable (typially created by var ...) to hold and access the content
//
// Note: All implentations should also provide a Let method which receives the represented type.
type Value interface {
	String() string
	Set(string) error
	Get() interface{}
}

// assure interface compliance

var _ Value = NewBool(true, new(bool))
var _ Value = NewInt(0, new(int))
var _ Value = NewInt64(0, new(int64))
var _ Value = NewUint(0, new(uint))
var _ Value = NewUint64(0, new(uint64))
var _ Value = NewFloat(0, new(float64))

func dur(duration string) time.Duration {
	dur, _ := time.ParseDuration(duration)
	return dur
}

var _ Value = NewDuration(dur("10h20m30s40ms50us60ns"), new(time.Duration))
