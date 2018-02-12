package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Create a worldList type to reference with our slices of strings
type wordList []string

// Primary coding challenge function, called from main func
func runCodingChallenge(controlWord string) {
	// Declare the words slice of strings
	words := wordList{controlWord, "a", "ab", "ac", "ad", "abc", "ade"}
	// Print the beginning values of the word list
	fmt.Println("_______\nOriginal Control Word and List\n_______")
	words.print()
	// Convert the slice (of strings) to string for file saving
	words.toString()
	// Save the string to words.txt
	words.saveToFile("words.txt")
	// Read back the words.txt file into a copy of the words slice
	words = readFromFile("words.txt")
	// Sort the words list based on match completeness requirements
	words.sortWordList(controlWord)
}

func (w wordList) toString() string {
	// Return the words element as a single string, values separated by comma
	return strings.Join([]string(w), ",")
}

func (w wordList) saveToFile(filename string) error {
	// Create the words.txt file while converting the words string to byte type
	return ioutil.WriteFile(filename, []byte(w.toString()), 0666)
}

func readFromFile(filename string) wordList {
	// Create a bytestream from words.txt
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Return the bytestream after converting to a string and splitting on commas
	s := strings.Split(string(bs), ",")
	return wordList(s)
}

func (w wordList) sortWordList(controlWord string) {
	// Create a new sortedList slice of wordList type
	sortedList := wordList{}
	// Iterate all words in the original word list slice referenced here via receiver...
	for _, word := range w {
		// Create the ld (length difference) element and start with the difference
		// between the control word length and the word list item length
		ld := len(controlWord) - len(word)
		// Start iterating through each word list item character to identify how
		// many characters match the control word and how many don't
		c := 0
		for c < len(word) {
			// If the characters of each identical index don't match...
			if string(controlWord[c]) != string(word[c]) {
				// Increment the ld (length difference) counter by 1
				ld++
			}
			c++
		}
		// Append each word list item to the sortedList slice, concatenating the
		// value of ld (length difference) at the beginning
		sortedList = append(sortedList, strconv.Itoa(ld)+"_"+word)
	}
	// Sort the fully prepared sortedList slice using Go's standard sort method,
	// which will first look at the numerical part of the string (reason we created
	// the ld element) and then the alphabetic part next
	sort.Strings(sortedList)
	// Now, iterate through each sortedList slice string again and using regex
	// replace all non a-z characters with empty values to remove the ld sorting
	// mechanism created prior
	c := 0
	for c < len(sortedList) {
		reg, err := regexp.Compile("[^a-z]+")
		if err != nil {
			log.Fatal(err)
		}
		// Replace non a-z values
		sortedList[c] = reg.ReplaceAllString(sortedList[c], "")
		c++
	}
	// Print the newly sorted list
	fmt.Println("_______\nSorted Words Based on Match Completeness of Word List against Control Word\n_______")
	sortedList.print()
}

// Function to print each index and word from the calling receiver
func (w wordList) print() {
	for i, word := range w {
		fmt.Println(i, word)
	}
}
