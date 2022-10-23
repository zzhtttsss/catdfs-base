package util

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

type Int int

type String string

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (s String) String() string {
	return string(s)
}

func TestNewQueue(t *testing.T) {
	queue := NewQueue[Int]()
	assert.NotNil(t, queue)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 3, queue.Len())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Pop())
}

func TestQueue_Pop(t *testing.T) {
	queue := Queue[String]{
		data: []String{"Hello", "World", "HuHu"},
		mu:   &sync.RWMutex{},
	}
	assert.Equal(t, "Hello", queue.Pop())
}

func TestQueue_Push(t *testing.T) {
	queue := Queue[String]{
		data: []String{},
		mu:   &sync.RWMutex{},
	}
	queue.Push("Hello")
	queue.Push("World")
	assert.Equal(t, 2, len(queue.data))
	assert.Equal(t, "Hello", queue.Pop())
}
