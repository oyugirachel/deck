//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit Creation type
type Suit uint8

// Spade declaration of a constant which will start counting from one
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker //Special case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank type creation
type Rank uint8

// Ace constant creation which will increment gradually
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Allows us to iterate over the constants
const (
	minRank = Ace
	maxRank = King
)

// Card type struct holding the Suit and Rank
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New function which will hold a slice of cards
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	//for each suit
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})

		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	//for each rank
	//for each card, suit

	return cards

}
 

// DefaultSort sorts the cars
func DefaultSort(cards []Card) []Card{
	sort.Slice(cards,Less(cards))
	return cards
}

// Sort function(generic version) takes in a slice of cards and returns the less function (Second part being the return type)
func Sort(less func(cards []Card) func(i,j int) bool ) func([]Card) []Card{
	return func(cards []Card) []Card{
		sort.Slice(cards, less(cards))
		return cards

	}
}

// Less function which sorts the provided slice check :https://golang.org/pkg/sort/ 
func Less(cards []Card ) func(i,j int) bool{
	return func (i,j int) bool{
		return absRank(cards[i]) < absRank(cards[j])
	}
}

 // absolute rank function
func absRank(c Card) int {
	 return int(c.Suit) * int(maxRank) + int(c.Rank)
	        
}

// Shuffle function takes in a slice of cards(Basically figuring out the order of the cards)
func Shuffle(cards []Card) []Card{
	// Creating a new slice of cards
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	// perm ={0,1,4,2,3} (j are those values inside)
	// i is the index of the permutation
	for i, j :=range perm{
		// cards is the original slice while j is the value of the permutation 
		ret[i] =cards[j]


	}
	return ret
	

}
// Jokers function that you pass the number of jokers to be placed in the deck
func Jokers(n int) func([]Card) []Card{
	return func(cards []Card) []Card{
		for i :=0; i<n; i++{
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit:Joker , 
			} )
		}
		return cards
	}
}

// Filter function
func Filter(f func(card Card) bool) func([]Card) []Card{
	return func (cards []Card) []Card{
		var ret []Card
		for _, c :=range cards{
			if !f(c){
				ret = append(ret,c)
			}

		}
		return ret

	}  
}

// Deck function that allows for the addition of new decks
func Deck(n int) func([]Card) []Card{
	return func(cards []Card) []Card{
		
	}
}

		
	

