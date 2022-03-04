package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	arrInput := []string{}
	for input.Scan() {
		arrInput = append(arrInput, input.Text())
	}
	if len(arrInput) <= 1 {
		return
	}
	for idx, row := range arrInput {
		sArr := strings.Split(row, " ")
		sort.Strings(sArr)
		for i := 0; i < len(sArr); i++ {
			for j := 0; j < len(sArr); j++ {
				

		}
		fmt.Printf("Case #%v: %v\n", idx, sArr)
	}

}
