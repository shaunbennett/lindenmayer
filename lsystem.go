package main

import (
	"math/rand"
	"strings"
	"time"
)

type LSystem struct {
	axiom   string
	rules   map[rune]*production
	outputs map[rune]func()
	rng     *rand.Rand
}

type successor struct {
	weight int
	value  string
}

type production struct {
	weightSum  int
	successors []successor
}

func New(axiom string) *LSystem {
	return &LSystem{
		axiom:   axiom,
		rules:   make(map[rune]*production),
		outputs: make(map[rune]func()),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (p *production) gen(rng *rand.Rand) string {
	// If only one successor, just use it
	if len(p.successors) == 1 {
		return p.successors[0].value
	}

	randValue := rng.Intn(len(p.successors))
	currentWeight := 0
	for _, s := range p.successors {
		currentWeight += s.weight
		if randValue < currentWeight {
			return s.value
		}
	}
	return p.successors[len(p.successors)-1].value
}

func (l *LSystem) AddProduction(c rune, value string) {
	l.AddWeightedProduction(c, value, 1)
}

func (l *LSystem) AddWeightedProduction(c rune, value string, weight int) {
	prod, ok := l.rules[c]
	if !ok {
		prod = &production{}
		l.rules[c] = prod
	}

	prod.successors = append(prod.successors, successor{weight, value})
	prod.weightSum += weight
}

func (l *LSystem) AddOutput(c rune, output func()) {
	l.outputs[c] = output
}

func (l *LSystem) Iterate(n int) string {
	var sb strings.Builder
	previousString := l.axiom

	for i := 0; i < n; i++ {
		for _, c := range previousString {
			rule, ok := l.rules[c]
			if ok {
				sb.WriteString(rule.gen(l.rng))
			} else {
				sb.WriteRune(c)
			}
		}
		previousString = sb.String()
		sb.Reset()
	}

	// Process outputs
	for _, c := range previousString {
		output, ok := l.outputs[c]
		if ok {
			output()
		}
	}

	return previousString
}

func (l *LSystem) IterateOnce() string {
	return l.Iterate(1)
}
