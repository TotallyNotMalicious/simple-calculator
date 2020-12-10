package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	repeat := true

	for repeat {
		fmt.Println("\nPlease Select An Option\n\nAddition\nSubtraction\nMultiplication\nDivision\n\n----------------------")
		var userinput string
		fmt.Scanln(&userinput)
		option := getOption(toLowercase(userinput))
		if option == Option(NONE) {
			fmt.Println("" + userinput + " Is An Invalid Option, Exiting Now")
			return
		}
		fmt.Println("----------------------\nEnter The First Number\n----------------------")
		fmt.Scanln(&userinput)
		int1 := makeInt(userinput)
		fmt.Println("----------------------\nEnter The Second Number\n----------------------")
		fmt.Scanln(&userinput)
		int2 := makeInt(userinput)
		var response string

		switch option {
		case DIVISION:
			response = makeString(int1 / int2)
			break
		case MULTIPLICATION:
			response = makeString(int1 * int2)
			break
		case SUBTRACTION:
			response = makeString(int1 - int2)
			break
		case ADDITION:
			response = makeString(int1 + int2)
			break
		default:
			response = "0"
			break
		}

		fmt.Println("----------------------\nAnswer: " + response)
		fmt.Println("----------------------\nWould You Like To Solve Another Problem? (y/n)\n----------------------")
		var again string
		fmt.Scanln(&again)
		if toLowercase(again) == "y" || toLowercase(again) == "yes" {
			repeat = true
		} else {
			fmt.Println("----------------------\nExiting Now")
			repeat = false
		}
	}
}

func makeString(value int) string {
	return strconv.Itoa(value)
}

func makeInt(value string) int {
	toreturn, _ := strconv.Atoi(value)
	return toreturn
}

func toLowercase(value string) string {
	return strings.ToLower(value)
}

func getOption(value string) Option {
	switch value {
	case "division":
		return DIVISION
	case "subtraction":
		return SUBTRACTION
	case "addition":
		return ADDITION
	case "multiplication":
		return MULTIPLICATION
	}

	return NONE
}

func toString(value Option) string {
	return string(value)
}

type Option string

const (
	DIVISION       = "division"
	MULTIPLICATION = "multiplication"
	ADDITION       = "addition"
	SUBTRACTION    = "subtraction"
	NONE           = "none"
)
