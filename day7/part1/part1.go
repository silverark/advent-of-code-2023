package part1

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
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type hand struct {
	cards    string
	freq     []cardFreq
	bid      int
	handType handType
}

func (h *hand) Type() handType {
	if h.handType != 0 {
		return h.handType
	}
	h.CalcFreq()
	if len(h.freq) == 1 { // 5 of a kind
		h.handType = TypeFiveOfAKind
		return h.handType
	}
	if h.freq[0].frequency == 4 { // 4 of a kind
		h.handType = TypeFourOfAKind
		return h.handType
	}
	if len(h.freq) == 2 { // Full house (just needs two types)
		h.handType = TypeFullHouse
		return h.handType
	}
	if h.freq[0].frequency == 3 { // 3 of a kind (123 or )
		h.handType = TypeThreeOfAKind
		return h.handType
	}
	if h.freq[0].frequency == 2 && h.freq[1].frequency == 2 { // 2 pairs
		h.handType = TypeTwoPairs
		return h.handType
	}
	if h.freq[0].frequency == 2 { // 1 pair
		h.handType = TypeOnePair
		return h.handType
	}
	h.handType = TypeHighCard
	return h.handType
}
func (h *hand) CalcFreq() {
	stringMap := make(map[byte]int)
	for i := 0; i < len(h.cards); i++ {
		stringMap[h.cards[i]]++
	}
	for key, value := range stringMap {
		i := cardFreq{
			char:      key,
			frequency: value,
		}
		h.freq = append(h.freq, i)
	}
	sort.Slice(h.freq, func(i, j int) bool {
		return h.freq[i].frequency > h.freq[j].frequency
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
			cards: items[0],
			bid:   bid,
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		// If they are the same hand, then compare highest card.
		if hands[i].Type() == hands[j].Type() {
			for card := 0; card < len(hands[i].cards); card++ {
				if cardVal[hands[i].cards[card]] == cardVal[hands[j].cards[card]] {
					continue
				}
				return cardVal[hands[i].cards[card]] < cardVal[hands[j].cards[card]]
			}
		}

		return hands[i].Type() > hands[j].Type()
	})
	total := 0
	for i, h := range hands {
		fmt.Println("Hand: ", h.cards, " Bid: ", h.bid, " Type: ", h.Type(), " Rank: ", i+1)
		total += h.bid * (i + 1)
	}
	return total
}
