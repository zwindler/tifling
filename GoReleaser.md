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

Let's now commit our new .goreleaser.yml file, add a git tag:

```
$ git add .
$ git commit -m "add simplest goreleaser example"
$ git tag -a 0.0.1 -m "simplest goreleaser example"
$ git push origin 0.0.1
```

And then run goreleaser

```
$ goreleaser build --clean
  • starting build...
  • loading                                          path=.goreleaser.yaml
  • loading environment variables
  • getting and validating git state
    • couldn't find any tags before "0.0.1"
    • git state                                      commit=2778c1e356e1b474ca9fe3081e450a3bd8a430e3 branch=main current_tag=0.0.1 previous_tag=<unknown> dirty=false
  • parsing tag
  • setting defaults
  • running before hooks
    • running                                        hook=go mod tidy
  • checking distribution directory
  • loading go mod information
  • build prerequisites
  • writing effective config file
    • writing                                        config=dist/config.yaml
  • building binaries
    • building                                       binary=dist/tifling_linux_amd64_v1/bin/tifling
  • storing release metadata
    • writing                                        file=dist/artifacts.json
    • writing                                        file=dist/metadata.json
  • build succeeded after 0s
  • thanks for using goreleaser!
```

We can now check if our binary is functional, and version is automatically set by goreleaser using the git tag:

```
$ dist/tifling_linux_amd64_v1/bin/tifling
## tifling version =0.0.1

Random Entry:
- Name: Coiff'hair
- Latitude/Longitude: 48.556123 -2.019321
```

Should we try to execute goreleaser without tagging a branch first, we will get an error:

```
  ⨯ build failed after 0s                    error=git is in a dirty state
```

or 

```
  ⨯ build failed after 0s                    error=git tag 0.0.1 was not made against commit 4a03b27dd1369cbeb2451d362493756fc90d7d48
```

This is normal, goreleaser relies on tags to work. In case we really want to use goreleaser and not tag the code, we call use the `--snapshot` option. This will be useful for snapshot/nightly releases for example:

```
$ goreleaser build --clean --snapshot
[...]
  • snapshotting
    • building snapshot...                           version=0.0.1-SNAPSHOT-4a03b27
[...]

$ dist/tifling_linux_amd64_v1/bin/tifling
## tifling version =0.0.1-SNAPSHOT-4a03b27
```

## 2. crossbuild

Now that we reproduced the basic functionality of the initial Makefile, we can try to go further.

Rather than limiting our build to amd64 on linux only, we can remove the `goos` and `goarch` sections completely to tell goreleaser to crossbuild to the most common archs. We also can keep those and explicitly define those we want.

```yaml
before:
  hooks:
    - go mod tidy

builds:
  - binary: bin/tifling
    env:
      - CGO_ENABLED=0
    ldflags:
      - -X main.Version=={{.Version}}
```

If we try to run goreleaser again, we will build much our binary for many more targets out of the box, which the need to write a loop in the Makefile or in a script:

```
$ git add .
$ git commit -m "goreleaser add more targets"
$ git tag -a 0.0.2 -m "goreleaser add more targets"
$ git push origin 0.0.2

$ goreleaser build --clean
[...]

$ ls dist
artifacts.json  tifling_darwin_amd64_v1  tifling_linux_amd64_v1  tifling_windows_amd64_v1
config.yaml     tifling_darwin_arm64     tifling_linux_arm64     tifling_windows_arm64
metadata.json   tifling_linux_386        tifling_windows_386
```

## 3. Push it somewhere

