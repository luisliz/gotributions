package main

import (
	"flag"
	"gotributions/internal"
)

type User struct {
	Uid      string
	Gid      string
	Username string
	Name     string
	HomeDir  string
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add new folder to scan git repos")
	flag.StringVar(&email, "email", "", "The email to scan")
	flag.Parse()

	if folder != "" {
		internal.Scan(folder)
		return
	}

	internal.Stats(email)
}
