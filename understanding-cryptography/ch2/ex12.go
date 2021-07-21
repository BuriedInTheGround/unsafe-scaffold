package main

import (
	"fmt"
)

type bit uint

type lfsr struct {
	feedbackBit    int
	feedforwardBit int
	state          []bit
}

func newLFSR(feedback, feedforward int, iv []bit) *lfsr {
	return &lfsr{
		feedbackBit:    feedback,
		feedforwardBit: feedforward,
		state:          iv,
	}
}

func (r *lfsr) shift(input bit) {
	feedback := r.state[r.feedbackBit]

	for i := len(r.state) - 1; i >= 1; i-- {
		r.state[i] = r.state[i-1]
	}

	r.state[0] = input ^ feedback
}

func (r *lfsr) output() bit {
	feedforward := r.state[r.feedforwardBit]
	and := r.state[len(r.state)-2] & r.state[len(r.state)-3]
	return feedforward ^ and ^ r.state[len(r.state)-1]
}

type trivium struct {
	a *lfsr
	b *lfsr
	c *lfsr
}

func newTrivium(iv, key []bit) *trivium {
	aiv := make([]bit, 93)
	copy(aiv, iv)
	a := newLFSR(69, 66, aiv)

	biv := make([]bit, 84)
	copy(biv, key)
	b := newLFSR(78, 69, biv)

	civ := make([]bit, 111)
	copy(civ[108:], []bit{1, 1, 1})
	c := newLFSR(87, 66, civ)

	return &trivium{
		a: a,
		b: b,
		c: c,
	}
}

func (t *trivium) shift() {
	aOut := t.a.output()
	bOut := t.b.output()
	cOut := t.c.output()
	t.a.shift(cOut)
	t.b.shift(aOut)
	t.c.shift(bOut)
}

func (t *trivium) output() bit {
	return t.a.output() ^ t.b.output() ^ t.c.output()
}

func main() {
	reg := newTrivium(
		[]bit{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		[]bit{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	)

	fmt.Printf("First 70 bits during the warm-up phase of Trivium:\n")
	for i := 0; i < 71; i++ {
		fmt.Printf("%d", reg.output())
		reg.shift()
	}
	fmt.Printf("\n")
}
