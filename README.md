# Turing Machine

A command-line utility to write and visualise the execution of turing machines.

## Usage

Compile with `go build` to produce the executable called `turingmachine`. Running the executable with the `-h` or `--help` flag prints the help page:

```bash
$ ./turingmachine --help
Usage of ./turingmachine:
  -e string
        halting state symbol (default "h")
  -i string
        string representing input tape
  -s string
        initial state symbol (default "s")
  -t string
        path to CSV file of turing machine rules (default "table.csv")
```

- By default, the halting symbol is `h`, but you can override this behavior with the `-e` flag, i.e. `-e H` to make `H` the halting symbol
- By default, the initial state symbol is `s`, but you can override this behavior with the `-s` flag, i.e. `-s S` to make `S` the halting symbol
- By default, the program looks for the state transition table in a file called `table.csv`, although you can tell it to look anywhere by specifying a path after the `-t` flag ("t" for table), i.e. `-t /path/to/table.csv`
- You must specify an input string with the `-i` flag, i.e. `-i 10101101`

## The state transition table

The state transition table is given by a CSV file. Refer to the following rules to create your own:

- The columns from left to right are: state, character read, next state, action
- Do not include the column names in the csv
- To read a blank (space), specify `_` in the character read column
- To erase a charcater on the tape, specify `_` in the action column
- To write a character on the tape, for example `1`, specify `1` in the action column
- To move the head right, specify `->` in the action column
- To move the head left, specify `<-` in the action column

Here is an example program which appends `010` to any binary number input greater than `0`:
```
s,0,n,_
n,_,n,1
n,1,m,->
m,_,h,0
s,1,z,<-
z,^,z,->
z,0,z,->
z,1,z,->
z,_,y,0
y,0,x,->
x,_,w,1
w,1,v,->
v,_,h,0
```
