package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/howeyc/gopass"
	"github.com/jessevdk/go-flags"
	"golang.org/x/crypto/ssh/terminal"
)

func newShadow(args []string) *shadow {
	var s shadow
	parser := flags.NewParser(&s, flags.None)
	parser.Usage = "[options] [name]"
	args, err := parser.ParseArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	if !s.SHA256 && !s.SHA512 && !s.MD5 {
		// default
		s.SHA512 = true
	}

	if s.Help {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if !s.OnlyEncrypt {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Login name is require.")
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
		s.LoginName = args[0]
	}

	return &s
}

func getPassword() []byte {
	if terminal.IsTerminal(0) {
		fmt.Fprintf(os.Stderr, "Enter Password: ")
		pass, err := gopass.GetPasswd()
		if err != nil {
			log.Fatal(err)
		}
		return bytes.TrimSpace(pass)
	} else {
		pass, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		return bytes.TrimSpace(pass)
	}
}

func main() {
	s := newShadow(os.Args[1:])
	password := getPassword()
	if s.OnlyEncrypt {
		fmt.Println(s.encryptedPassword(password))
	} else {
		fmt.Println(s.make(password))
	}
}
