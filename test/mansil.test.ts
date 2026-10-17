import { test } from "node:test";
import assert from "node:assert/strict";

import Mansil from "../npm/index.ts";

const LINK = "\u001b]8;;https://example.com\u0007";
const LINK_END = "\u001b]8;;\u0007";
const RED = Mansil.RedFg;
const RESET = Mansil.Reset;

test("controls", () => {
    assert.equal(Mansil.clearScreenDown, "\x1b[J");
    assert.equal(Mansil.clearScreenUp, "\x1b[1J");
    assert.equal(Mansil.clearLineRight, "\x1b[K");
    assert.equal(Mansil.clearLineLeft, "\x1b[1K");
    assert.equal(Mansil.hideCursor, "\x1b[?25l");
    assert.equal(Mansil.showCursor, "\x1b[?25h");
});

test("cursor movement", () => {
    const moves: Record<string, (n: number) => string> = {
        A: Mansil.cursorUp,
        B: Mansil.cursorDown,
        C: Mansil.cursorRight,
        D: Mansil.cursorLeft,
        E: Mansil.cursorNextLine,
        F: Mansil.cursorPrevLine,
    };
    // The helpers are passed around unbound, so they must not rely on `this`
    for (const [final, move] of Object.entries(moves)) {
        assert.equal(move(3), "\x1b[3" + final);
        assert.equal(move(0), "");
        assert.equal(move(-2), "");
    }

    // Absolute positions are 1-based and are not changed for 0
    assert.equal(Mansil.cursorColumn(0), "\x1b[0G");
    assert.equal(Mansil.cursorPosition(2, 5), "\x1b[2;5H");
});

test("strip", () => {
    const cases: [string, string][] = [
        ["", ""],
        ["hello", "hello"],
        [RED + "he" + Mansil.Bold + "llo" + RESET, "hello"],
        ["a" + Mansil.hideCursor + "b" + Mansil.cursorPosition(3, 4) + "c", "abc"],
        [LINK + "link" + LINK_END, "link"],
        ["\x1b]8;;https://example.com\x1b\\link\x1b]8;;\x1b\\", "link"],
        ["a\x1b7b\x1b8c", "abc"],
        ["a\x1b(Bb", "ab"],
        ["abc\x1b", "abc"],
        ["abc\x1b[31", "abc"],
        ["日本" + Mansil.GreenFg + "語" + RESET, "日本語"],
    ];
    for (const [s, want] of cases) {
        assert.equal(Mansil.strip(s), want, JSON.stringify(s));
    }
});

test("visibleWidth", () => {
    const cases: [string, number][] = [
        ["", 0],
        ["hello", 5],
        [RED + "hello" + RESET, 5],
        [RED + Mansil.Bold + RESET + Mansil.cursorUp1 + Mansil.clearScreenDown, 0],
        [LINK + "link" + LINK_END, 4],
        ["⠹ 12 passed", 11],
        ["日本", 4],
        ["✅ ok", 5],
        ["🎉", 2],
        ["e\u0301", 1],
        ["a\u200bb", 2],
        ["a\tb\n", 2],
        ["⚠", 1],
        ["⚠\u{FE0F} ok", 5],
        ["✅\u{FE0F}", 2],
    ];
    for (const [s, want] of cases) {
        assert.equal(Mansil.visibleWidth(s), want, JSON.stringify(s));
    }
});

test("truncate", () => {
    const cases: [string, string, number, string, string][] = [
        ["fits", "hello", 5, "…", "hello"],
        ["styled and fits", RED + "hello" + RESET, 5, "…", RED + "hello" + RESET],
        ["plain cut", "hello world", 8, "…", "hello w…"],
        ["no tail", "hello world", 3, "", "hel"],
        ["code in the middle", "hello " + RED + "world" + RESET, 8, "…", "hello " + RED + "w…" + RESET],
        ["cut inside a styled span", RED + "hello world" + RESET, 6, "…", RED + "hello…" + RESET],
        ["tail as wide as the width", RED + "hello" + RESET, 3, "...", RED + "..." + RESET],
        ["tail wider than the width", "hello", 2, "...", "he"],
        ["width 0", "hello", 0, "…", ""],
        ["negative width", "hello", -1, "…", ""],
        ["only codes", RED + RESET, 3, "…", RED + RESET],
        ["wide characters", "日本語", 5, "…", "日本…"],
        ["wide character at the cut", "日本語", 4, "…", "日…"],
        ["emoji", "🎉🎉🎉", 5, "…", "🎉🎉…"],
        ["emoji selector", "⚠\u{FE0F}ok", 3, "…", "⚠\u{FE0F}…"],
        ["emoji selector at the cut", "⚠\u{FE0F}ok", 2, "…", "⚠…"],
        ["hyperlink", LINK + "click here" + LINK_END, 6, "…", LINK + "click…" + LINK_END],
    ];
    for (const [name, s, width, tail, want] of cases) {
        const got = Mansil.truncate(s, width, tail);
        assert.equal(got, want, name);
        if (width > 0) {
            assert.ok(Mansil.visibleWidth(got) <= width, `${name}: ${JSON.stringify(got)} is wider than ${width}`);
        }
    }

    assert.equal(Mansil.truncate("hello world", 5), "hello", "default tail");
});
