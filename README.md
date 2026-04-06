# ascii-art-fs

A command-line tool written in Go that converts text into ASCII art using a variety of banner fonts.

## Usage

```bash
go run . "your text here"
```

By default, the program uses the `standard` font. You can specify a different font as a second argument:

```bash
go run . "your text here" shadow
go run . "your text here" thinkertoy
```

### Newlines

To render text on multiple lines, use `\n` in your input:

```bash
go run . "Hello\nWorld"
```

## Available fonts

- `standard`
- `shadow`
- `thinkertoy`

## How it works

Each font is stored as a `.txt` file where every printable ASCII character (from space to `~`) is represented as 8 rows of ASCII art. Characters are separated by a blank line, making each character block 9 lines tall.

When given an input string, the program looks up each character in the banner file using its ASCII value to calculate the correct line index, then renders the full string row by row.

## Project structure

```
.
├── main.go
└── internal/
    └── functions/
        └── asciiConv.go
```

## Running tests

```bash
go test ./...
```

## Author

[muhsar27](https://github.com/muhsar27)
