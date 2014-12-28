package example

import (
	"fmt"
)

// Example is the struct that will hold optional values.
//go:generate optioner -type Example
type Example struct {
	N      int
	FSlice []float64 `opt:"myName,omitempty"`
	Map    map[string]int
	Name   string
	ff     func(int) int
}

// NewExample creates an example.
// name is required.
func NewExample(name string, options ...Option) *Example {

	// Set required values and initialize optional fields with default values.
	ex := &Example{
		Name:   name,
		N:      10,
		FSlice: make([]float64, 0, 100),
		Map:    make(map[string]int),
	}

	// Set options.
	ex.init(options...)

	fmt.Printf("Example initalized: %+v", ex)
	return ex
}

// Options

// Option type is used to pass options to Example.
type Option func(*Example)

// N sets field N in Example.
func N(n int) Option {
	return func(ex *Example) {
		ex.N = n
	}
}

// FSlice sets field FSlice in Example.
func FSlice(f []float64) Option {
	return func(ex *Example) {
		ex.FSlice = f
	}
}

// Map sets field Map in Example.
func Map(m map[string]int) Option {
	return func(ex *Example) {
		ex.Map = m
	}
}

// initExample applies options to Example.
func (ex *Example) init(options ...Option) {
	for _, option := range options {
		option(ex)
	}
}

type ZZZ struct {
	f1 string
	f2 uint64
}
