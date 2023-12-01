package part2

import (
	"bitbucket.org/silverark/aoc-2023/pkg/file"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input2_test.txt"))

	expect := 281
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
	assert.Equal(t, "8", FirstNumber("eightwothree"))
	assert.Equal(t, "2", FirstNumber("xtwone3four"))
}

func TestDay1_LastNumber(t *testing.T) {
	assert.Equal(t, "9", LastNumber("XXX1XXXX9"))
	assert.Equal(t, "9", LastNumber("1XXXX9XXX"))
	assert.Equal(t, "1", LastNumber("XXXXXXXX1"))
	assert.Equal(t, "3", LastNumber("eightwothree"))
	assert.Equal(t, "4", LastNumber("xtwone3four"))
}
