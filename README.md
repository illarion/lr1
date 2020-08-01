# lr1 parser 

This project is an example of LR(1) parser, and is useful for quick and dirty parsing of files that you don't have parser library for.
It consists of 2 elements -- `lr1.Parse(items []interface{}, rules ...Rule) ([]interface{}, error)` function and a number of rules `Rule(items []interface{}) (interface{}, error)`. You have to specify rules yourself. 

Algorithm is extremely simple - given the number of items, initially it may be just a list of tokenized strings from some text, Parse() takes last 1 element, then last 2 elements, 3, 4 and so on, and applies all specified rules (one by one) to this sublist. If some rule is applicable and returns the new value (`[]interface{}`), parser replaces the current end of the list with the replacement provided by rule. 

For example, let's say we parse following text, consisting of some imaginary language function calls:

```
foo(1,2);
bar(a,b,c);
```

## Step 1 - tokenize like this (see https://golang.org/pkg/text/scanner/ for more details) : 
`["foo", "(", "1", ",", "2", ")", ";", "bar", "(", "a", ",", "b", ",", "c", ")", ";"]`

## Step 2 - implement following rule functions
```
[token, ARGS, SEMICOLON] --> struct Function{ name: ..., }
["(", END_OF_LIST] -> ARGS
[token, ",", END_OF_LIST] -> END_OF_LIST // with appended token
[")"] -> END_OF_LIST`
[";"] -> SEMICOLON
```

## Step 3 - run the Parser
Result - list of Function{} structs, parsed from the initial list of string tokens, or error if it's not possible to parse


## Links
https://en.wikipedia.org/wiki/Canonical_LR_parser
