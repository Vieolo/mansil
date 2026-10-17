package mansil_test

import (
	"testing"

	"github.com/vieolo/mansil"
)

const link = "\033]8;;https://example.com\a"
const linkEnd = "\033]8;;\a"

func TestControls(t *testing.T) {
	cases := map[string]string{
		mansil.ClearScreenDown: "\x1b[J",
		mansil.ClearScreenUp:   "\x1b[1J",
		mansil.ClearLineRight:  "\x1b[K",
		mansil.ClearLineLeft:   "\x1b[1K",
		mansil.HideCursor:      "\x1b[?25l",
		mansil.ShowCursor:      "\x1b[?25h",
	}
	for got, want := range cases {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestCursorMovement(t *testing.T) {
	moves := map[string]func(int) string{
		"A": mansil.CursorUp,
		"B": mansil.CursorDown,
		"C": mansil.CursorRight,
		"D": mansil.CursorLeft,
		"E": mansil.CursorNextLine,
		"F": mansil.CursorPrevLine,
	}
	for final, move := range moves {
		if got := move(3); got != "\x1b[3"+final {
			t.Errorf("move %s by 3: got %q", final, got)
		}
		if got := move(0); got != "" {
			t.Errorf("move %s by 0: got %q, want an empty string", final, got)
		}
		if got := move(-2); got != "" {
			t.Errorf("move %s by -2: got %q, want an empty string", final, got)
		}
	}

	// Absolute positions are 1-based and are not changed for 0
	if got := mansil.CursorColumn(0); got != "\x1b[0G" {
		t.Errorf("CursorColumn(0): got %q", got)
	}
	if got := mansil.CursorPosition(2, 5); got != "\x1b[2;5H" {
		t.Errorf("CursorPosition(2, 5): got %q", got)
	}
}

func TestStrip(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"hello", "hello"},
		{mansil.RedFG + "he" + mansil.Bold + "llo" + mansil.Reset, "hello"},
		{"a" + mansil.HideCursor + "b" + mansil.CursorPosition(3, 4) + "c", "abc"},
		{link + "link" + linkEnd, "link"},
		{"\033]8;;https://example.com\033\\link\033]8;;\033\\", "link"},
		{"a\0337b\0338c", "abc"},
		{"a\033(Bb", "ab"},
		{"abc\033", "abc"},
		{"abc\033[31", "abc"},
		{"日本" + mansil.GreenFG + "語" + mansil.Reset, "日本語"},
	}
	for _, c := range cases {
		if got := mansil.Strip(c.in); got != c.want {
			t.Errorf("Strip(%q): got %q, want %q", c.in, got, c.want)
		}
	}
}

func TestVisibleWidth(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"hello", 5},
		{mansil.RedFG + "hello" + mansil.Reset, 5},
		{mansil.RedFG + mansil.Bold + mansil.Reset + mansil.CursorUp1 + mansil.ClearScreenDown, 0},
		{link + "link" + linkEnd, 4},
		{"⠹ 12 passed", 11},
		{"日本", 4},
		{"✅ ok", 5},
		{"🎉", 2},
		{"e\u0301", 1},
		{"a\u200bb", 2},
		{"a\tb\n", 2},
		{"⚠", 1},
		{"⚠\U0000FE0F ok", 5},
		{"✅\U0000FE0F", 2},
	}
	for _, c := range cases {
		if got := mansil.VisibleWidth(c.in); got != c.want {
			t.Errorf("VisibleWidth(%q): got %d, want %d", c.in, got, c.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	red, reset := mansil.RedFG, mansil.Reset
	cases := []struct {
		name  string
		in    string
		width int
		tail  string
		want  string
	}{
		{"fits", "hello", 5, "…", "hello"},
		{"styled and fits", red + "hello" + reset, 5, "…", red + "hello" + reset},
		{"plain cut", "hello world", 8, "…", "hello w…"},
		{"no tail", "hello world", 3, "", "hel"},
		{"code in the middle", "hello " + red + "world" + reset, 8, "…", "hello " + red + "w…" + reset},
		{"cut inside a styled span", red + "hello world" + reset, 6, "…", red + "hello…" + reset},
		{"tail as wide as the width", red + "hello" + reset, 3, "...", red + "..." + reset},
		{"tail wider than the width", "hello", 2, "...", "he"},
		{"width 0", "hello", 0, "…", ""},
		{"negative width", "hello", -1, "…", ""},
		{"only codes", red + reset, 3, "…", red + reset},
		{"wide characters", "日本語", 5, "…", "日本…"},
		{"wide character at the cut", "日本語", 4, "…", "日…"},
		{"emoji", "🎉🎉🎉", 5, "…", "🎉🎉…"},
		{"emoji selector", "⚠\U0000FE0Fok", 3, "…", "⚠\U0000FE0F…"},
		{"emoji selector at the cut", "⚠\U0000FE0Fok", 2, "…", "⚠…"},
		{"hyperlink", link + "click here" + linkEnd, 6, "…", link + "click…" + linkEnd},
	}
	for _, c := range cases {
		got := mansil.Truncate(c.in, c.width, c.tail)
		if got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
		if c.width > 0 && mansil.VisibleWidth(got) > c.width {
			t.Errorf("%s: %q is wider than %d", c.name, got, c.width)
		}
	}
}
