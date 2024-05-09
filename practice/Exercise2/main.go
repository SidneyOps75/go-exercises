package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatalf("[Iternal Error] \"standard.txt\" File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// skip first empty new line
	scanner.Scan()

	alphabet := map[byte][]string{}
	count := 0
	var letter byte = ' '
	alphabet[letter] = []string{}
	for scanner.Scan() {
		if count == 8 {
			letter++
			count = 0
		} else {
			alphabet[letter] = append(alphabet[letter], scanner.Text())
			count++
		}
	}

	args := strings.Split(os.Args[1], "\\n")

	for _, word := range args {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, l := range word {
					if l == 10 {
						fmt.Println()
					} else {
						fmt.Printf("%s", alphabet[byte(l)][i])
					}
				}
				fmt.Println()
			}
			// fmt.Println()
		}
	}
}
