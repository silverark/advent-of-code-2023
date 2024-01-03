package part1

import (
	"fmt"
	"log"
	"silverark/aoc-2023/pkg/shared"
	"strings"
)

var workflows map[string]*Workflow

type Workflow struct {
	name  string
	rules []Rule
}

// Process returns true if accepted, or false if rejected
func (w *Workflow) Process(p *Part) bool {

	log.Println("Processing", w.name, p)
	for _, r := range w.rules {
		log.Println("Checking Rule", r)
		if r.Op == "" {
			log.Println("No Op", r)
			if r.Dest == "A" {
				return true
			}
			if r.Dest == "R" {
				return false
			}
			return workflows[r.Dest].Process(p)
		}
		if r.Op == "<" {
			if p.getVal(r.Category) < r.Value {
				log.Println("Rule Passed")
				if r.Dest == "A" {
					return true
				}
				if r.Dest == "R" {
					return false
				}
				return workflows[r.Dest].Process(p)
			}
		}
		if r.Op == ">" {

			if p.getVal(r.Category) > r.Value {
				log.Println("Rule Passed")
				if r.Dest == "A" {
					return true
				}
				if r.Dest == "R" {
					return false
				}
				return workflows[r.Dest].Process(p)
			}
		}
	}
	return false
}

type Part struct {
	X, M, A, S int
}

func (p *Part) String() string {
	return fmt.Sprintf("X=%v,M=%v,A=%v,S=%v", p.X, p.M, p.A, p.S)
}

func (p *Part) getVal(s string) int {
	switch s {
	case "x":
		return p.X
	case "m":
		return p.M
	case "a":
		return p.A
	case "s":
		return p.S
	}
	panic("Unknown category")
}

type Rule struct {
	Category string
	Op       string
	Value    int
	Dest     string
}

func process(input []string) int {

	workflows = make(map[string]*Workflow)
	var parts []*Part

	for _, row := range input {

		if row == "" {
			continue
		}

		if row[0] == '{' {
			//{X=787,M=2655,A=1222,S=2876}
			items := strings.Split(row[1:len(row)-1], ",")

			parts = append(parts, &Part{
				X: shared.Atoi(strings.Split(items[0], "=")[1]),
				M: shared.Atoi(strings.Split(items[1], "=")[1]),
				A: shared.Atoi(strings.Split(items[2], "=")[1]),
				S: shared.Atoi(strings.Split(items[3], "=")[1]),
			})
		} else {
			//px{A<2006:qkq,M>2090:A,rfg}
			cleanUp := strings.Split(row, "{")
			wf := Workflow{name: cleanUp[0]}
			workflows[cleanUp[0]] = &wf
			rules := strings.ReplaceAll(cleanUp[1], "}", "")
			for _, p := range strings.Split(rules, ",") {

				// Handle single destination Rule (fallthrough)
				if !strings.Contains(p, ":") {
					wf.rules = append(wf.rules, Rule{
						Dest: p,
					})
					continue
				}

				finalParts := strings.Split(p[2:], ":")
				wf.rules = append(wf.rules, Rule{
					Category: string(p[0]),
					Op:       string(p[1]),
					Value:    shared.Atoi(finalParts[0]),
					Dest:     finalParts[1],
				})
			}
		}

	}

	// For each Part pass though the workflows
	totalAccepted := 0
	for _, p := range parts {
		log.Println("\n\nProcessing Part", p)
		workflowStart := workflows["in"]
		accepted := workflowStart.Process(p)

		if accepted {
			log.Println("Accepted", p)
			totalAccepted += p.X + p.M + p.A + p.S
		} else {
			log.Println("Rejected", p)
		}
	}

	return totalAccepted
}
