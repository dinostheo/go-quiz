package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func exit(count int, total int) {
	fmt.Printf("\nYou answered %d/%d correctly", count, total)

	os.Exit(0)
}

func shuffleQuestions(qas []string) []string {
	l := len(qas)

	for i := range qas {
		j := rand.Intn(l)
		qas[i], qas[j] = qas[j], qas[i]
	}

	return qas
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	csvFile := flag.String("csv", "problems.csv", "The path to the questions & answers csv file")
	limit := flag.Int("limit", 30, "The time to run the quiz in seconds (default is 30 seconds)")
	shuffle := flag.Bool("shuffle", false, "Shuffle the questions")

	flag.Parse()

	data, err := ioutil.ReadFile(*csvFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dataStr := strings.Split(string(data), "\n")

	correctCounter := 0
	expr := make([]string, 0)

	for _, v := range dataStr {
		if strings.TrimSpace(v) != "" {
			expr = append(expr, v)
		}
	}
	fmt.Println(">>>>>> ", *shuffle)
	if *shuffle {
		expr = shuffleQuestions(expr)
	}

	go func() {
		time.Sleep(time.Duration(*limit) * time.Second)
		fmt.Println("\nTime's up!!!")

		exit(correctCounter, len(expr))
	}()

	for _, v := range expr {
		parts := strings.Split(v, ",")

		fmt.Printf("%s= ", parts[0])
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if strings.ToLower(strings.TrimSpace(text)) == strings.ToLower(parts[1]) {
			correctCounter++
		}
	}

	exit(correctCounter, len(expr))
}
