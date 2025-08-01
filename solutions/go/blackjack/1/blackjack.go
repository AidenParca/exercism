package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch {
    case card == "ace" : return 11
    case card == "two" : return 2
    case card == "three" : return 3
    case card == "four" : return 4
    case card == "five" : return 5
    case card == "six" : return 6
    case card == "seven" : return 7
    case card == "eight" : return 8
    case card == "nine" : return 9
    case card == "ten" : return 10
    case card == "jack" || card == "queen" || card == "king": return 10
    default: return 0
        }
    panic("Please implement the ParseCard function")
    
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    c1 := ParseCard(card1)
    c2 := ParseCard(card2)
    dil := ParseCard(dealerCard)
    switch {
        case c1+c2 == 22 : return "P"
        case c1+c2 == 21 && dil <= 9: return "W"
        case c1+c2 == 21 && dil > 9: return "S"
        case c1+c2 > 16 && c1+c2 < 21 : return "S"
        case c1+c2 > 11 && c1+c2 < 17 && dil < 7: return "S" // 10,6,6
        case c1+c2 > 11 && c1+c2 < 17 && dil >= 7: return "H"
        case c1+c2 < 12 : return "H"
        
    }
	panic("Please implement the FirstTurn function")
}
