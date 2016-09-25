package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"bufio"
	"os"
)

func TranslateWord(words []string, mistakes int) {
   
        rand.Seed(time.Now().Unix())
        random := rand.Intn(len(words))
        translateWord := strings.Split(words[random], ":")
         
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Translate: " + translateWord[0])
        text, _ := reader.ReadString('\n')

        if text == translateWord[1]+"\n" && len(words)>1{
                fmt.Println("Correct!")
                words = append(words[:random], words[random+1:]...)
                TranslateWord(words,mistakes)
        } else if text == translateWord[1]+"\n" && len(words)<=1{
                fmt.Println("That's all! Mistakes: ",mistakes)
                os.Exit(0)
        } else {
                fmt.Println("Wrong motherfucker! You'll be asked this again")
                mistakes = mistakes + 1
                TranslateWord(words,mistakes)
        }
}

func GetWordList(listName string) []string {
	file, err := os.Open(listName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
func main() {
	words := GetWordList(os.Args[1])
	TranslateWord(words, 0)

}
