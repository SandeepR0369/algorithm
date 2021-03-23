package main 

import (
"fmt"
"encoding/csv"
"os"
)

const (
	EXE_FILE = "problems.csv"
)

func main(){
quiz, err := os.Open("/Users/V771254/go/src/quiz/"+ EXE_FILE)
if err != nil {
        return
    }
defer quiz.Close()

lines, err := csv.NewReader(quiz).ReadAll()
    if err != nil {
        return 
    }
fmt.Println("LINES", lines)


}
