package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)


type Question struct {
	Statement string
	UserAns   int
	CompAns   int
}

const (
	TotalQuestions = 5
	MaxScore       = 5
	MinScore       = 1
)

//Initialize Questions and randomly generated Computer responses
var questions = map[int]Question{
	1: {Statement: "I think Sundays are perfect for lazy mornings and breakfast in bed.", CompAns: rand.Intn(MaxScore) + 1},
	2: {Statement: "I love spontaneous road trips and discovering new places.", CompAns: rand.Intn(MaxScore) + 1},
	3: {Statement: "I think dogs are the best companions.", CompAns: rand.Intn(MaxScore) + 1},
	4: {Statement: "Broccoli is totally underrated.", CompAns: rand.Intn(MaxScore) + 1},
	5: {Statement: "I think pineapple belongs on pizza.", CompAns: rand.Intn(MaxScore) + 1},
}

var thresholds = []struct {
	minScore int
	message  string
}{
	{22, "True love!"},
	{12, "Just friends."},
	{0, "Run away."},
}

func main() {
	welcomeMsg()
	totalScore := displayQs(questions)
	percentageScore := (float64(totalScore) / float64(TotalQuestions*MaxScore)) * 100
	fmt.Printf("\nYour overall compatibility score is: %.2f%%\n", percentageScore)

	for _, threshold := range thresholds {
		if totalScore >= threshold.minScore {
			fmt.Println(threshold.message)
			break
		}
	}
	fmt.Println("Thank you for using Match Maker!")
}

func welcomeMsg() {
	fmt.Println("***********************************")
	fmt.Println("Welcome to Match Maker made with Go")
	fmt.Print("***********************************\n\n")
	fmt.Println("Here's how it works:")
	fmt.Println("You will be asked 5 questions and to each of the questions you must respond with a number between 1 - 5.")
	fmt.Println("1 being strongly disagree and 5 being strongly agree.")
	fmt.Println("Once finished, the computer will compare your answers and determine a compatibility score with a max value of 100.")
	fmt.Println("It will also tell you if it is true love, just friends, or run away.")
}

/*
Displays each question in the map and compares user ans with computer
generated answer to calculate score for each question and overall compatibility
*/
func displayQs(questions map[int]Question) int {
	scanner := bufio.NewScanner(os.Stdin)
	totalScore := 0

	//Loops through each question of the map
	for key, q := range questions {
		fmt.Println(q.Statement)
		fmt.Print("1 Strongly Disagree, 5 Strongly Agree: ")

		var userAns int
		for {
			scanner.Scan()
			input := scanner.Text()
			var err error
			userAns, err = validateAns(input)
			if err != nil {
				fmt.Println(err)
				fmt.Print("Please enter a valid number between 1 and 5: ")
				continue
			}
			break
		}
		questions[key] = Question{
			Statement: q.Statement,
			UserAns:   userAns,
			CompAns:   q.CompAns,
		}

		// Calculates compatibility score for the question
		questionScore := MaxScore - abs(q.CompAns-userAns)
		fmt.Printf("You scored %d on this question.\n\n", questionScore)
		totalScore += questionScore
	}
	return totalScore
}

//validates user answer to make sure input is valid
func validateAns(input string) (int, error) {
	score, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || score < MinScore || score > MaxScore {
		return 0, fmt.Errorf("invalid input")
	}
	return score, nil
}

//Helper function to calculate absolute value of a given int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
