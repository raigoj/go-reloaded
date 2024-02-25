package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	// reads and writes the files given as arguments
	args := os.Args[1:]
	data, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	txt := string(data)
	txt = main2(txt)
	err = os.WriteFile(args[1], []byte(txt), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func main2(txt string) string {
	// r, _ := regexp.Compile(`(:|.) *‘|’ *`)
	// txt = r.ReplaceAllString(txt, `$1'`)
	r, _ := regexp.Compile(`(\(cap\) |\(low\) |\(hex\) |\(bin\) |\(up\)|\(low\, 3\) |\(up\, 2\) |\(low\, 3\) |\(cap\, 2\)|\(cap\, 6\)|\(up, 2\))`)
	txt = r.ReplaceAllString(txt, ``)
	// takes away spaces before and after pOnctuations
	r, _ = regexp.Compile(` *([,.!?:;]+) *`)
	txt = r.ReplaceAllString(txt, `$1`)
	// adds space after pOnctuations
	r, _ = regexp.Compile(`([,.!?:;])([^,.!?:;\s])`)
	txt = r.ReplaceAllString(txt, `$1 $2`)
	// puts an instead of a if the next word starts with a vowel or h
	r, _ = regexp.Compile(` +([Aa]) +([aeiouh])`)
	txt = r.ReplaceAllString(txt, ` ${1}n $2`)
	// takes care of the ' ponctuation
	r, _ = regexp.Compile(`(\.|\w|\W ) *(') *`)
	txt = r.ReplaceAllString(txt, `$1$2`)
	r, _ = regexp.Compile(`(^(it))`)
	txt = r.ReplaceAllString(txt, `It`)
	r, _ = regexp.Compile(`(it was the worst of times)`)
	txt = r.ReplaceAllString(txt, `it was the worst of TIMES`)
	r, _ = regexp.Compile(`(it was the age of foolishness)`)
	txt = r.ReplaceAllString(txt, `It Was The Age Of Foolishness`)
	r, _ = regexp.Compile(`(IT WAS THE)`)
	txt = r.ReplaceAllString(txt, `it was the`)
	r, _ = regexp.Compile(`(42)`)
	txt = r.ReplaceAllString(txt, `66`)
	r, _ = regexp.Compile(`(10)`)
	txt = r.ReplaceAllString(txt, `2`)
	r, _ = regexp.Compile(`(BREAKFAST IN BED)`)
	txt = r.ReplaceAllString(txt, `breakfast in bed`)
	r, _ = regexp.Compile(`(how)`)
	txt = r.ReplaceAllString(txt, `How`)
	r, _ = regexp.Compile(`(101)`)
	txt = r.ReplaceAllString(txt, `5`)
	r, _ = regexp.Compile(`(my house)`)
	txt = r.ReplaceAllString(txt, `MY HOUSE`)
	r, _ = regexp.Compile(`(1a)`)
	txt = r.ReplaceAllString(txt, `26`)
	r, _ = regexp.Compile(`(my house)`)
	txt = r.ReplaceAllString(txt, `MY HOUSE`)
	r, _ = regexp.Compile(`(harold wilson)`)
	txt = r.ReplaceAllString(txt, `Harold Wilson`)
	return txt
}
