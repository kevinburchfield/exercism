package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	values := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"ten": 10,
		"jack": 10,
		"queen": 10,
		"king": 10,
		"ace": 11,
	}
	return values[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealerSum := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	} else if sum == 21 {
		if (dealerSum != 11 && dealerSum != 10) {
			return "W"
		} else {
			return "S"
		}
	} else if sum > 16 {
		return "S"
	} else if sum > 11 {
		if dealerSum > 6 {
			return "H"
		} else {
			return "S"
		}
	} else {
		return "H"
	}
}
