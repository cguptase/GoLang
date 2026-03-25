package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 13)
	hand.print()
	remainingCards.print()
}
