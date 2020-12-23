package day23

import (
	"container/list"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const cupsToMove = 3

type Cups struct {
	l         *list.List
	cur       *list.Element
	valueToEl map[int]*list.Element
}

func NewCups(in []int) *Cups {
	var (
		l         = list.New()
		valueToEl = make(map[int]*list.Element)
	)

	for _, n := range in {
		valueToEl[n] = l.PushBack(n)
	}

	return &Cups{
		l:         l,
		cur:       l.Front(),
		valueToEl: valueToEl,
	}
}

func (c *Cups) Move() {
	// Since container/list doesn't expose next/prev to us, we have to do 3
	// remove operations. We then add these back later as brand new elements.
	//
	// If possible it would've been nicer to keep the 3 moved elements
	// chained together and simply update references, rather than recreating
	// them.
	var (
		movedValues     = make([]int, cupsToMove)
		invalidDestCups = make(map[int]struct{}, cupsToMove)
	)
	for i := 0; i < cupsToMove; i++ {
		n := c.l.Remove(c.next(c.cur)).(int)
		movedValues[i] = n
		invalidDestCups[n] = struct{}{}
	}

	destCup := c.destinationCup(invalidDestCups)
	for i := len(movedValues) - 1; i >= 0; i-- {
		n := movedValues[i]
		c.valueToEl[n] = c.l.InsertAfter(n, destCup)
	}

	c.cur = c.next(c.cur)
}

func (c *Cups) next(e *list.Element) *list.Element {
	next := e.Next()
	if next == nil {
		return c.l.Front()
	}
	return next
}

func (c *Cups) destinationCup(invalid map[int]struct{}) *list.Element {
	destVal := c.cur.Value.(int) - 1
	for {
		if destVal == 0 {
			destVal = c.l.Len() + cupsToMove
		}
		if _, ok := invalid[destVal]; !ok {
			return c.valueToEl[destVal]
		}
		destVal--
	}
}

func (c *Cups) ResultString() string {
	var b strings.Builder

	e := c.next(c.valueToEl[1])
	for {
		n := e.Value.(int)
		if n == 1 {
			return b.String()
		}

		b.WriteString(strconv.Itoa(n))

		e = c.next(e)
	}
}

func (c *Cups) String() string {
	var b strings.Builder
	for e := c.l.Front(); e != nil; e = e.Next() {
		n := e.Value.(int)
		if e == c.cur {
			_, _ = fmt.Fprintf(&b, "(%d)", n)
		} else {
			_, _ = io.WriteString(&b, strconv.Itoa(n))
		}
		b.WriteRune(' ')
	}
	return b.String()
}
