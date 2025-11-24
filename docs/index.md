# Learn Go

A personal companion to *Learning Go* that mixes short-form notes with runnable code experiments. Each chapter mirrors the book structure so you can jump between narrative explanations and the matching Go programs under `chNN/`.

## Working with the Example Code

The repository mirrors the book's progression:

1. Each chapter lives in `chNN/` (for example, `ch03/`) and every topic or exercise resides in its own folder with a small `package main` program.
2. Run `make list` to see every runnable topic in the form `ch03-slices -> ./ch03/slices`.
3. Execute a single sample with `make run-<target>`, build everything with `make build-all`, or run `make ci` to format, vet, test, and build before commits.

## Building the Documentation

This documentation is built using [MkDocs](https://www.mkdocs.org/) with the [Material theme](https://squidfunk.github.io/mkdocs-material/).

To build and serve the documentation locally:

```bash
make docs        # Build static site to site/
make docs-serve  # Start live-reload dev server at http://127.0.0.1:8000
```

## Chapters

- [Chapter 1: Setting Up Your Environment](ch01.md)
- [Chapter 2: Predeclared Types and Declarations](ch02.md)
- [Chapter 3: Composite Types](ch03.md)
