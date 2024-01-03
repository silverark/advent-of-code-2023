package part2

import (
	"fmt"
	"log"
	"silverark/aoc-2023/pkg/shared"
	"sort"
	"strings"
)

var workflows map[string]*Workflow

type Workflow struct {
	name  string
	rules []Rule
}

// Process returns true if accepted, or false if rejected
func (w *Workflow) Process(p *Part) bool {

	//log.Println("Processing", w.name, p)
	for _, r := range w.rules {
		//log.Println("Checking Rule", r)
		if r.Op == "" {
			//log.Println("No Op", r)
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
				//log.Println("Rule Passed")
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
				//log.Println("Rule Passed")
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

func process(input []string) uint64 {

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

	// Funk some calcs. After failing miserably I took inspiration from https://github.com/macos-fuse-t/aoc/blob/main/2023/19/main.go
	dmin := map[string][]int{}
	dmax := map[string][]int{}
	for _, w := range workflows {
		for _, r := range w.rules {
			if r.Op == ">" {
				dmin[r.Category] = append(dmin[r.Category], r.Value+1)
				dmax[r.Category] = append(dmax[r.Category], r.Value)
			} else if r.Op == "<" {
				dmin[r.Category] = append(dmin[r.Category], r.Value)
				dmax[r.Category] = append(dmax[r.Category], r.Value-1)
			}
		}
	}
	for p := range dmin {
		dmin[p] = append(dmin[p], 1)
		sort.Slice(dmin[p], func(i, j int) bool {
			return dmin[p][i] < dmin[p][j]
		})
	}
	for p := range dmax {
		dmax[p] = append(dmax[p], 4000)
		sort.Slice(dmax[p], func(i, j int) bool {
			return dmax[p][i] < dmax[p][j]
		})
	}

	n := uint64(0)
	for xi, x := range dmin["x"] {
		log.Println("Processing X", x)
		for mi, m := range dmin["m"] {
			for si, s := range dmin["s"] {
				for ai, a := range dmin["a"] {
					p := Part{X: x, M: m, A: a, S: s}
					workflowStart := workflows["in"]
					accepted := workflowStart.Process(&p)

					if accepted {
						n += (uint64(dmax["a"][ai]) - uint64(a) + 1) *
							(uint64(dmax["s"][si]) - uint64(s) + 1) *
							(uint64(dmax["m"][mi]) - uint64(m) + 1) *
							(uint64(dmax["x"][xi]) - uint64(x) + 1)
					}
				}
			}
		}
	}

	return n
}
