package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello world!!")
	var intNum int = 32
	var floatNum float32 = 32.12
	var floatNum64 float64 = 32.1222

	fmt.Println(intNum)
	fmt.Println(floatNum)
	fmt.Println(floatNum64)

	// arth
	fmt.Println("Multiply", intNum*2)
	fmt.Println("Division", intNum/2)
	fmt.Println("Mod rem", intNum%2)

	var name string = `Sarvesh Sawant`
	var age int8 = 32
	fmt.Printf("My name is %s and age is %d \n", name, age)

	fmt.Printf("Number of characters in string %d \n", utf8.RuneCountInString(name))

	company_name := "xyz"
	fmt.Printf("Company name %s \n", company_name)

	salary, age := 1, 1
	fmt.Printf("%d, %d \n", salary, age)

	const myName string = "Lakhan"

	fmt.Printf("My name is %s \n", myName)

	printMe("Sarvesh")

	var div, rem, err = division(4, 2)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("Reminder is %v \n", rem)
	default:
		fmt.Printf("Diviser is %v and reminder is %v", div, rem)
	}

	switch rem {
	case 0:
		fmt.Printf("Reminder is zero \n")

	case 1, 2:
		fmt.Printf("Reminder is either 1 or 2 \n")

	default:
		fmt.Printf("Different rem \n")
	}
}

func printMe(printvalue string) {
	fmt.Println(printvalue)
}

func division(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	return num / den, num % den, nil
}
