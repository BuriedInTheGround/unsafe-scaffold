package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"filippo.io/age"
	"golang.org/x/term"
)

var usage = `Usage: age-encrypt FILE
Symmetrically encrypts a file with a passphrase using age.

The passphrase is requested interactively via standard input. The encrypted
file goes to standard output.
`

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()

	fmt.Fprint(os.Stdin, "Enter passphrase: ")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Fprint(os.Stdin, "\n")

	r, err := age.NewScryptRecipient(string(pass))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	r.SetWorkFactor(16)

	name := flag.Arg(0)
	if name == "" {
		log.Fatalf("error: no input file specified", err)
	}
	in, err := os.Open(name)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer in.Close()

	var out io.Writer = os.Stdout

	w, err := age.Encrypt(out, r)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if _, err := io.Copy(w, in); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
