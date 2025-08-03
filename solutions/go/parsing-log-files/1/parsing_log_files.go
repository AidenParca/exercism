package parsinglogfiles

import "strings"
import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	valids := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
    for _, st := range valids {
        if strings.HasPrefix(text, st) {
			return true
		}
    }
    return false
}

func SplitLogLine(text string) []string {
    sanit := []string{"~","*","=","-"}
    for i := range sanit {
        text = strings.ReplaceAll(text, sanit[i], "")
    }
	return strings.Split(text, "<>")
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*?password.*?"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}


func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\bUser\s+([a-zA-Z0-9]+)`)
	result := []string{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			userName := matches[1]
			taggedLine := fmt.Sprintf("[USR] %s %s", userName, line)
			result = append(result, taggedLine)
		} else {
			result = append(result, line)
		}
	}

	return result
}

