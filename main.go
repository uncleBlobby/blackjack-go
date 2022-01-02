package main

import "fmt"
import "math/rand"
import "time"

type Player struct {
    hand []Card
    score int
}

type Card struct {
    suit string
    value int
    name string
}

type Suit struct {
    symbol string
    colour string
    name string
}

type Deck struct {
    size int
    cards []Card	
}

func (d *Deck) shuffle() Deck {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(d.cards), func(i, j int) {
    d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
    })
    return *d
}

func (d *Deck) dealCard(p *Player) Deck {
	newCard := Card{}
	newCard, d.cards = d.cards[0], d.cards[1:]
	p.hand = append(p.hand, newCard)
    return *d
}

func (p *Player) getScore() Player {
    sum := 0
    //fmt.Println("Hand: ", p.hand)
    for i := 0; i < len(p.hand); i++ {
	sum = sum + p.hand[i].value
    }
    p.score = sum
    return *p
}

func generateDeck(size int) Deck { 
	deck := Deck{}
	suit := ""
    for i := 0; i < 4; i++ {
	for j := 1; j < 14; j++ {
	    if i == 0 {
		suit = "Hearts"
	    }
	    if i == 1 {
		suit = "Diamonds"
	    }
	    if i == 2 {
	    	suit = "Clubs"
	    }
	    if i == 3 {
		suit = "Spades"
	    }
	    card := Card{suit: suit, value: j, name: ""} 
	    deck.cards = append(deck.cards, card)
	}
    }
    return deck
}

func addValuesToCards(deck Deck) Deck {    
    for i := 0; i < len(deck.cards); i++ {
	value := deck.cards[i].value	
	switch value {
	    case 1:
		deck.cards[i].name = "one"
	    case 2:
		deck.cards[i].name = "two"
	    case 3:
		deck.cards[i].name = "three"
	    case 4:
		deck.cards[i].name = "four"
	    case 5:
		deck.cards[i].name = "five"
	    case 6:
		deck.cards[i].name = "six"
	    case 7:
		deck.cards[i].name = "seven"
	    case 8:
		deck.cards[i].name = "eight"
	    case 9:
		deck.cards[i].name = "nine"
	    case 10:
		deck.cards[i].name = "ten"
	    case 11:
		deck.cards[i].name = "jack"
	    case 12:
		deck.cards[i].name = "queen"
	    case 13:
		deck.cards[i].name = "king"
	}
    }

    return deck
}

func prettyPrintDeck(deck Deck) {
    for i := 0; i < 13; i++ {
	fmt.Printf("%v	,  %v	,  %v	,  %v	\n", deck.cards[i], deck.cards[i+13], deck.cards[i+26], deck.cards[i+39])
    }
}

func main(){

    player := Player{}
    playerp := &player
    dealer := Player{}
    dealerp := &dealer
    deck := generateDeck(52)
    deckp := &deck
    //prettyPrintDeck(deck)
    deck = addValuesToCards(deck)
    //prettyPrintDeck(deck)
    deck = deckp.shuffle()
    deck = deckp.dealCard(playerp)
    deck = deckp.dealCard(dealerp)
    deck = deckp.dealCard(playerp)
    deck = deckp.dealCard(dealerp)
    player = playerp.getScore()
    dealer = dealerp.getScore()
    fmt.Printf("Player hand: %v, Player score: %v\n", player.hand, player.score)
    fmt.Printf("Dealer hand: %v, Dealer score: %v\n", dealer.hand, dealer.score)
    fmt.Printf("%v cards remain in deck.\n", len(deck.cards))

}
