package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type config struct {
	owner  string
	repo   string
	sshUrl string
	token  string
}

func execCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
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

	cfg.sshUrl = fmt.Sprintf("git@github.com:%v/%v.git", cfg.owner, cfg.repo)

	ctx := context.Background()
	sts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.token})
	client := github.NewClient(oauth2.NewClient(ctx, sts))

	forks, _, err := client.Repositories.ListForks(ctx, cfg.owner, cfg.repo, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	execCommand("git", []string{"clone", cfg.sshUrl}...)
	execCommand("git", []string{"-C", cfg.repo, "remote", "rename", "origin", "upstream"}...)

	for _, fork := range forks {
		execCommand("git", []string{"-C", cfg.repo, "remote", "add", *fork.Owner.Login, *fork.SSHURL}...)
	}
}
