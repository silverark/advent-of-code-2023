package part2

import (
	"github.com/stretchr/testify/assert"
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))

	expect := 5905
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)
}

func TestTypes(t *testing.T) {

	process := func(cards string) *hand {
		h := hand{Cards: cards}
		return &h
	}

	// Stupid Bug Somewhere, let's test the J's
	assert.Equal(t, TypeFiveOfAKind, process("AAAAA").Type())
	assert.Equal(t, TypeFiveOfAKind, process("JAAAA").Type())
	assert.Equal(t, TypeFiveOfAKind, process("JJAAA").Type())
	assert.Equal(t, TypeFiveOfAKind, process("JJJAA").Type())
	assert.Equal(t, TypeFiveOfAKind, process("JJJJA").Type())
	assert.Equal(t, TypeFiveOfAKind, process("JJJJJ").Type())
	assert.Equal(t, TypeFullHouse, process("AAJBB").Type())
	assert.Equal(t, TypeFourOfAKind, process("ABJJJ").Type())
	assert.Equal(t, TypeFourOfAKind, process("JJJAB").Type())
	assert.Equal(t, TypeThreeOfAKind, process("JJABC").Type())
	assert.Equal(t, TypeThreeOfAKind, process("JAABC").Type())
	assert.Equal(t, TypeOnePair, process("JABCD").Type())

}
