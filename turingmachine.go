package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func display(input string, head int) {
	fmt.Println(input)
	fmt.Println(strings.Repeat(" ", head) + "@")
}

func wait(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
	// convert csv to map
	table := make(map[string][]string)
	for _, record := range readCsvFile("table.csv") {
		table[record[0]+record[1]] = []string{record[2], record[3]}
	}

	args := os.Args
	state := args[1]
	halt := args[2]
	input := "^" + args[3]
	head := 1

	clear()
	display(input, head)
	wait(250)

	for {
		// end program when the halting state has been reached
		if state == halt {
			break
		}

		clear()

		instruction := table[state+string(input[head])]
		state = instruction[0]

		if instruction[1] == "<-" {
			if head == 0 {
				log.Fatal("cannot move head to the left - already at the start of input tape")
				os.Exit(1)
			}
			head = head - 1
		} else if instruction[1] == "->" {
			if head == len(input)-1 {
				input = input + "_"
			}
			head = head + 1
		} else {
			input = input[:head] + instruction[1] + input[head+1:]
		}

		display(input, head)
		wait(250)
	}
}
