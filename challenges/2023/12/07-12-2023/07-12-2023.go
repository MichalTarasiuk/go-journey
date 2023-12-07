package main

import (
	"challenges/lib"
	"fmt"
	"sort"
)

func main() {
	var hands []hand
	for _, ln := range lib.AocInputLines(2023, 07) {
		var cards string
		var bid int
		lib.Extract(ln, `^([2-9TJQKA]{5}) (\d+)$`, &cards, &bid)
		hands = append(hands, makeHand(cards, bid))
	}

	winnings := func(fn func(h hand) (handType, int)) int {
		sort.Slice(hands, func(i, j int) bool {
			ti, si := fn(hands[i])
			tj, sj := fn(hands[j])
			if ti < tj {
				return true
			} else if ti > tj {
				return false
			}
			return si < sj
		})
		var sum int
		for i, h := range hands {
			rank := i + 1
			sum += h.bid * rank
		}
		return sum
	}

	fmt.Println(winnings(func(h hand) (handType, int) { return h.handType, h.cardsScore }))
	fmt.Println(winnings(func(h hand) (handType, int) { return h.handType2, h.cardsScore2 }))
}

type hand struct {
	cards       string
	bid         int
	handType    handType
	cardsScore  int
	handType2   handType
	cardsScore2 int
}

func makeHand(cards string, bid int) hand {
	h := hand{cards: cards, bid: bid}

	cardCounts := make(map[rune]int, 13)
	for _, ch := range cards {
		cardCounts[ch]++
	}
	counts := lib.MapValues(cardCounts)
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	h.handType = classifyHand(counts)

	cardCounts2 := make(map[rune]int, 13)
	var jokers int
	for _, ch := range cards {
		if ch == 'J' {
			jokers++
		} else {
			cardCounts2[ch]++
		}
	}
	counts = lib.MapValues(cardCounts2)
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	if len(counts) == 0 {
		counts = []int{0} // handle all-jokers case
	}
	counts[0] += jokers
	h.handType2 = classifyHand(counts)

	h.cardsScore = computeCardsScore(cards, cardScores)
	h.cardsScore2 = computeCardsScore(cards, cardScores2)

	return h
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

func computeCardsScore(cards string, scores map[rune]int) int {
	var total int
	for _, ch := range cards {
		score, ok := scores[ch]
		lib.Assertf(ok, "No score for %q", ch)
		total = total*len(scores) + score
	}
	return total
}

var cardScores = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var cardScores2 = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
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

func classifyHand(counts []int) handType {
	switch {
	case counts[0] == 5:
		return fiveKind
	case counts[0] == 4:
		return fourKind
	case counts[0] == 3 && counts[1] == 2:
		return fullHouse
	case counts[0] == 3:
		return threeKind
	case counts[0] == 2 && counts[1] == 2:
		return twoPair
	case counts[0] == 2:
		return onePair
	default:
		return highCard
	}
}
