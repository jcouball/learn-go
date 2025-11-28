# Learn Go

A personal companion to *Learning Go* that mixes short-form notes with runnable code experiments. Each chapter mirrors the book structure so you can jump between narrative explanations and the matching Go programs under `chNN/`.

## Installation

### Prerequisites

Install the following tools to work with this repository:

- **Go** - The Go programming language
- **Task** - Modern task runner (replaces Make)
- **Python 3** - Required for MkDocs documentation generator

The instructions below use [Homebrew](https://brew.sh/) for macOS, but you can install these tools using any package manager or installation method you prefer.

**Go** - The Go programming language

```bash
# macOS with Homebrew
brew install go

# Or download from https://go.dev/dl/
```

**Task** - Modern task runner (replaces Make)

```bash
# macOS with Homebrew
brew install go-task/tap/go-task

# Or using Go
go install github.com/go-task/task/v3/cmd/task@latest
```

**Python 3** - Required for MkDocs documentation generator

```bash
# macOS with Homebrew
brew install python

# Or use asdf, pyenv, or download from https://www.python.org/
```

### Setup

After cloning the repository, run the setup task to verify dependencies and initialize the project:

```bash
task setup
```

This will:

- Verify that Go and Python 3 are installed
- Install MkDocs and Python dependencies from `requirements.txt`
- Install Go dependencies with `go mod tidy`

## Working with the Example Code

The repository mirrors the book's progression:

1. Each chapter lives in `chNN/` (for example, `ch03/`) and every topic or exercise resides in its own folder with a small `package main` program.
2. Run `task list-examples` to see all runnable packages and their binary names.
3. Execute a single sample with `task run -- ch03/slices`, build everything with `task build-examples`, or run `task ci` to format, vet, test, and build before commits.

### Common Commands

```bash
# List all available tasks
task

# List all example packages
task list-examples

# Run a specific package
task run -- ch03/slices

# Build a specific package
task build -- ch03/slices

# Build all examples
task build-examples

# Run all examples
task run-examples

# Development workflow
task fmt         # Format code
task vet         # Run go vet
task test        # Run tests with coverage
task ci          # Run fmt, vet, test, and build-examples

# Clean build artifacts
task clean
```

## Building the Documentation

This documentation is built using [MkDocs](https://www.mkdocs.org/) with the [Material theme](https://squidfunk.github.io/mkdocs-material/).

To build and serve the documentation locally:

```bash
task docs        # Build static site to site/
task docs-serve  # Start live-reload dev server at http://127.0.0.1:8000
```

The documentation is organized into chapters that mirror the book structure. See the `docs/` directory for the source files.
