package sampler

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterator(t *testing.T) {
	s := []interface{}{0, 1, "two", 3, "vier"}
	it := FromSlice(s)
	for _, expected := range s {
		v, err := it.Next()
		assert.Equal(t, expected, v)
		assert.NoError(t, err)
	}
	v, err := it.Next()
	assert.Nil(t, v)
	if assert.Error(t, err) {
		assert.Equal(t, io.EOF, err)
	}
}

func TestChanIterator(t *testing.T) {
	s := []interface{}{0, 1, "two", 3, "vier"}
	ch := make(chan interface{})
	go func() {
		for _, v := range s {
			ch <- v
		}
		close(ch)
	}()

	it := FromCh(ch)
	for _, expected := range s {
		v, err := it.Next()
		assert.Equal(t, expected, v)
		assert.NoError(t, err)
	}
	v, err := it.Next()
	assert.Nil(t, v)
	if assert.Error(t, err) {
		assert.Equal(t, io.EOF, err)
	}
}
