package main

import (
	"fmt"
)

//Given a list of words, return the shortest unique prefix of each word. For example, given the list:
//Input: dog cat apple apricot fish
//Return the list: d c app apr f

func main(){
	
}

func prefix(str []string)[]string{
	str1 := strings.Split(str)
}

func main() {
	str := []string{"Hello", "Playground", "Jaffa", "Badkav"}
	sort.Strings(str)

	for i := 0; i <= len(str)-1; i++ {
		str1 := strings.ToLower(str[i])
		str2 := strings.Split(str1, "")
		fmt.Println(str2)
	}
	fmt.Println("S1",str)

}

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := []string{"Hello", "Apricot", "Playground", "Apple", "Jaffa", "Badkav"}
	//strings.ToLower(str)
	sort.Strings(str)

	fmt.Println(str)

	mapper := make(map[int][]string)

	for i := 0; i <= len(str)-1; i++ {
		str1 := strings.ToLower(str[i])
		str2 := strings.Split(str1, "")

		mapper[i] = str2
	}

	for k, v := range mapper {
		fmt.Println("K", k, "V", v)
	}
	fmt.Println("M", mapper)

}

func main() {
	str := []string{"Hello", "Apricot", "Playground", "Apple", "Jaffa", "Badkav"}
	//var abc []string
	sort.Strings(str)

	fmt.Println(str)

	mapper := make(map[int][]string)

	for i := 0; i <= len(str)-1; i++ {
		str1 := strings.ToLower(str[i])
		str2 := strings.Split(str1, "")
		mapper[i] = str2
	}

	fmt.Println("M", mapper)

	for j := 0; j <= len(mapper); j++ {
		fmt.Println("J", mapper[j])

		for k, v := range mapper[j] {
			fmt.Println(k, v)
		}

	}

	//for k:=0; k <= len

}

