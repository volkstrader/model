package model

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfile_Dating(t *testing.T) {
	a := Profile{
		Name: gofakeit.Name(),
		Age:  25,
	}

	assert.True(t, a.Dating())
	a.Age = 26
	assert.False(t, a.Dating())
}
