// Package proverb generates a proverb from a list of strings.
package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	// If the input slice is empty, there's no proverb to generate.
	if len(rhyme) == 0 {
		return []string{}
	}

	var proverbLines []string

	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverbLines = append(proverbLines, line)
	}

	finalLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverbLines = append(proverbLines, finalLine)

	return proverbLines
}
