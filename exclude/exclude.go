package exclude

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Exclude(alphabet string, inputPath string, outputPath string) string {
	inputFile, inputErr := os.Open(inputPath)
	outputFile, outputErr := os.Create(outputPath)
	if inputErr != nil || outputErr != nil {
		fmt.Println(inputErr)
		fmt.Println(outputErr)
		log.Fatal(outputErr)
	}
	defer inputFile.Close()
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	letterRegex, rErr := regexp.Compile("^[a-z]*$")
	repeatString := "(a[a-z]*a)"
	for _, v := range alphabet {
		repeatString += fmt.Sprintf("|(%c[a-z]*%c)", v, v)
	}

	duplicateRegex, dErr := regexp.Compile(repeatString)
	if rErr != nil {
		log.Fatal(rErr)
	}

	if dErr != nil {
		log.Fatal(dErr)
	}

	fmt.Println("Here")
	count := 0
	for scanner.Scan() {
		s := scanner.Text()
		if letterRegex.MatchString(s) && !duplicateRegex.MatchString(s) {
			if _, err := outputFile.Write([]byte(s + "\n")); err != nil {
				panic(err)
			}
		}
		count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "nah"
}
