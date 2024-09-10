package wordfunctions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordStorer() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input the first letter for the word you want to generate here: ")
	firstLetter, _ := reader.ReadString('\n')
	firstLetter = strings.TrimSpace(firstLetter)

	fmt.Println("The first letter will be", firstLetter)
}
