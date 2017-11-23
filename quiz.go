package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "The path to the questions & answers csv file")
	flag.Parse()

	data, err := ioutil.ReadFile(*csvFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Go quiz yourself! ", *csvFile)

	dataStr := strings.Split(string(data), "\n")

	correctCounter := 0
	expr := make([]string, 0)

	for _, v := range dataStr {
		if strings.TrimSpace(v) != "" {
			expr = append(expr, v)
		}
	}

	for _, v := range expr {
		parts := strings.Split(v, ",")

		fmt.Printf("%s= ", parts[0])
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if strings.TrimSpace(text) == parts[1] {
			correctCounter++
		}
	}

	fmt.Printf("You answered %d/%d correctly", correctCounter, len(expr))
}
