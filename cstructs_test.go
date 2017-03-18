package cstructs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A
type A struct {
	Name  string
	Age   int
	About []byte
}

// B
type B struct {
	Name  string
	Age   int
	About []byte
}

// C
type C struct {
	Name   string
	Age    int
	About  []byte
	Gender string
}

// D
type D struct {
	NaMe   string
	AGe    int
	About  []byte
	Gender string
}

func TestCStruct1(t *testing.T) {
	SetStrictMode(true)
	ben := A{
		Name:  "benedith",
		Age:   20,
		About: []byte("Ben is a great sports person"),
	}

	var pete B
	Copy(&ben, &pete)

	assert.Equal(t, ben.Name, pete.Name)
	assert.Equal(t, ben.Age, pete.Age)
	assert.Equal(t, ben.About, pete.About)
}

func TestCStruct2(t *testing.T) {
	SetStrictMode(true)
	ben := A{
		Name:  "benedith",
		Age:   20,
		About: []byte("Ben is a great sports person"),
	}

	var pete C
	Copy(&ben, &pete)

	assert.Equal(t, ben.Name, pete.Name)
	assert.Equal(t, ben.Age, pete.Age)
	assert.Equal(t, ben.About, pete.About)
	assert.Empty(t, pete.Gender)
}

func TestCStruct3StrictModeDisabled(t *testing.T) {
	SetStrictMode(false)
	ben := A{
		Name:  "benedith",
		Age:   20,
		About: []byte("Ben is a great sports person"),
	}

	var pete D
	Copy(&ben, &pete)

	assert.Equal(t, ben.Name, pete.NaMe)
	assert.Equal(t, ben.Age, pete.AGe)
	assert.Equal(t, ben.About, pete.About)
}

func TestCStruct4StrictModeEnabled(t *testing.T) {
	SetStrictMode(true)
	ben := A{
		Name:  "benedith",
		Age:   20,
		About: []byte("Ben is a great sports person"),
	}

	var pete D
	Copy(&ben, &pete)

	assert.NotEqual(t, ben.Name, pete.NaMe)
	assert.NotEqual(t, ben.Age, pete.AGe)
	assert.Equal(t, ben.About, pete.About)
}
