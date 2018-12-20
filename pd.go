package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"

	"github.com/aicodes/sandwich/taxonomy"
)

const (
	//CorrectLikelyhood is the likelyhood that curator can label things correctly
	CorrectLikelyhood = 0.90
	//ErrorRate is the probability that curator may incorrectly label things.
	ErrorRate = 0.10
)

func normalize(p []float64) ([]float64, error) {
	if len(p) == 0 {
		return nil, errors.New("Empty probability distribution")
	}

	newp := make([]float64, len(p))
	sum := 0.0
	for _, v := range p {
		sum += v
	}
	if math.Abs(sum) < 1e-9 {
		return nil, errors.New("Probability underflow")
	}
	for i, v := range p {
		newp[i] = v / sum
	}
	return newp, nil
}

//Label to denote the labels in classification
type Label string

//ProbabilityDistribution represents a distribution of name/label pair
type ProbabilityDistribution struct {
	Name          string
	Labels        []Label
	Probabilities []float64
}

//Entropy of the probability distribution
func (p ProbabilityDistribution) Entropy() float64 {
	s := 0.0
	for _, p := range p.Probabilities {
		if p == 0.0 {
			continue
		}
		s -= p * math.Log(p)
	}
	return s
}

/*
func (p ProbabilityDistribution) Spread() {

}
*/

/*Mark the label and update probability distribution using
Bayesian rule. */
func (p *ProbabilityDistribution) Mark(label Label) {
	// calculate posterior distribution using Bayesian rule.
	posterior := make([]float64, len(p.Probabilities))
	for i, l := range p.Labels {
		if l == label {
			posterior[i] = p.Probabilities[i] * CorrectLikelyhood
		} else {
			posterior[i] = p.Probabilities[i] * ErrorRate
		}
	}
	prob, err := normalize(posterior)
	if err == nil {
		p.Probabilities = prob
	}
}

//LoadDistribution loads the default names and distribution from file.
func LoadDistribution(fname string, tax *taxonomy.Taxonomy) ([]ProbabilityDistribution, error) {
	var ps []ProbabilityDistribution

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	nLabels := len(tax.Root.Children)
	// fix this guy later.
	labels := []Label{"a", "b"}

	uniformDistributin := make([]float64, nLabels)
	for i := range uniformDistributin {
		uniformDistributin[i] = 1.0 / float64(nLabels)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ps = append(ps, ProbabilityDistribution{scanner.Text(), labels, uniformDistributin})
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ps, err
}
