package shape

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a letter: ")

	letter, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return letter
}

func getTirnagleWidth(letter rune) int {
	//Check what number a letter represents between 1 and 26 (A-Z)
	if unicode.IsLower(letter) {
		return int(letter-'a') + 1
	} else {
		return int(letter-'A') + 1
	}
}

func getLetterFromNumber(number int) rune {
	//Get the letter that represents a number between 1 and 26
	return rune(number + 'A' - 1)
}

func drawDiamondWithLetter(width int) {
	// Dar the outlike of a diamond shape with the letter on the outside

	// Upper triangle
	for i := 0; i < width; i++ {
		for j := 0; j < width-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			if j == 0 || j == 2*i {
				fmt.Print(string(getLetterFromNumber(i + 1)))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// Lower Triangle
	for i := width - 2; i >= 0; i-- {
		for j := 0; j < width-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			if j == 0 || j == 2*i {
				fmt.Print(string(getLetterFromNumber(i + 1)))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Diamond() {
	//accept a letter as an input from stdin
	input := getInput()

	//handle the case where the input is more than one character
	if len(input) > 2 {
		fmt.Println("Please enter only one letter")
		return
	}
	//handle the case where the input is not a letter
	if !unicode.IsLetter(rune(input[0])) {
		fmt.Println("Please enter a letter")
		return

	}

	//first character of the input
	character := rune(input[0])

	drawDiamondWithLetter(getTirnagleWidth(character))
}
