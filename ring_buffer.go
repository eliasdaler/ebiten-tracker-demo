package main

import (
	"fmt"
	"time"
)

type RingBuffer struct {
	head     int
	tail     int
	capacity int

	buf []byte
}

func NewRingBuffer(capacity int) *RingBuffer {
	if capacity <= 0 {
		panic("invalid buffer capacity")
	}
	return &RingBuffer{
		head:     0,
		tail:     0,
		capacity: capacity + 1,
		buf:      make([]byte, capacity+1),
	}
}

func (rb *RingBuffer) Append(b byte) {
	if (rb.tail+1)%rb.capacity == rb.head {
		panic("buffer overflow")
	}

	rb.buf[rb.tail] = b
	rb.tail = (rb.tail + 1) % rb.capacity
}

func (rb *RingBuffer) Empty() bool {
	return rb.Size() == 0
}

func (rb *RingBuffer) Size() int {
	if rb.tail == rb.head {
		return 0
	}

	if rb.tail > rb.head {
		return rb.tail - rb.head
	}
	return rb.capacity - rb.head + rb.tail
}

func (rb *RingBuffer) Pop() byte {
	if rb.Empty() {
		return 0
	}

	b := rb.buf[rb.head]
	rb.head = (rb.head + 1) % rb.capacity
	return b
}

var lag bool

func (rb *RingBuffer) Read(b []byte) (n int, err error) {
	fmt.Println("READ", len(b))
	if lag {
		time.Sleep(5000 * time.Millisecond)
		fmt.Println("LAG", len(b))
		lag = false
		return 0, nil
	}
	if rb.Empty() {
		for i := 0; i < len(b); i++ {
			b[i] = 0
		}
		return len(b), nil
	}

	for i := 0; i < len(b); i++ {
		b[i] = rb.Pop()
	}

	return len(b), nil
}
