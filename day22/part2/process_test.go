package part2

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
	"time"
)

func TestProcessTest(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))
	expect := 7
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
}

func TestProcessActual(t *testing.T) {

	fileData := file.GetFile("../input.txt")
	timer := time.Now()
	value := process(fileData)
	log.Println("Time taken:", time.Since(timer))
	if value <= 77911 {
		t.Fatalf("Received %v, but expected a higher value than 77911", value)
	}
	log.Println("The answer is", value)
}
