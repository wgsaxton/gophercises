package deck

import (
	"fmt"
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