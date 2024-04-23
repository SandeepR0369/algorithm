// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Reverse:", reverse("Hello"))
	fmt.Println("Missing:", missing([]int{3, 0, 1}))
	fmt.Println("Merge:", merge([]int{1, 3, 5}, []int{2, 4, 6}))
	fmt.Println("SubString:", substring("abcabcdbb"))
}

// ind All Anagrams in a String:
// Input: s = "cbaebabacd", p = "abc"
// Output: [0, 6]

//Longest Substring Without Repeating Characters:
//Input: "abcabcdbb"
//Output: 3 (for substring "abc")

func substring(s string) string {
	maxLen := 0
	maxStr := ""

	for i := 0; i < len(s); i++ {
		sub := ""
		seen := make(map[byte]bool)
		//	fmt.Println("UY", i, seen)
		for j := i; j < len(s); j++ {
			//	fmt.Println("JJ", string(s[j]))
			if seen[s[j]] {
				//	fmt.Println("BRK", string(s[j]), seen[s[j]])
				break // <-- This break statement terminates the inner loop
			}
			//	fmt.Println("AB", j, seen, sub)
			seen[s[j]] = true
			sub += string(s[j])
			if len(sub) > maxLen {
				maxLen = len(sub)
				maxStr = sub
			}
		}
	}

	return maxStr
}

// Reverse a String:
// Input: "hello"
// Output: "olleh"
func reverse(s string) string {
	a := strings.Split(s, "")
	/*if err != nil {
		fmt.Println("Unable to split")
	}*/
	b := make([]string, len(a))
	for k, v := range a {
		b[len(a)-k-1] = v
	}

	return strings.Join(b, "")
}

// Find the Missing Number:
// Input: [3,0,1]
// Output: 2
func missing(a []int) int {
	sort.Ints(a)
	min := a[0]
	max := a[len(a)-1]
	num := 0
	for i := 0; i < len(a); i++ {
		if a[i+1]-a[i] == 1 {
			continue
		} else {
			num = a[i] + 1
			return num
		}
	}

	if min != 0 {
		return 0
	}
	return max + 1
}

// Merge Two Sorted Arrays:
// Input: [1, 3, 5], [2, 4, 6]
// Output: [1, 2, 3, 4, 5, 6]
func merge(A, B []int) []int {
	for _, v := range B {
		A = append(A, v)
	}
	sort.Ints(A)
	return A
}
