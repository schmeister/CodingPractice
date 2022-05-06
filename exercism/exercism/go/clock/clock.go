package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func (c Clock) compute(h, m int) Clock {

	h += c.hour
	m += c.minute

	for m >= 60 {
		m -= 60
		h += 1
	}
	for m < 0 {
		m += 60
		h -= 1
	}

	for h >= 24 {
		h -= 24
	}
	for h < 0 {
		h += 24
	}

	c.minute = m
	c.hour = h
	return c
}
func New(h, m int) Clock {
	c := Clock{}
	c = c.compute(h, m)
	return c
}

func (c Clock) Add(m int) Clock {
	c = c.compute(0, m)
	return c
}

func (c Clock) Subtract(m int) Clock {
	c = c.compute(0, m*-1)
	return c
}

func (c Clock) String() string {
	clock := fmt.Sprintf("%02d:%02d", c.hour, c.minute)
	return clock
}
