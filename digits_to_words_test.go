package digits_to_words

import "testing"

func TestDigitsToWords(t *testing.T) {
	t1 := DigitsToWords(1111)
	if t1 != "one thousand one hundred and eleventh" {
		t.Errorf("Return was incorrect, got: %s, want: one thousand one hundred and eleventh", t1)
	}
	t2 := DigitsToWords(1)
	if t2 != "first" {
		t.Errorf("Return was incorrect, got: %s, want: first", t2)
	}
}
