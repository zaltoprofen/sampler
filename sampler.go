package sampler

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

var defaultSampler *Sampler

func init() {
	defaultSampler = New(rand.NewSource(time.Now().UnixNano()))
}

// Iterator iterate values for reservoir sampling
type Iterator interface {
	Next() (value interface{}, err error) // Next should return io.EOF after last value
}

// IteratorFunc is an adapter function for Iterator
type IteratorFunc func() (value interface{}, err error)

// Next calls f()
func (f IteratorFunc) Next() (interface{}, error) {
	return f()
}

// FromCh convert channel of interface to Iterator
func FromCh(source chan interface{}) Iterator {
	retval := func() (interface{}, error) {
		v, ok := <-source
		if ok {
			return v, nil
		}
		return nil, io.EOF
	}
	return IteratorFunc(retval)
}

// FromSlice convert slice to Iterator
func FromSlice(source []interface{}) Iterator {
	i := 0
	f := func() (interface{}, error) {
		if i < len(source) {
			v := source[i]
			i++
			return v, nil
		}
		return nil, io.EOF
	}
	return IteratorFunc(f)
}

// A Sampler will sample
type Sampler struct {
	r *rand.Rand
}

// New returns a Sampler that use given random source
func New(randSource rand.Source) *Sampler {
	return &Sampler{rand.New(randSource)}
}

// Sample returns reservoir sampled values with defaultSampler
func Sample(k int, source Iterator) ([]interface{}, error) {
	return defaultSampler.Sample(k, source)
}

// Sample returns reservoir sampled values
func (s *Sampler) Sample(k int, source Iterator) ([]interface{}, error) {
	retval := make([]interface{}, k)
	for i := 0; i < k; i++ {
		var err error
		if retval[i], err = source.Next(); err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("given values fewer than k=%d", k)
			}
			return nil, err
		}
	}
	n := k
	for {
		val, err := source.Next()
		if err == io.EOF {
			return retval, nil
		} else if err != nil {
			return retval, err
		}
		n++
		if i := s.r.Intn(n); i < k {
			retval[i] = val
		}
	}
}
