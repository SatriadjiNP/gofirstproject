package main

import (
	"fmt"
	"strings"
	"unicode"
)

// challenge soal 1
func countVocalConsonant(word string) {
	// var word string
	// fmt.Println("Masukan kalimat : ")
	// fmt.Scan(&word)

	// word = strings.ToLower(word)

	// vocal := 0
	// consonant := 0

	// for _, kata := range word { //
	// 	if kata == 'a' || kata == 'i' || kata == 'u' || kata == 'e' || kata == 'o' {
	// 		vocal++
	// 	} else if kata == ' ' {
	// 		continue
	// 	} else {
	// 		consonant++
	// 	}
	// }
	// fmt.Println("Huruf Vocal : ", vocal)
	// fmt.Println("Huruf Consonant : ", consonant)

}

// soal 2
func countWordWithSpace(word string) {
	// spasi := 1
	// for _, total := range word {
	// 	if total == ' ' {
	// 		spasi++
	// 	}
	// }
	// fmt.Println("Total Word : ", spasi)

	// kata := strings.Fields(word)
	// counter := 0

	// for index, value := range kata {
	// 	runes := []rune(value)
	// 	counter++
	// 	kata[index] = string(runes)
	// }
	// fmt.Println("Total Word : ", counter)

	kata := strings.Split(word, " ")
	count := len(kata)

	fmt.Println("Total Word : ", count)
}

// soal 3
func countWordWithUpper(word string) {
	upper := 0
	for _, total := range word {
		if unicode.IsUpper(total) {
			upper++
		}
	}
	fmt.Println("Total Word : ", upper)
}

// soal 4
func findWord(word string, search string) {
	kata := strings.Split(word, " ")
	katas := 0
	for _, value := range kata { //value variable penampung untuk kata yang udah di split
		if strings.ToLower(value) == strings.ToLower(search) {
			katas++
		}

	}
	fmt.Printf("Word go found %d Times ", katas)
}

// soal 5
func replaceMiddleCharWithStar(word string) {
	/* kata := strings.Fields(word) //mengubah struktur kata jadi array, dan mempunyai index

	for index, kalimat := range kata{
		runes := []rune(kalimat) //
		for char := 1; char < len(runes)-1; char++ {
			runes[char] = '*'
		}
		kata[index] = string(runes)
	}
	hasil_kalimat := strings.Join(kata, " ")
	fmt.Println("Result : ", hasil_kalimat) */

	wordSplit := strings.Split(word, " ")

	for i, value := range wordSplit {
		if len(value) > 2 {
			firstLetter := value[:1]
			sensorLetter := strings.Repeat("*", len(value)-2)
			endLetter := value[len(value)-1:]

			sensorWord := firstLetter + sensorLetter + endLetter
			wordSplit[i] = sensorWord
		}
	}

	fmt.Println("Result : ", strings.Join(wordSplit, " "))

	//cara for biasa bukan range
	// words := strings.Split(input, " ")
	// output := ""

	// for count := 0; count < len(words); count++ {
	// 	word := words[i]

	// 	if len(word) > 2 {
	// 		middle := word[:len(word)-1]
	// 		star := strings.Repeat("*", len(middle))

	// 		word = string(word[0]) + star + string(word[len(word)-1])
	// 	}
	// 	output += word + " "
	// }
	// fmt.Println("Result: ", output)

}

func main() {
	// countVocalConsonant("Rockbye Baby")
	// countWordWithSpace("How Many Word In This Sentence")
	// countWordWithUpper("HowManyWordInThisSentence")
	// findWord("Bootcamp GO di codeid, run go Run", "go")
	// replaceMiddleCharWithStar("aku suka join codeid bootcamp")
}
