package main

import (
	"fmt"
	"sort"
	"strings"
	"reflect"
)

func main() {
	str := []string{"Hello", "Apricot", "Playground", "Apple", "Jaffa", "Badkav"}
	//var abc []string
	fmt.Println(prefixer(str))

	/*sort.Strings(str)

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
			//fmt.Println(k, v)
			for h := 0; h <= len(mapper[j]); h++ {
				if v = mapper[j][h]; ok {
					return v
				}
			}
		}

	}*/

}

func prefixer(str []string) string {
	sort.Strings(str)
	mapper := make(map[int][]string)
	//map2 := make(map[int][]string)

	for i := 0; i <= len(str)-1; i++ {
		//fmt.Println("jaffa", str[i])
		str1 := strings.ToLower(str[i])
		str2 := strings.Split(str1, "")
		mapper[i] = str2
	}

	fmt.Println("M", mapper)
	fmt.Println("Len", len(mapper))


	for j := 0; j <= len(mapper); j++ {
	
		fmt.Println("Len1", len(mapper))
		fmt.Println("J", mapper[j])
		fmt.Println("Type", reflect.TypeOf(mapper[j]))

		for _, v := range mapper[j] {
			fmt.Println("V1", mapper[j])

			for h := 0; h <= len(mapper[j]); h++ {

				fmt.Println("bsdk", mapper[j][h])
				//str3 := strings.Split(mapper[j], "")

				//fmt.Println("3", str3)
				if v == mapper[j][h] {
					return v
				} else {
					return "Erri Pooka"
				}
			}
		}

	}
	return "Pora Pooka"
}
