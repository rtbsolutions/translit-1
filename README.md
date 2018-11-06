# translit [![GoDoc](https://godoc.org/github.com/mxmCherry/translit?status.svg)](https://godoc.org/github.com/mxmCherry/translit) [![Build Status](https://travis-ci.org/mxmCherry/translit.svg?branch=master)](https://travis-ci.org/mxmCherry/translit) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/translit)](https://goreportcard.com/report/github.com/mxmCherry/translit) [![codecov](https://codecov.io/gh/mxmCherry/translit/branch/master/graph/badge.svg)](https://codecov.io/gh/mxmCherry/translit)

Go (Golang) utilities for (mostly Cyrillic) transliteration.

PROJECT STATE: pre-v1 grooming (mostly README).

This project aims to provide:

- easy-to-configure transliteration [golang.org/x/text/transform.Transformer](https://godoc.org/golang.org/x/text/transform#Transformer)
- language-specific transliterations (only Ukrainian/National and Russian/ICAO so far, PRs/issues are welcome)

This project is intended to be used with [golang.org/x/text/transform](https://godoc.org/golang.org/x/text/transform) - really well-thought-out/convenient base for streaming text transforming (at least, it abstracts all the tricky buffering).

If you a) don't need to build a custom transliterator b) are fine with custom license - take a look at [essentialkaos/translit](https://github.com/essentialkaos/translit) - it has plenty (like, really plenty!) of standards implemented for Cyrillic and should be very fast.

## Features

- simple, but flexible configuration - `translit.New` accepts string-to-string transliteration rule map, so it should be really easy to implement/add custom transliterations
- properly handles multi-char rules, would work for latin-to-cyrillic transliterations like `z` -> `з` and `zh` -> `ж` (and not `zh` -> `зх`)
- expected decent performance - no regular expressions or anything like that (though that could actually be quite fast for transliteration use case)
- [The Unlicense](https://tldrlegal.com/license/unlicense): Anyone is free to copy, modify, publish, use, compile, sell, or distribute this software, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.

## Problems

- not flexible enough to differentiate transliteration rules for first/non-first word letter for now: for example, Ukrainain national transliteration standard differentiates `я` at the beginning of the word `ya` and `я` in the middle of the word `ia` (though this looks solveable, just a bit later; should be fine for common usage even with this "disadvantage")

## Usage

```shell
go get -u github.com/mxmCherry/translit
```

```go
package translit_test

import (
	"fmt"

	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
)

func ExampleNew() {
	tr := translit.New(map[string]string{
		"л":  "l",
		"Л":  "L",
		"ля": "lya",
		"Ля": "Lya",
	})

	var s string

	s, _, _ = transform.String(tr, "Л - л")
	fmt.Println(s) // L - l

	s, _, _ = transform.String(tr, "Ля-лЯ-ля")
	fmt.Println(s) // Lya-lЯ-lya

	// Output:
	// L - l
	// Lya-lЯ-lya
}
```

## Guidelines

TODO: still subject to change!

This package aims to provide default transliterations for some languages.

Subpackage names for these transliterations should be made of [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code and the standard name, for example: `uknational`, where `uk` is the language code and `national` is a standard, defined by national government.

One subpackage per language/standard recommendation is given reduce memory footprint: pay (with memory) only for what you actually use.

These subpackages should expose at least one [transform.Transformer](https://godoc.org/golang.org/x/text/transform#Transformer) constructor, ideally - two (two-way transformers, like `ToLatin`/`FromLatin`).

```go
package main

import "github.com/mxmCherry/translit/uknational"

var uk = uknational.ToLatin() // global one (package-local)

func main() {
	s, _ _ := transform.String(uk, "Український трансліт")
	println(s) // Ukrainskyi translit
}
```

### Code style

- [editorconfig](https://editorconfig.org/) (recommended)
- [gofmt](https://blog.golang.org/go-fmt-your-code) (MUST)
- [goreportcard](https://goreportcard.com/report/github.com/mxmCherry/translit) (MUST)
- [codecov](https://codecov.io/gh/mxmCherry/translit) (recommended)

## Motivation

At the moment of writing this lib, there were 3 analogs:

[gen1us2k/go-translit](github.com/gen1us2k/go-translit) at [46f1a0b](https://github.com/gen1us2k/go-translit/commit/46f1a0be552caadbdbc19cf9a6705c4402b7ab47):

- bad: only Latin -> Russian (depends on use case though)
- minor: unable to reuse for other languages
- very bad: anti-optimal at least  - looping through rules, compiling regexp and replacing entire string for each rule

[essentialkaos/translit](https://github.com/essentialkaos/translit) at [d62c0f9](https://github.com/essentialkaos/translit/commit/d62c0f98f9b32cda180f3e875d80a6afbaf34d9b):

- neutral: only Russian -> Latin
- good: really nice README, really nice code
- good: plenty of transliteration standards implemented
- good: has support for previous+current rune analysis for special case converting
- good: nice tests (probably)
- minor: too much "branding"
- bad (IMO): custom license (and legal stuff is always scary, at least for me)

[dchest/translit.go](https://github.com/dchest/translit.go) at [5528f11](https://github.com/dchest/translit.go/commit/5528f1177236f74b86bf5eecb7381bcda1074cba):

- neutral: only Russian -> Latin
- fatal: doesn't even compile with Go 1.11
- bad: last updated in 2011
- minor: single-rune conversion rules only (like, no way for special multi-char cases, though that may be relevant only for Ukrainian and/or Bulgarian, not sure)
- nice: attempt to handle cases like `ЩИ` -> `SCHI`, `Щи` -> `Schi`, though don't think it really matters / cannot be worked around with case-sensitive converter

Summary: [essentialkaos/translit](https://github.com/essentialkaos/translit) is non-flexible, but extremely nice for Russian -> Latin transliteration (plenty of standards implemented, handles special multi-char cases), but has custom license.
