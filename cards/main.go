package main

func main() {
	//cards := newDeckFromFile("my")

	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards.saveToFile("my_cards")*/

	/*hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()*/

}
