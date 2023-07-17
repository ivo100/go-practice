package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {

	l1 := Node[int]{
		Value: 5,
	}

	l2 := Node[int]{
		Value: 5,
	}

	r1 := Node[int]{
		Value: 20,
	}

	r2 := Node[int]{
		Value: 20,
	}

	t1 := Node[int]{
		Value: 10,
		Left:  &l1,
		Right: &r1,
	}
	t2 := Node[int]{
		Value: 10,
		Left:  &l2,
		Right: &r2,
	}
	assert.True(t, Equals[int](&t1, &t2))

	r2.Value = 42
	assert.False(t, Equals[int](&t1, &t2))
}
