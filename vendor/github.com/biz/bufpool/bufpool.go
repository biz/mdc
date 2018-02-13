// Package bufpool is a thin wrapper around
// sync.Pool that stores *bytes.Buffer
package bufpool

import (
	"bytes"
	"sync"
)

type Pool struct {
	sync.Pool
}

func (p *Pool) Get() *bytes.Buffer {
	return p.Pool.Get().(*bytes.Buffer)
}

func (p *Pool) Put(b *bytes.Buffer) {
	b.Reset()
	p.Pool.Put(b)
}

func New() *Pool {
	p := &Pool{}

	p.Pool.New = func() interface{} {
		return &bytes.Buffer{}
	}

	return p
}
