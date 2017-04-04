# cmote

## Introduction

If you're using [GitHub flow](https://guides.github.com/introduction/flow/index.html) on a team, cmote can help you easily set up repositories locally.

With it, you can clone an upstream repository, rename the provided remote from `origin` to `upstream`, and add a remote for every fork that exists on GitHub in one step. This makes it easy to set up a new project where you anticipate interacting with the existing remotes without needing to repeatedly run `git remote add $FORK`.

You probably don't want to use cmote with larger open source projects with lots of forks.

## Quick Start

Before you start, you'll need to [install Go](https://golang.org/doc/install) and create a personal [GitHub auth token](https://github.com/blog/1509-personal-api-tokens).

```bash
$ go install github.com/mmwtsn/cmote
$ cmote --owner $OWNER --repo $REPO --token $TOKEN
```
