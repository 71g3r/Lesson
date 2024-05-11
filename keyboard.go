// Package Keyboard  reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFloat Hello ! U can write number her

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num, err
}
