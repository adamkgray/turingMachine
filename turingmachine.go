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

func display(input string, head int) {
	fmt.Println(input)
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

func main() {
	initial := flag.String("s", "s", "initial state symbol")
	halt := flag.String("h", "h", "halting state symbol")
	input := flag.String("i", "", "input tape")
	file := flag.String("t", "table.csv", "turing machine rules")
	flag.Parse()

	tape := "^" + *input

	m := machine{
		state: *initial,
		halt:  *halt,
		head:  1,
	}

	rules := readCsvFile(*file)

	table := make(map[string]instruction)
	for _, rule := range rules {
		// the key is the unique combination of
		// current state and character read
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

	clear()
	display(tape, m.head)
	wait(250)

	for {
		if m.state == m.halt {
			break
		}

		clear()

		rule := m.state + string(tape[m.head])
		result := table[rule]
		m.state = result.next

		if result.action == "<-" {
			if m.head == 0 {
				panic("attempted to move head past start of tape")
			}
			m.head--
		} else if result.action == "->" {
			if m.head == len(tape)-1 {
				tape += "_"
			}
			m.head++
		} else {
			tape = tape[:m.head] + result.action + tape[m.head+1:]
		}

		display(tape, m.head)
		wait(250)
	}
}
