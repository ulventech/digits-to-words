package digits_to_words

import (
	"strconv"
	"strings"
	"fmt"
)

func DigitsToWords(i int) string {
	if i > 10000 || i < 0 {
		return "err: invalid input"
	}

	if i == 10000 {
		return "ten thousandth"
	}

	if i < 20 {
		return ordinalNumbers[i]
	}

	// any digits less than 100
	if i < 100 {
		return toOrdinal(i)
	}

	s := fmt.Sprintf("%d", i)
	// any digits more than 100
	switch len(s) {
	case 3: // hundreds
		lastTwoDigits := i % 100
		result := lessThanTwenty[s[0]-48]
		if lastTwoDigits == 0 {
			return result + " hundredth"
		}
		result = result + " hundred " + toOrdinal(i)
		return result
	case 4: // thousands
		lastTwoDigits := i % 100
		lastThreeDigits := i % 1000
		result := lessThanTwenty[s[0]-48]
		if lastThreeDigits == 0 {
			result = result + " thousandth"
			return result
		}
		if lastThreeDigits < 100 {
			result = result + " thousand " + toOrdinal(i)
			return result
		}

		if lastTwoDigits == 0 {
			result = result + " thousand " + lessThanTwenty[s[1]-48] + " hundredth " + toOrdinal(i)
			return result
		}
		result = result + " thousand " + lessThanTwenty[s[1]-48] + " hundred " + toOrdinal(i)
		return result
	case 5: // ten thousands

	default:
		println(i)
	}
	return toOrdinal(i)
}


func toOrdinal(i int) string {
	lastTwoDigits := i % 100

	if lastTwoDigits == 0 {
		return ""
	}

	if lastTwoDigits < 20 {
		result := ordinalNumbers[lastTwoDigits]
		return "and " + result
	}
	s := strings.Split(strconv.Itoa(lastTwoDigits), "")
	first, _ := strconv.Atoi(s[0])
	second, _ := strconv.Atoi(s[1])
	result := tensLessThanHundred[first]

	if second == 0 {
		result = strings.Replace(result, "y", "ieth", -1)
		if i > 100 {
			return "and " + result
		}
		return result
	}

	if i > 100 {
		result = "and " + result + " " + ordinalNumbers[second]
		return result
	}
	return result + " " + ordinalNumbers[second]
}

var ordinalNumbers = []string{"",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
	"thirteenth",
	"fourteenth",
	"fifteenth",
	"sixteenth",
	"seventeenth",
	"eighteenth",
	"nineteenth",
}

var lessThanTwenty = []string{"",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tensLessThanHundred = []string{"", "",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}