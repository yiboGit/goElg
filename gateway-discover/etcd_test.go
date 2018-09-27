package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetService(t *testing.T) {
	r, err := getService("hello")
	assert.Equal(t, nil, err)
	assert.Equal(t, "", r)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	randIndex := rand.Intn(2)
	assert.Equal(t, 8, randIndex)
}
