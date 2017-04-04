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

func checkArg(arg string, msg string) {
	if arg == "" {
		fmt.Println(msg)
		os.Exit(-1)
	}
}

func parseFlags() *config {
	cfg := new(config)

	flag.StringVar(&cfg.owner, "owner", "", "GitHub owner")
	flag.StringVar(&cfg.repo, "repo", "", "GitHub repo")
	flag.StringVar(&cfg.token, "token", "", "GitHub token")

	flag.Parse()

	checkArg(cfg.owner, "Missing GitHub owner")
	checkArg(cfg.repo, "Missing Github repo")
	checkArg(cfg.token, "Missing GitHub token")

	cfg.sshUrl = fmt.Sprintf("git@github.com:%v/%v.git", cfg.owner, cfg.repo)

	return cfg
}

func newClient(ctx context.Context, cfg *config) *github.Client {
	sts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.token})
	return github.NewClient(oauth2.NewClient(ctx, sts))
}

func listForks(ctx context.Context, cfg *config, client *github.Client) []*github.Repository {
	forks, _, err := client.Repositories.ListForks(ctx, cfg.owner, cfg.repo, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return forks
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
	cfg := parseFlags()
	ctx := context.Background()
	client := newClient(ctx, cfg)
	forks := listForks(ctx, cfg, client)

	execCommand("git", []string{"clone", cfg.sshUrl}...)
	execCommand("git", []string{"-C", cfg.repo, "remote", "rename", "origin", "upstream"}...)

	for _, fork := range forks {
		execCommand("git", []string{"-C", cfg.repo, "remote", "add", *fork.Owner.Login, *fork.SSHURL}...)
	}
}
