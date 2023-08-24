package game

func insertLineBreaks(input string, charsPerLine int) string {
	processedOutput := ""
	for i, runeValue := range input {
		if i > 0 && i%charsPerLine == 0 {
			processedOutput += "\n"
		}
		processedOutput += string(runeValue)
	}
	return processedOutput
}
