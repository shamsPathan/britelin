package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your_script.go <value_to_add>")
		return
	}

	input := os.Args[1]

	var valueToAdd int
	var err error

	if strings.HasPrefix(input, "-") {
		valueToAdd, err = strconv.Atoi(input[1:]) // Extract the number part without the '-'
		valueToAdd = -valueToAdd                  // Convert it to a negative value
	} else {
		valueToAdd, err = strconv.Atoi(input)
	}

	if err != nil {
		fmt.Println("Error parsing input value:", err)
		return
	}

	filePath := "/sys/class/backlight/acpi_video0/brightness"

	content, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	currentBrightness, err := strconv.Atoi(strings.TrimSpace(string(content)))
	if err != nil {
		fmt.Println("Error parsing brightness value:", err)
		return
	}

	updatedBrightness := currentBrightness + valueToAdd

	updatedBrightnessStr := strconv.Itoa(updatedBrightness)

	err = os.WriteFile(filePath, []byte(updatedBrightnessStr), 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
