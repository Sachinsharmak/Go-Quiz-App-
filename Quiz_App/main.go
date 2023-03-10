package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	//  read the proble from the quiz csv file
	// 1.open the file
	if fObj, err := os.Open(fileName); err == nil {
		// 2. we will create a new reader
		csvR := csv.NewReader(fObj)
		// 3. it will need to read the file
		if cLines, err := csvR.ReadAll(); err == nil {
			// 4. call the parse function
			return ParseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("Erroe in Reading the File"+"Format from %s file; %s", fileName, err.Error())
		}

	} else {
		return nil, fmt.Errorf("Error in Opening the %s File; %s", fileName, err.Error())
	}

}
func main() {
	// 1. input the name of the file
	fName := flag.String("f", "Quiz.csv", "Path of the Csv File")
	// 2. set the duration of the timer
	timer := flag.Int("t", 30, "this is the timer of the quiz")
	flag.Parse()
	// 3. Pull the proble from the file ( calling Problempuller)
	problems, err := problemPuller(*fName)
	// 4. Handle the error
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s\n", err.Error()))
	}
	// 5. create Variable to count the correct answers
	correctAns := 0
	// 6. using the duration of the timer , we want to initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	// 7. loops through the problems and ask the user the question
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println("Time is Up")
			break problemLoop
		case answer = <-ansC:
			if answer == p.a {
				correctAns++
			}
			if i == len(problems) {
				close(ansC)
			}
		}
	}
	// 8. print the number of correct answers
	fmt.Printf("You scored %d out of %d\n", correctAns, len(problems))
	fmt.Println("Thank you for playing the quiz")
	fmt.Printf("Press Enter the Exit The Quiz")
	<-ansC
}

func ParseProblem(lines [][]string) []problem {
	// go over the lines and parse them ,with problem struct
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
