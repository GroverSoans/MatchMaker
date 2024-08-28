package main
import "fmt"
import "math/rand"

type Question struct {
	Statement string
	UserAns int
	CompAns int
}

var questions = map[int]Question{
	1: {
		Statement: "I think Sundays are perfect for lazy mornings and breakfast in bed.",
		UserAns: 0,
		CompAns: rand.Intn(5) + 1,
	},
	2: {
		Statement: "I love spontaneous road trips and discovering new places.",
		UserAns: 0,
		CompAns: rand.Intn(5) + 1,
	},
	3: {
		Statement: "I think dogs are the best companions.",
		UserAns: 0,
		CompAns: rand.Intn(5) + 1,
	},
	4: {
		Statement: "Broccoli is totally underrated.",
		UserAns: 0,
		CompAns: rand.Intn(5) + 1,
	},
	5: {
		Statement: "I think pineapple belongs on pizza.",
		UserAns: 0,
		CompAns: rand.Intn(5) + 1,
	},
}

func main(){
	welcomeMsg()


}
func welcomeMsg(){
	fmt.Println("***********************************")
	fmt.Println("Welcome to Match Maker made with Go")
	fmt.Print("***********************************\n\n")
	fmt.Println("Heres how it works:")
	fmt.Println("You will be asked 5 questions and to ")
	fmt.Println("each of the questions you must respond")
	fmt.Println("with a number between 1 - 5. 1 being")
	fmt.Println("strongly disagree and 5 being strongly")
	fmt.Println("agree. Once finished the computer will")
	fmt.Println("compare your answers and determine a")
	fmt.Println("compatiblity score that has a max value")
	fmt.Println("of 100. It will also tell you if it is")
	fmt.Print("true love, just friends, or run away.\n")
}

