package main

import (
	"fmt"
	"math"
	"testing"
)

const epsilon = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestPd(t *testing.T) {
	/*
		p := new(ProbabilityDistribution)
		p.Name = "test"
		p.Labels = []Label{"a", "b", "c"}
		p.Probabilities = []float64{1.0 / 2, 1.0 / 2, 0.0}
	*/
	p := ProbabilityDistribution{"test", []Label{"a", "b", "c"}, []float64{1.0 / 2, 1.0 / 2, 0.0}}
	if !almostEqual(p.Entropy(), math.Log(2.0)) {
		t.Errorf("Incorrect entropy fair coin  %v", p.Entropy())
	}
	fmt.Printf("%v\n", p)
	p.Mark("a")
	fmt.Printf("%v\n", p)
	p.Mark("b")
	fmt.Printf("%v\n", p)
	p.Mark("b")
	fmt.Printf("%v\n", p)
	p.Mark("b")
	fmt.Printf("%v\n", p)
}
