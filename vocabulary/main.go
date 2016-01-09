package main

// Simple example usage of
//  github.com/karan/vocabulary

import (
	"fmt"
	"log"
	"os"

	"github.com/karan/vocabulary"
)

func main() {
	// Set the API keys
	// Some functions require API keys. Refer to docs.
	// If API keys are not required, simple set empty strings as config:
	BigHugeLabsApiKey := ""
	WordnikApiKey := ""
	c := &vocabulary.Config{BigHugeLabsApiKey: BigHugeLabsApiKey, WordnikApiKey: WordnikApiKey}

	// Instantiate a Vocabulary object with your config
	v, err := vocabulary.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new vocabulary.Word object, and collects all possible information.
	// word, err := v.Word("vuvuzela")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("word.Word = %s \n", word.Word)
	// fmt.Printf("word.Meanings = %s \n", word.Meanings)
	// fmt.Printf("word.Synonyms = %s \n", word.Synonyms)
	// fmt.Printf("word.Antonyms = %s \n", word.Antonyms)
	// fmt.Printf("word.PartOfSpeech = %s \n", word.PartOfSpeech)
	// fmt.Printf("word.UsageExample = %s \n", word.UsageExample)

	if len(os.Args) < 2 {
		fmt.Println("Provide word as argument")
		os.Exit(1)
	}

	fmt.Println("Word:", os.Args[1])

	// Get just the synonyms
	fmt.Println("\n\nSynonym:\n-------")
	synonyms, err := v.Synonyms(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range synonyms {
		fmt.Println(s)
	}

	fmt.Println("\n\nMeaning:\n-------")
	meanings, err := v.Meanings(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range meanings {
		fmt.Println(s)
	}

	fmt.Println("\n\nUsage Example:\n-------")
	usages, err := v.UsageExample(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range usages {
		fmt.Println(s)
	}

	// Get just the antonyms
	// ants, err := v.Antonyms("love")
	// if err != nil {
	//   log.Fatal(err)
	// }
	// for _, a := range ants {
	//   fmt.Println(a)
	// }

	// Get just the part of speech
	// pos, err := v.PartOfSpeech("love")
	// if err != nil {
	//   log.Fatal(err)
	// }
	// for _, a := range pos {
	//   fmt.Println(a)
	// }

	// Can also use:
	//  v.UsageExample(word)

}
