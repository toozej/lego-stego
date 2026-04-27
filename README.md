# lego-stego

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/toozej/lego-stego)
[![Go Report Card](https://goreportcard.com/badge/github.com/toozej/lego-stego)](https://goreportcard.com/report/github.com/toozej/lego-stego)
![GitHub Actions CI Workflow Status](https://img.shields.io/github/actions/workflow/status/toozej/lego-stego/ci.yaml)
![GitHub Actions Release Workflow Status](https://img.shields.io/github/actions/workflow/status/toozej/lego-stego/release.yaml)
![GitHub Actions Weekly Docker Refresh Workflow Status](https://img.shields.io/github/actions/workflow/status/toozej/lego-stego/weekly-docker-refresh.yaml)
![Docker Pulls](https://img.shields.io/docker/pulls/toozej/lego-stego)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/toozej/lego-stego/total)

Go-based CLI tool for [Steganography](https://en.wikipedia.org/wiki/Steganography)

<img src="img/avatar.png" alt="lego-stego avatar" style="background-color: #FFFFFF;" />

## features of this starter template
- follows common Golang best practices in terms of repo/project layout, and includes explanations of what goes where in README files
- Cobra library for CLI handling, Logrus for logging, and GoDotEnv and Env libraries for reading config files already plugged in and ready to expand upon
- Goreleaser to build Docker images and most standard package types across Linux, MacOS and Windows
    - also includes auto-generated manpages and shell autocompletions
- Makefile for easy building, deploying, testing, updating, etc. both Dockerized and using locally installed Golang toolchain
- docker-compose project for easily hosting built Dockerized Golang project, with optional support for Golang web services
- scripts to make using the starter template easy, and to update the Golang version when a new one comes out
- Dev Container with built in Go-related VSCode extensions, and [llm](https://llm.datasette.io/) tool + plugins pre-configured to use GitHub Copilot
- built-in security scans, vulnerability warnings and auto-updates via Dependabot and GitHub Actions
- auto-generated documentation
- pre-commit hooks for ensuring formatting, linting, security checks, etc.

## changes required to use this as a starter template
- set up new repository in quay.io web console
    - (DockerHub and GitHub Container Registry do this automatically on first push/publish)
    - name must match Git repo name
    - grant robot user with username stored in QUAY_USERNAME "write" permissions (your quay.io account should already have admin permissions)
- set built packages visibility in GitHub packages to public
    - navigate to https://github.com/users/$USERNAME/packages/container/$REPO/settings
    - scroll down to "Danger Zone"
    - change visibility to public

