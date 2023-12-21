package part2

import (
	"slices"
	"strings"
)

var modules map[string]Module

type Pulse struct {
	pulseType bool // High or Low, 0 or 1
	from      string
	to        string
}

type Module interface {
	ProcessPulse(pulse *Pulse) []*Pulse
	GetDestinations() []string
	GetName() string
}

type ModuleBase struct {
	name         string
	destinations []string
}

func (m *ModuleBase) GetDestinations() []string {
	return m.destinations
}
func (m *ModuleBase) GetName() string {
	return m.name
}

type ModuleFlipFlop struct {
	ModuleBase
	state bool
}

func (m *ModuleFlipFlop) ProcessPulse(pulse *Pulse) []*Pulse {
	var pulses []*Pulse
	if pulse.pulseType == true {
		return pulses
	}
	newPulse := false
	if m.state == true {
		m.state = false
	} else {
		m.state = true
		newPulse = true
	}
	for _, module := range m.destinations {
		pulses = append(pulses, &Pulse{
			pulseType: newPulse,
			from:      m.name,
			to:        module,
		})
	}

	return pulses
}

type ModuleConjunction struct {
	ModuleBase
	LastReceived map[string]bool
	Cycle        int
}

func (m *ModuleConjunction) ProcessPulse(pulse *Pulse) []*Pulse {
	var pulses []*Pulse
	if m.LastReceived == nil {
		m.LastReceived = make(map[string]bool)
		// Add each input
		for _, module := range modules {
			if slices.Contains(module.GetDestinations(), m.name) {
				m.LastReceived[module.GetName()] = false
			}
		}
	}
	m.LastReceived[pulse.from] = pulse.pulseType

	allHigh := true
	for _, value := range m.LastReceived {
		if value == false {
			allHigh = false
			break
		}
	}
	newPulse := true
	if allHigh == true {
		newPulse = false
	}
	for _, module := range m.destinations {
		pulses = append(pulses, &Pulse{
			pulseType: newPulse,
			from:      m.name,
			to:        module,
		})
	}
	return pulses
}

type ModuleBroadcast struct {
	ModuleBase
}

func (m *ModuleBroadcast) ProcessPulse(pulse *Pulse) []*Pulse {
	var pulses []*Pulse
	for _, module := range m.destinations {
		pulses = append(pulses, &Pulse{
			pulseType: pulse.pulseType,
			from:      m.name,
			to:        module,
		})
	}
	return pulses
}

func process(input []string) int {

	var moduleVd *ModuleConjunction
	rxDownstream := make(map[string]*ModuleConjunction)

	modules = make(map[string]Module)
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		name := strings.TrimSpace(parts[0][1:len(parts[0])])
		switch parts[0][0] {
		case '%': // FlipFlop
			newModule := &ModuleFlipFlop{}
			newModule.name = name
			newModule.destinations = getDestinations(parts[1])
			modules[newModule.name] = newModule
		case '&': // Conjunction
			newModule := &ModuleConjunction{}
			newModule.name = name
			newModule.destinations = getDestinations(parts[1])
			modules[newModule.name] = newModule

			// Set VD to keep track of the cycle, and it's downstream (upstream?) modules
			if name == "vd" {
				moduleVd = newModule
			}
			if slices.Contains(newModule.destinations, "vd") {
				rxDownstream[newModule.name] = newModule
			}

		case 'b': // Broadcast
			newModule := &ModuleBroadcast{}
			newModule.name = "broadcaster"
			newModule.destinations = getDestinations(parts[1])
			modules[newModule.name] = newModule
		}
	}

	i := 0
	for {
		var pulses []*Pulse

		// Push button
		pulses = append(pulses, &Pulse{
			pulseType: false,
			from:      "button",
			to:        "broadcaster",
		})

		// Process all the pulses.
		for j := 0; j < len(pulses); j++ {
			module, found := modules[pulses[j].to]
			if found == false {
				continue
			}

			// See if the module is downstream of rx
			if pulses[j].pulseType == true {
				foundAll := true
				for _, mod := range rxDownstream {
					//log.Println("rx downstream", mod.GetName())
					if mod.Cycle == 0 && moduleVd.LastReceived[mod.name] == true {
						mod.Cycle = i + 1
					}
					if mod.Cycle == 0 {
						foundAll = false
					}
				}
				if foundAll == true {
					cycles := make([]int, 0)
					for _, mod := range rxDownstream {
						cycles = append(cycles, mod.Cycle)
					}
					totalButtonPresses := LCM(cycles[0], cycles[1], cycles[2:]...)
					return totalButtonPresses
				}
			}
			newPulses := module.ProcessPulse(pulses[j])
			pulses = append(pulses, newPulses...)
		}

		i++
	}

	return 0
}

func getDestinations(input string) []string {
	fields := strings.FieldsFunc(input, func(r rune) bool {
		return r == ','
	})
	for i := 0; i < len(fields); i++ {
		fields[i] = strings.TrimSpace(fields[i])
	}
	return fields
}

// GCD (greatest common divisor) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM (Least Common Multiple) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
