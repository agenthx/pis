package main

import (
	"bufio"
	"fmt"
	"os"
	piscine "piscine/functions"
	"regexp"
	"strings"
)

func main() {
	file, ferr := os.Open("text files/sample.txt") //open the file sample.txt
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //read the contents of the file and make the changes
	var items []string
	for scanner.Scan() {
		line := scanner.Text()
		items = strings.Split(line, " ")
		for i, word := range items {
			if word == "(bin)" {
				items[i-1] = piscine.Bin2Dec(items[i-1])
			} else if word == "(hex)" {
				items[i-1] = piscine.Hex2Dec(items[i-1])
			} else if word == "(up)" {
				items[i-1] = strings.ToUpper(items[i-1])
			} else if word == "(low)" {
				items[i-1] = strings.ToLower(items[i-1])
			} else if word == "(cap)" {
				items[i-1] = strings.Title(items[i-1])
			}
		}
	}
	//store the new sentence and remove the (hex), etc..
	sentence := ""
	for _, word := range items {
		sentence += word + " "
	}
	regex := regexp.MustCompile(`\(bin\)|\(hex\)|\(up\)|\(low\)|\(cap\)`)
	sentence = regex.ReplaceAllLiteralString(sentence, "")
	//fix vowels
	sentence = piscine.Vowels(sentence)
	//remove extra spaces caused by regex
	sentence = piscine.Spaces(sentence)
	//fix punctuations
	sentence = piscine.Punct(sentence)
	//create result file
	file, err := os.Create("text files/result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//write the sentence in the file
	_, werr := file.WriteString(sentence)
	if err != nil {
		panic(werr)
	}
	defer file.Close()

	fmt.Println("READY!!!")
}
