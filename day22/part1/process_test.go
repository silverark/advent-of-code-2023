package part1

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcessTest(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))
	expect := 5
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
}

func TestProcessActual(t *testing.T) {
	value := process(file.GetFile("../input.txt"))
	if value >= 1237 {
		log.Fatalf("Received %v, but expected a lower value. Should be lower than 1237", value)
	}
	log.Println("The answer is", value)
}
