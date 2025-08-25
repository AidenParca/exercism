package twelve

import "fmt"
import "strings"

// gifts lists the gifts for each day of Christmas.
var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// days lists the ordinal numbers for each day.
var days = []string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

func Verse(i int) string {
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[i-1])
	var verseGifts []string
	for j := i - 1; j >= 0; j-- {
		verseGifts = append(verseGifts, gifts[j])
	}
	if len(verseGifts) == 1 {
		return verse + verseGifts[0] + "."
	}
	allButLastGift := strings.Join(verseGifts[:len(verseGifts)-1], ", ")

	lastGift := verseGifts[len(verseGifts)-1]

	return fmt.Sprintf("%s%s, and %s.", verse, allButLastGift, lastGift)
}

// Song generates the full lyrics for "The Twelve Days of Christmas".
func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
