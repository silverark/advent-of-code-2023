package part1

import (
	"log"
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
	lastReceived map[string]bool
}

func (m *ModuleConjunction) ProcessPulse(pulse *Pulse) []*Pulse {
	var pulses []*Pulse
	if m.lastReceived == nil {
		m.lastReceived = make(map[string]bool)
		// Add each input
		for _, module := range modules {
			if slices.Contains(module.GetDestinations(), m.name) {
				m.lastReceived[module.GetName()] = false
			}
		}
	}
	m.lastReceived[pulse.from] = pulse.pulseType

	allHigh := true
	for _, value := range m.lastReceived {
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

	highCount := 0
	lowCount := 0

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
		case 'b': // Broadcast
			newModule := &ModuleBroadcast{}
			newModule.name = "broadcaster"
			newModule.destinations = getDestinations(parts[1])
			modules[newModule.name] = newModule
		}
	}

	for i := 0; i < 1000; i++ {
		var pulses []*Pulse

		// Push button
		pulses = append(pulses, &Pulse{
			pulseType: false,
			from:      "button",
			to:        "broadcaster",
		})

		// Process all the pulses.
		for i := 0; i < len(pulses); i++ {
			module, found := modules[pulses[i].to]
			if found == false {
				continue
			}
			newPulses := module.ProcessPulse(pulses[i])
			pulses = append(pulses, newPulses...)
		}
		for _, pulse := range pulses {
			if pulse.pulseType == true {
				highCount++
			} else {
				lowCount++
			}
		}

	}

	log.Println("High count", highCount, "Low count", lowCount)

	return highCount * lowCount
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
