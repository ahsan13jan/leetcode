package medium

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
//Given a string s, find the length of the longest substring without repeating characters.
//Example 1:
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// TODO Slow Need to optimize it
func lengthOfLongestSubstring(s string) int {

	maxSubStrCount := 0
	subStrCount := 0
	occurrence := map[byte]int{}
	bytes := []byte(s)

	index := 0
	for index < len(bytes) {
		c := bytes[index]

		val, ok := occurrence[c]
		if ok {
			occurrence = map[byte]int{}
			subStrCount = 0
			index = val
		} else {
			occurrence[c] = index
			subStrCount++

			if subStrCount > maxSubStrCount {
				maxSubStrCount = subStrCount
			}
		}
		index++
	}

	return maxSubStrCount
}