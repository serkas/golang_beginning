package main

import (
	"flag"
	"fmt"
	"regexp"
)

var (
	// for static expressions
	emailRegExp = regexp.MustCompile(`^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$`) // https://www3.ntu.edu.sg/home/ehchua/programming/howto/Regexe.html#zz-1.9
)

func main() {
	emails := []string{
		"lama.loca.loca123@inca.com",
		"pio_pio@factory.com",
		"carnival666@hellmail.com",
		"la-lalai@gmail.com",
		" testmail@mail.com",
		".testmail@mail.com",
		"testmail@mail.com.",
		"testm ail@mail.com",
	}

	for _, e := range emails {
		isValid := emailRegExp.MatchString(e)
		fmt.Printf("%q %t\n", e, isValid)
	}

	//For dynamic expressions
	//go run . -pattern '\d{3}\s\w+'
	pattern := flag.String("pattern", "", "specify pattern to search for")
	flag.Parse()
	if *pattern == "" {
		flag.Usage()
		return
	}

	re, err := regexp.Compile(*pattern)
	if err != nil {
		fmt.Printf("compiling pattern: %s", err)
		return
	}

	text := `some text example 123 abc`
	if re.MatchString(text) {
		fmt.Printf("%q MATCHED %q\n", *pattern, text)
	} else {
		fmt.Printf("%q NOT MATCHED %q\n", *pattern, text)
	}

}
