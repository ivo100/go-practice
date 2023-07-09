package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	l := New[int]()
	assert.Equal(t, 0, l.Len())
	l.Append(1)
	assert.Equal(t, 1, l.Len())
	l.Append(2)
	assert.Equal(t, 2, l.Len())
	v := *l.First()
	assert.Equal(t, 1, v)
	v = *l.Next()
	assert.Equal(t, 2, v)
	assert.Nil(t, l.Next())
}

func TestRemove(t *testing.T) {
	l := New[int]()
	l.Append(1)
	assert.Equal(t, 1, l.Len())
	v := l.Remove()
	assert.Equal(t, 1, *v)
	assert.Equal(t, 0, l.Len())
}
