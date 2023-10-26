# GoReleaser

I'm going to use this repository to integrate and replace my Makefile with GoReleaser

## Initial state

To automate some of the steps, I'm using a good ol' Makefile

To compile this software in amd64, I just have to run:

```
VERSION=0.0.1 make build
```

To create a docker image for this, I need to run:

```
TAG=0.0.1 make dockerbuild
```

To push the newly created docker image to hub.docker.com, I need to run:

```
TAG=0.0.1 make dockerpush
```

That's better than doing `go build`s, `docker build`s and `docker push`es by hand, but not by much...

## Enter GoReleaser

[GoReleaser](https://goreleaser.com/) is an open source software that aims to automate most boring tasks from shipping go software. 

> With GoReleaser, you can:
>
>    Cross-compile your Go project
>    Release to GitHub, GitLab and Gitea
>    Create nightly builds
>    Create Docker images and manifests
>    Create Linux packages and Homebrew taps
>    Sign artifacts, checksums and container images
>    Announce new releases on Twitter, Slack, Discord and others
>    Generate SBOMs (Software Bill of Materials) for binaries and container images
>    ... and much more!

There is a Pro version that supports the developer, but we will only use the opencore one in this document.

## Initialize GoReleaser

You can generate a basic working GoReleaser config file with the command

```bash
$ goreleaser init
  • Generating .goreleaser.yaml file
  • config created; please edit accordingly to your needs file=.goreleaser.yaml
  • thanks for using goreleaser!
```

This file is a good start but for the sake of the demonstration, I'll start with my own simpler file first.

## 1. build

I'm going to replace the first parts of the Makefile

```makefile
prepare:
	go mod tidy

build: prepare
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/tifling -ldflags "-X main.Version=$$VERSION" main.go
```

With GoReleaser, it should look like this

```yaml
before:
  hooks:
    - go mod tidy

builds:
  - binary: bin/tifling
    env:
      - CGO_ENABLED=0
    goos:
      - linux
    goarch:
      - "amd64"
    ldflags:
      - -X main.Version=={{.Version}}
```

Section before.hooks will allow us to do some things before building, just like we did with the "build: prepare" line in the Makefile.

`{{.Version}}` is an automated variable from GoReleaser. We'll use it to get rid of the manual task of setting the `$$VERSION` and `$$TAG` from our Makefile. GoReleaser is going to leverage git tags for this.

Note: there are many variables and they will be super useful in many usecases. 

Let's now commit our new .goreleaser.yml file, add a git tag and try to run goreleaser.

```
git add .
git commit -m "add simplest goreleaser example"
git tag -a 0.0.1 -m "simplest goreleaser example"
git push origin 0.0.1
```
