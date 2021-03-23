package methods

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircle(t *testing.T) {
	parameters := []struct {
		r, expectedDiameter, expectedLength float64
	}{
		{123.4, 246.8, 775.345066905961},
		{1, 2, 6.283185307179586},
		{0, 0, 0},
	}

	for i := range parameters {
		d := getDiameter(parameters[i].r)
		l := getLength(d)

		assert.Equal(t, d, parameters[i].expectedDiameter)
		assert.Equal(t, l, parameters[i].expectedLength)
	}
}
