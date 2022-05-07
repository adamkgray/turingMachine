package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Parse command line args
	initial := flag.String("s", "s", "initial state symbol")
	halt := flag.String("h", "h", "halting state symbol")
	input := flag.String("i", "", "string representing input tape")
	file := flag.String("t", "table.csv", "path to CSV file of turing machine rules")
	flag.Parse()

	// Prepend the input tape with '^' for clearer visualisation
	tape := "^" + *input

	// Init the turing machine struct with the initial tape, halting symbol and the head at position 1
	m := machine{
		state: *initial,
		halt:  *halt,
		head:  1,
	}

	// Read the state transition table
	rules := readCsvFile(*file)

	// Parse the state transition table into a map of instructions
	table := make(map[string]instruction)
	for _, rule := range rules {
		// The key is the unique combination of the current state and the character read
		current := rule[0]
		read := rule[1]
		next := rule[2]
		action := rule[3]
		key := current + read
		table[key] = instruction{
			next:   next,
			action: action,
		}
	}

	// Clear the terminal, display the input tape, and wait
	clear()
	display(tape, m.head)
	wait(250)

	for {
		// Exit when the turing machine reaches the halting state
		if m.state == m.halt {
			break
		}

		// Clear the terminal
		clear()

		// Read the state and the character under the head together as a string
		rule := m.state + string(tape[m.head])
		// Given that string, pull the associated rules from the table
		result := table[rule]
		// Determine the next state
		m.state = result.next

		if result.action == "<-" { // Move the head to the left
			if m.head == 0 { // Throw error if the state transition table is buggy
				panic("attempted to move head past start of tape")
			}
			// Decrement the head position
			m.head--
		} else if result.action == "->" { // Move the head to the right
			if m.head == len(tape)-1 { // If the head moves past the end of the tape, we treat it as adding a blank space to the end of the tape
				tape += "_"
			}
			// Increment the head position
			m.head++
		} else { // Update or erase the character under the head
			tape = tape[:m.head] + result.action + tape[m.head+1:]
		}

		// Display the new state of the tape and wait
		display(tape, m.head)
		wait(250)
	}
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read table file '%s': %s", filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV for '%s': %s", filePath, err)
	}
	return records
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func display(tape string, head int) {
	fmt.Println(tape)
	fmt.Println(strings.Repeat(" ", head) + "@")
}

func wait(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

type machine struct {
	state string
	halt  string
	head  int
}

type instruction struct {
	next   string
	action string
}
