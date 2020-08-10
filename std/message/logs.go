package main

import (
	"fmt"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor		= "\033[1;34m%s\033[0m"
	WarningColor	= "\033[1;33m%s\033[0m"
	ErrorColor		= "\033[1;31m%s\033[0m"
)

func main() {
	showMessage(INFO, "Welcome to logs!")
	showMessage(INFO, fmt.Sprintf("About to open %s", fileName))

	fileName := "logs.txt"
	file, err := os.Open(fileName)
	if err != nil {
		showMessage(ERROR, err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func showMessage(messagetype messageType, message string) {
	switch messagetype {
		case INFO: 
			printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
			fmt.Printf(InfoColor, printMessage)
		case WARNING: 
			printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
			fmt.Printf(InfoColor, printMessage)
		case ERROR: 
			printMessage := fmt.Sprintf("\nError: \n%s\n", message)
			fmt.Printf(InfoColor, printMessage)
	}
}
