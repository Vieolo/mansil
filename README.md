# mansil
Mansil is a ***M***ulti-language ***ANSI*** escape code ***L***ibrary

It provides the commonly used ANSI escape codes and helper functions for:

- Dart
- Go
- JS/TS
- Python
- Rust


## Install

```bash
# Dart
dart pub add mansil

# Go
go get github.com/vieolo/mansil

# JS/TS
npm install @vieolo/mansil

# Python
pip install mansil # using pip
uv add mansil # using uv

# Rust
cargo add mansil
```

## Helpers

Besides the codes, every language has the same helpers, named in its own style:

| Helper | Go | Python, Rust | JS/TS, Dart |
|---|---|---|---|
| Move the cursor | `CursorUp(n)`, ... | `cursor_up(n)`, ... | `Mansil.cursorUp(n)`, ... |
| Remove escape sequences | `Strip(s)` | `strip(s)` | `Mansil.strip(s)` |
| Measure in terminal columns | `VisibleWidth(s)` | `visible_width(s)` | `Mansil.visibleWidth(s)` |
| Cut to a width | `Truncate(s, width, tail)` | `truncate(s, width, tail)` | `Mansil.truncate(s, width, tail)` |

- The relative movements (up, down, right, left, next line and previous line) return an empty string for `n <= 0`, so a distance that comes out as 0 does not move the cursor. `CursorColumn` and `CursorPosition` are absolute and 1-based.
- `VisibleWidth` counts CJK and most emoji as two columns, and so is a narrow character followed by the emoji presentation selector (U+FE0F). Combining marks, zero-width and control characters count as none. The characters of a joined emoji are counted one by one, so it can come out wider than the terminal draws it.
- `Truncate` drops only the visible characters after the cut and keeps every escape sequence, so the styles and hyperlinks that the text closes stay closed. The tail counts towards the width.

For example, to redraw a block of lines in place in Go:

```go
fmt.Print(mansil.CursorPrevLine(lastHeight) + mansil.ClearScreenDown)
for _, line := range lines {
	fmt.Println(mansil.Truncate(line, termWidth, "…"))
}
lastHeight = len(lines)
```

## Development

The codes and the character width tables live in `ansi_codes.json`. The generator writes them between the `GEN START` and `GEN END` comments of each language's source. Everything outside those comments is written by hand.

```bash
go run ./cmd/main.go generate      # regenerate the codes
uv run --no-project agent.py       # regenerate, then check and test every language
uv run --no-project agent.py go    # only some groups: generate, go, python, rust, ts, dart
```

The tests are in `mansil_test.go` (Go), `src/lib.rs` (Rust) and `test/` (Dart, Python and JS/TS).
