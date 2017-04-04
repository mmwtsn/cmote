package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	owner string
	repo  string
	token string
}

func main() {
	cfg := new(config)

	flag.StringVar(&cfg.owner, "owner", "", "GitHub owner")
	flag.StringVar(&cfg.repo, "repo", "", "GitHub repo")
	flag.StringVar(&cfg.token, "token", "", "GitHub token")

	flag.Parse()

	if cfg.owner == "" {
		fmt.Println("Missing GitHub owner")
		os.Exit(-1)
	}

	if cfg.repo == "" {
		fmt.Println("Missing GitHub repo")
		os.Exit(-1)
	}

	if cfg.token == "" {
		fmt.Println("Missing GitHub token")
		os.Exit(-1)
	}
}
