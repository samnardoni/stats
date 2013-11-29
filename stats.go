package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Stats interface {
	Step(float64)
	Result() float64
}

type Mean struct {
	sum   float64
	count int
}

func NewMean() *Mean               { return &Mean{} }
func (m *Mean) Step(value float64) { m.sum += value; m.count++ }
func (m *Mean) Result() float64    { return m.sum / float64(m.count) }

type Min struct{ min float64 }

func NewMin() *Min                { return &Min{math.MaxFloat64} }
func (m *Min) Step(value float64) { m.min = math.Min(m.min, value) }
func (m *Min) Result() float64    { return m.min }

type Max struct{ max float64 }

func NewMax() *Max                { return &Max{0} }
func (m *Max) Step(value float64) { m.max = math.Max(m.max, value) }
func (m *Max) Result() float64    { return m.max }

type FloatScanner struct {
	scanner *bufio.Scanner
	value   float64
}

func NewFloatScanner(r io.Reader) *FloatScanner {
	return &FloatScanner{scanner: bufio.NewScanner(r)}
}

func (f *FloatScanner) Scan() bool {
	for f.scanner.Scan() {
		value, err := strconv.ParseFloat(f.scanner.Text(), 64)
		if err != nil {
			continue
		}
		f.value = value
		return true
	}
	return false
}

func (f *FloatScanner) Value() float64 {
	return f.value
}

func main() {
	stats := map[string]Stats{
		"Average": NewMean(),
		"Max":     NewMax(),
		"Min":     NewMin(),
	}

	for scanner := NewFloatScanner(os.Stdin); scanner.Scan(); {
		for _, s := range stats {
			s.Step(scanner.Value())
		}
	}

	for name, s := range stats {
		fmt.Printf("%-7s %g\n", name, s.Result())
	}
}
