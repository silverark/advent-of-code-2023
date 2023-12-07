package part2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	TypeFiveOfAKind handType = iota + 1
	TypeFourOfAKind
	TypeFullHouse
	TypeThreeOfAKind
	TypeTwoPairs
	TypeOnePair
	TypeHighCard
)

var cardVal = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type hand struct {
	Cards    string
	Freq     []cardFreq
	CharMap  map[byte]int
	bid      int
	handType handType
}

func (h *hand) Type() handType {
	if h.handType != 0 {
		return h.handType
	}
	h.CalcFreq()
	if len(h.Freq) == 1 || h.Freq[0].frequency+h.CharMap['J'] == 5 { // 5 of a kind
		h.handType = TypeFiveOfAKind
		return h.handType
	}
	if h.Freq[0].frequency == 4 || h.Freq[0].frequency+h.CharMap['J'] == 4 { // 4 of a kind
		h.handType = TypeFourOfAKind
		return h.handType
	}
	if len(h.Freq) == 2 || h.Freq[0].frequency+h.Freq[1].frequency+h.CharMap['J'] == 5 { // Full house (just needs two types)
		h.handType = TypeFullHouse
		return h.handType
	}
	if h.Freq[0].frequency == 3 || h.Freq[0].frequency+h.CharMap['J'] == 3 { // 3 of a kind
		h.handType = TypeThreeOfAKind
		return h.handType
	}
	if h.Freq[0].frequency == 2 && h.Freq[1].frequency == 2 { // 2 pairs
		h.handType = TypeTwoPairs
		return h.handType
	}
	if h.Freq[0].frequency == 2 || h.CharMap['J'] > 0 { // 1 pair
		h.handType = TypeOnePair
		return h.handType
	}
	h.handType = TypeHighCard
	return h.handType
}

func (h *hand) CalcFreq() {
	stringMap := make(map[byte]int)
	for i := 0; i < len(h.Cards); i++ {
		stringMap[h.Cards[i]]++
	}
	h.CharMap = stringMap
	for key, value := range stringMap {
		h.Freq = append(h.Freq, cardFreq{
			char:      key,
			frequency: value,
		})
	}
	sort.Slice(h.Freq, func(i, j int) bool {
		// Push J to the right in terms of frequency
		if h.Freq[i].char == 'J' {
			return false
		}
		if h.Freq[j].char == 'J' {
			return true
		}
		return h.Freq[i].frequency > h.Freq[j].frequency
	})
}

type cardFreq struct {
	char      byte
	frequency int
}

type handType int

func process(input []string) int {
	var hands []hand
	for _, line := range input {
		items := strings.Fields(line)
		bid, _ := strconv.Atoi(items[1])
		hands = append(hands, hand{
			Cards: items[0],
			bid:   bid,
		})
	}
	// Sort the slice
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type() == hands[j].Type() { // If they are the same hand, then compare highest card.
			for card := 0; card < len(hands[i].Cards); card++ {
				if hands[i].Cards[card] == hands[j].Cards[card] { // Same card skip
					continue
				}
				return cardVal[hands[i].Cards[card]] < cardVal[hands[j].Cards[card]] // Check value of card from map
			}
		}
		return hands[i].Type() > hands[j].Type()
	})
	total := 0
	for i, h := range hands {
		fmt.Println("Hand: ", h.Cards, "\tBid: ", h.bid, "\tType: ", h.Type(), "\tRank: ", i+1, "\tFreq: ", h.Freq)
		total += h.bid * (i + 1)
	}
	return total
}
