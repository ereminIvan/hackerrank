package strings

import (
	"fmt"
	"os"
	sp "strings"
)

func CamelCase() {
	var inputString string
	var count int
	input := os.Stdin

	fmt.Fscanf(input, "%v", &inputString)
	if len(inputString) > 0 {
		count = 1
	}
	for _, cr := range inputString {
		if string(cr) == sp.ToUpper(string(cr)) {
			count++
		}
	}
	fmt.Println(count)
}
