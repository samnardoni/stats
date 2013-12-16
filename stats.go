package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Tally struct {
	count int
	sum   float64
	min   float64
	max   float64
}

func NewTally() Tally {
	return Tally{
		count: 0,
		sum:   0,
		min:   math.MaxFloat64,
		max:   0, // should be min float?
	}
}

func (t *Tally) Step(v float64) {
	t.count += 1
	t.sum += v
	t.min = math.Min(t.min, v)
	t.max = math.Max(t.max, v)
}

func (t *Tally) Average() float64 { return t.sum / float64(t.count) }
func (t *Tally) Min() float64     { return t.min }
func (t *Tally) Max() float64     { return t.max }

// Scans for line-separated floats
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

func PrintSummary(w io.Writer, s map[string]float64) {
	for name, stat := range s {
		fmt.Printf("%-7s %g\n", name, stat)
	}
}

func main() {
	tally := NewTally()

	for scanner := NewFloatScanner(os.Stdin); scanner.Scan(); {
		tally.Step(scanner.Value())
	}

	PrintSummary(os.Stdout, map[string]float64{
		"Average": tally.Average(),
		"Max":     tally.Max(),
		"Min":     tally.Min(),
	})
}
