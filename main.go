package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Read console input
	fmt.Println("Please enter your hex: ")
	reader := bufio.NewReader(os.Stdin)
	HexInput, _ := reader.ReadString('\n')

	HexInput = strings.TrimSpace(HexInput)                     //Trim space
	Base64Output := ConvertHex(HexInput)                       //Convert Hex to Base64
	fmt.Printf("Your Base64 encoded Hex is: %v", Base64Output) //Print output to console
}
