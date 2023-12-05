package part1

import (
	"github.com/stretchr/testify/assert"
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))

	expect := 142
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = process(file.GetFile("../input1.txt"))
	log.Println("The answer is", value)
}

func TestDay1_FirstNumber(t *testing.T) {
	assert.Equal(t, "1", FirstNumber("XXX1XXXX9"))
	assert.Equal(t, "1", FirstNumber("1XXXXXXX9"))
	assert.Equal(t, "1", FirstNumber("XXXXXXXX1"))
}

func TestDay1_LastNumber(t *testing.T) {
	assert.Equal(t, "9", LastNumber("XXX1XXXX9"))
	assert.Equal(t, "9", LastNumber("1XXXX9XXX"))
	assert.Equal(t, "1", LastNumber("XXXXXXXX1"))
}
