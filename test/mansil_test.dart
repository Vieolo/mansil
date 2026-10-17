import 'package:mansil/mansil.dart';
import 'package:test/test.dart';

const link = "\u001b]8;;https://example.com\u0007";
const linkEnd = "\u001b]8;;\u0007";
const red = Mansil.redFg;
const reset = Mansil.reset;

void main() {
  test("controls", () {
    expect(Mansil.clearScreenDown, "\x1b[J");
    expect(Mansil.clearScreenUp, "\x1b[1J");
    expect(Mansil.clearLineRight, "\x1b[K");
    expect(Mansil.clearLineLeft, "\x1b[1K");
    expect(Mansil.hideCursor, "\x1b[?25l");
    expect(Mansil.showCursor, "\x1b[?25h");
  });

  test("cursor movement", () {
    final moves = <String, String Function(int)>{
      "A": Mansil.cursorUp,
      "B": Mansil.cursorDown,
      "C": Mansil.cursorRight,
      "D": Mansil.cursorLeft,
      "E": Mansil.cursorNextLine,
      "F": Mansil.cursorPrevLine,
    };
    moves.forEach((last, move) {
      expect(move(3), "\x1b[3$last");
      expect(move(0), "");
      expect(move(-2), "");
    });

    // Absolute positions are 1-based and are not changed for 0
    expect(Mansil.cursorColumn(0), "\x1b[0G");
    expect(Mansil.cursorPosition(2, 5), "\x1b[2;5H");
  });

  test("strip", () {
    final cases = <(String, String)>[
      ("", ""),
      ("hello", "hello"),
      ("${red}he${Mansil.bold}llo$reset", "hello"),
      ("a${Mansil.hideCursor}b${Mansil.cursorPosition(3, 4)}c", "abc"),
      ("${link}link$linkEnd", "link"),
      ("\x1b]8;;https://example.com\x1b\\link\x1b]8;;\x1b\\", "link"),
      ("a\x1b7b\x1b8c", "abc"),
      ("a\x1b(Bb", "ab"),
      ("abc\x1b", "abc"),
      ("abc\x1b[31", "abc"),
      ("日本${Mansil.greenFg}語$reset", "日本語"),
    ];
    for (final (s, want) in cases) {
      expect(Mansil.strip(s), want, reason: s);
    }
  });

  test("visibleWidth", () {
    final cases = <(String, int)>[
      ("", 0),
      ("hello", 5),
      ("${red}hello$reset", 5),
      (
        "$red${Mansil.bold}$reset${Mansil.cursorUp1}${Mansil.clearScreenDown}",
        0,
      ),
      ("${link}link$linkEnd", 4),
      ("⠹ 12 passed", 11),
      ("日本", 4),
      ("✅ ok", 5),
      ("🎉", 2),
      ("e\u0301", 1),
      ("a\u200bb", 2),
      ("a\tb\n", 2),
      ("⚠", 1),
      ("⚠\u{FE0F} ok", 5),
      ("✅\u{FE0F}", 2),
    ];
    for (final (s, want) in cases) {
      expect(Mansil.visibleWidth(s), want, reason: s);
    }
  });

  test("truncate", () {
    final cases = <(String, String, int, String, String)>[
      ("fits", "hello", 5, "…", "hello"),
      ("styled and fits", "${red}hello$reset", 5, "…", "${red}hello$reset"),
      ("plain cut", "hello world", 8, "…", "hello w…"),
      ("no tail", "hello world", 3, "", "hel"),
      (
        "code in the middle",
        "hello ${red}world$reset",
        8,
        "…",
        "hello ${red}w…$reset",
      ),
      (
        "cut inside a styled span",
        "${red}hello world$reset",
        6,
        "…",
        "${red}hello…$reset",
      ),
      (
        "tail as wide as the width",
        "${red}hello$reset",
        3,
        "...",
        "$red...$reset",
      ),
      ("tail wider than the width", "hello", 2, "...", "he"),
      ("width 0", "hello", 0, "…", ""),
      ("emoji selector", "⚠\u{FE0F}ok", 3, "…", "⚠\u{FE0F}…"),
      ("emoji selector at the cut", "⚠\u{FE0F}ok", 2, "…", "⚠…"),
      ("negative width", "hello", -1, "…", ""),
      ("only codes", "$red$reset", 3, "…", "$red$reset"),
      ("wide characters", "日本語", 5, "…", "日本…"),
      ("wide character at the cut", "日本語", 4, "…", "日…"),
      ("emoji", "🎉🎉🎉", 5, "…", "🎉🎉…"),
      (
        "hyperlink",
        "${link}click here$linkEnd",
        6,
        "…",
        "${link}click…$linkEnd",
      ),
    ];
    for (final (name, s, width, tail, want) in cases) {
      final got = Mansil.truncate(s, width, tail);
      expect(got, want, reason: name);
      if (width > 0) {
        expect(
          Mansil.visibleWidth(got),
          lessThanOrEqualTo(width),
          reason: name,
        );
      }
    }

    expect(Mansil.truncate("hello world", 5), "hello", reason: "default tail");
  });
}
