package problem0340

func LengthOfLongestSubstringKDistrinct(s string, k int) int {
	maxLength, windowStart := 0, 0

	//MOC
	//convert string into an array of chars
	//create a slice of length k to save the current alphabets in scope
	//create counters for start and end of the window
	//save current alphabet into slice
	//move to next and see if the current alphabet exists in slice already
	//if yes then move to next ..
	//if no then add the char to slice and move to next
	chars := []rune(s)
	temp := make([]rune, k)

	for windowEnd := 0; windowEnd < len(chars); windowEnd++ {
		rightChar := chars[windowEnd]
		if !contains(temp, rightChar) {
			temp = append(temp, rightChar)
		}
	}
	maxLength = len(temp)

	return maxLength
}

func contains(s []rune, str rune) bool {

	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false

}
