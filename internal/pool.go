package internal

import "sync"

type Pool[T any] struct {
	pool sync.Pool
	New  func() *T
}

func (p *Pool[T]) Get() *T {
	if x := p.pool.Get(); x != nil {
		return x.(*T)
	}

	return p.New()
}

func (p *Pool[T]) Put(x *T) {
	p.pool.Put(x)
}
