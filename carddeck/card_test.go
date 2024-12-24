package deck

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Suit: Joker})
	fmt.Println(Card{Rank: Two, Suit: Spade})

	// Output:
	// Ace of Hearts
	// Joker
	// Two of Spades
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Errorf("Got: %d, Want: %d", len(cards), 13*4)
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
		t.Errorf("Got: %v, Want: %v", cards[0], Card{Rank: Ace, Suit: Spade})
	}
}

func TestShuffle(t *testing.T) {
	// Make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [24 3 ... ]
	shuffleRand = rand.New(rand.NewPCG(1, 2))

	orig := New()
	first := orig[24]
	second := orig[3]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Got: %s, Want: %s", cards[0], first)
	}
	if cards[1] != second {
		t.Errorf("Got: %s, Want: %s", cards[1], second)
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Got: %d, Want: %d", count, 3)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("Got %v - Expected all Twos and Threes to be filtered out.", c)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 Ranks * 4 Suits * 3 Decks
	want := 13 * 4 * 3
	if len(cards) != want {
		t.Errorf("Got: %d cards, Want: %d cards", len(cards), want)
	}
}
