package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/** TODO
  *  - allow users to customize length of sentence to type
  *  - allow users to specify alphabet
  *  - allow users to specify a range for length of words
  *  - improve the calcMistakes method
  *  - add weighting so that certain characters are more likely to appear
  *  - add a random word from dictionary option, real words selected from a dictionary
  *  - add high scores
**/

type Stats struct {
	time        float64
	length      int
	typedLength int

	cps float64
	cpm float64
	wps float64
	wpm float64

	mistakes int
}

func (s Stats) String() string {
	return fmt.Sprintf("time: %.2vs\nInput Length: %v\nTyped Length: %v\nChars/second: %.4v\n"+
		"Chars/minute: %.4v\nWords/second: %.4v\nWords/minutes: %.4v\nMistakes: %v\n", s.time, s.length,
		s.typedLength, s.cps, s.cpm, s.wps, s.wpm, s.mistakes)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trackTime(start time.Time, name string) (elapsed time.Duration) {
	elapsed = time.Since(start)
	fmt.Printf("\"%s\" took %s\n", name, elapsed)

	return elapsed
}

func calcMistakes(input string, typed string) (numMistakes int) {

	numMistakes = 0

	for i, _ := range typed {
		if input[i] != typed[i] {
			numMistakes++
		}
	}

	return numMistakes
}

func calcStats(input string, typed string, elapsed time.Duration) (stats Stats) {
	stats = Stats{}

	stats.time = elapsed.Seconds()
	stats.length = len(input)
	stats.typedLength = len(typed)

	stats.cps = float64(stats.length) / float64(stats.time)
	stats.cpm = stats.cps * 60
	// average of 4 chars per word, used to calculate wpm
	stats.wps = float64(stats.length) / 4 / float64(stats.time)
	stats.wpm = stats.wps * 60

	stats.mistakes = calcMistakes(input, typed)

	return stats

}

func readFromFileTest() {
	input, err := ioutil.ReadFile("input.txt")
	check(err)
	fmt.Println(string(input))

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")

	start := time.Now()

	typed, _ := reader.ReadString('\n')
	elapsed := trackTime(start, typed)

	stats := calcStats(string(input), typed, elapsed)
	fmt.Println(stats)

}

func randomWord() (str string) {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	// random length between 4 and 8
	length := rand.Intn(8) + 4

	var b bytes.Buffer

	for i := 0; i < length; i++ {
		b.WriteRune(alphabet[rand.Intn(len(alphabet)-1)])
	}
	str, err := b.ReadString(0x00)
	if err != io.EOF {
		check(err)
	}
	return str
}

func randomWordsTest() {
	const numRandomWords int = 10
	sentence := make([]string, numRandomWords, numRandomWords)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numRandomWords; i++ {
		if i < numRandomWords-1 {
			sentence[i] = randomWord() + " "
		} else {
			sentence[i] = randomWord() + "\n"
		}

	}
	input := strings.Join(sentence, "")

	fmt.Print(input)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")

	start := time.Now()

	typed, _ := reader.ReadString('\n')
	elapsed := trackTime(start, typed)

	stats := calcStats(string(input), typed, elapsed)
	fmt.Println(stats)

}

func main() {
	//readFromFileTest()
	randomWordsTest()

}
