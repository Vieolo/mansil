import unittest

import mansil

LINK = "\033]8;;https://example.com\a"
LINK_END = "\033]8;;\a"
RED = mansil.RED_FG
RESET = mansil.RESET


class TestControls(unittest.TestCase):
    def test_codes(self):
        self.assertEqual(mansil.CLEAR_SCREEN_DOWN, "\x1b[J")
        self.assertEqual(mansil.CLEAR_SCREEN_UP, "\x1b[1J")
        self.assertEqual(mansil.CLEAR_LINE_RIGHT, "\x1b[K")
        self.assertEqual(mansil.CLEAR_LINE_LEFT, "\x1b[1K")
        self.assertEqual(mansil.HIDE_CURSOR, "\x1b[?25l")
        self.assertEqual(mansil.SHOW_CURSOR, "\x1b[?25h")


class TestCursorMovement(unittest.TestCase):
    def test_relative(self):
        moves = {
            "A": mansil.cursor_up,
            "B": mansil.cursor_down,
            "C": mansil.cursor_right,
            "D": mansil.cursor_left,
            "E": mansil.cursor_next_line,
            "F": mansil.cursor_prev_line,
        }
        for final, move in moves.items():
            with self.subTest(final=final):
                self.assertEqual(move(3), "\x1b[3" + final)
                self.assertEqual(move(0), "")
                self.assertEqual(move(-2), "")

    def test_absolute(self):
        # Absolute positions are 1-based and are not changed for 0
        self.assertEqual(mansil.cursor_column(0), "\x1b[0G")
        self.assertEqual(mansil.cursor_position(2, 5), "\x1b[2;5H")


class TestStrip(unittest.TestCase):
    def test_strip(self):
        cases = [
            ("", ""),
            ("hello", "hello"),
            (RED + "he" + mansil.BOLD + "llo" + RESET, "hello"),
            ("a" + mansil.HIDE_CURSOR + "b" + mansil.cursor_position(3, 4) + "c", "abc"),
            (LINK + "link" + LINK_END, "link"),
            ("\033]8;;https://example.com\033\\link\033]8;;\033\\", "link"),
            ("a\0337b\0338c", "abc"),
            ("a\033(Bb", "ab"),
            ("abc\033", "abc"),
            ("abc\033[31", "abc"),
            ("日本" + mansil.GREEN_FG + "語" + RESET, "日本語"),
        ]
        for s, want in cases:
            with self.subTest(s=s):
                self.assertEqual(mansil.strip(s), want)


class TestVisibleWidth(unittest.TestCase):
    def test_visible_width(self):
        cases = [
            ("", 0),
            ("hello", 5),
            (RED + "hello" + RESET, 5),
            (RED + mansil.BOLD + RESET + mansil.CURSOR_UP_1 + mansil.CLEAR_SCREEN_DOWN, 0),
            (LINK + "link" + LINK_END, 4),
            ("⠹ 12 passed", 11),
            ("日本", 4),
            ("✅ ok", 5),
            ("🎉", 2),
            ("e\u0301", 1),
            ("a\u200bb", 2),
            ("a\tb\n", 2),
            ("⚠", 1),
            ("⚠\U0000FE0F ok", 5),
            ("✅\U0000FE0F", 2),
        ]
        for s, want in cases:
            with self.subTest(s=s):
                self.assertEqual(mansil.visible_width(s), want)


class TestTruncate(unittest.TestCase):
    def test_truncate(self):
        cases = [
            ("fits", "hello", 5, "…", "hello"),
            ("styled and fits", RED + "hello" + RESET, 5, "…", RED + "hello" + RESET),
            ("plain cut", "hello world", 8, "…", "hello w…"),
            ("no tail", "hello world", 3, "", "hel"),
            ("code in the middle", "hello " + RED + "world" + RESET, 8, "…", "hello " + RED + "w…" + RESET),
            ("cut inside a styled span", RED + "hello world" + RESET, 6, "…", RED + "hello…" + RESET),
            ("tail as wide as the width", RED + "hello" + RESET, 3, "...", RED + "..." + RESET),
            ("tail wider than the width", "hello", 2, "...", "he"),
            ("width 0", "hello", 0, "…", ""),
            ("negative width", "hello", -1, "…", ""),
            ("only codes", RED + RESET, 3, "…", RED + RESET),
            ("wide characters", "日本語", 5, "…", "日本…"),
            ("wide character at the cut", "日本語", 4, "…", "日…"),
            ("emoji selector", "⚠\U0000FE0Fok", 3, "…", "⚠\U0000FE0F…"),
            ("emoji selector at the cut", "⚠\U0000FE0Fok", 2, "…", "⚠…"),
            ("emoji", "🎉🎉🎉", 5, "…", "🎉🎉…"),
            ("hyperlink", LINK + "click here" + LINK_END, 6, "…", LINK + "click…" + LINK_END),
        ]
        for name, s, width, tail, want in cases:
            with self.subTest(name):
                got = mansil.truncate(s, width, tail)
                self.assertEqual(got, want)
                if width > 0:
                    self.assertLessEqual(mansil.visible_width(got), width)

    def test_default_tail(self):
        self.assertEqual(mansil.truncate("hello world", 5), "hello")


if __name__ == "__main__":
    unittest.main()
