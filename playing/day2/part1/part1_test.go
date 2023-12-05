package part1

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {

	gameLimit := ColourCount{red: 12, green: 13, blue: 14}
	value := Process(file.GetFile("../input_test.txt"), gameLimit)

	expect := 8
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = Process(file.GetFile("../input.txt"), gameLimit)
	log.Println("The answer is", value)
}
