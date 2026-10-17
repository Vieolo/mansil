//! Mansil ANSI Codes

// GEN START

// Styles
pub const RESET: &str = "\x1b[0m";
pub const BOLD: &str = "\x1b[1m";
pub const DIM: &str = "\x1b[2m";
pub const ITALIC: &str = "\x1b[3m";
pub const UNDERLINE: &str = "\x1b[4m";
pub const BLINK: &str = "\x1b[5m";
pub const INVERSE: &str = "\x1b[7m";
pub const HIDDEN: &str = "\x1b[8m";
pub const STRIKETHROUGH: &str = "\x1b[9m";

// Colors
pub const BLACK_FG: &str = "\x1b[30m";
pub const BLACK_BG: &str = "\x1b[40m";
pub const BLACK_FG_BRIGHT: &str = "\x1b[90m";
pub const BLACK_BG_BRIGHT: &str = "\x1b[100m";
pub const RED_FG: &str = "\x1b[31m";
pub const RED_BG: &str = "\x1b[41m";
pub const RED_FG_BRIGHT: &str = "\x1b[91m";
pub const RED_BG_BRIGHT: &str = "\x1b[101m";
pub const GREEN_FG: &str = "\x1b[32m";
pub const GREEN_BG: &str = "\x1b[42m";
pub const GREEN_FG_BRIGHT: &str = "\x1b[92m";
pub const GREEN_BG_BRIGHT: &str = "\x1b[102m";
pub const YELLOW_FG: &str = "\x1b[33m";
pub const YELLOW_BG: &str = "\x1b[43m";
pub const YELLOW_FG_BRIGHT: &str = "\x1b[93m";
pub const YELLOW_BG_BRIGHT: &str = "\x1b[103m";
pub const BLUE_FG: &str = "\x1b[34m";
pub const BLUE_BG: &str = "\x1b[44m";
pub const BLUE_FG_BRIGHT: &str = "\x1b[94m";
pub const BLUE_BG_BRIGHT: &str = "\x1b[104m";
pub const MAGENTA_FG: &str = "\x1b[35m";
pub const MAGENTA_BG: &str = "\x1b[45m";
pub const MAGENTA_FG_BRIGHT: &str = "\x1b[95m";
pub const MAGENTA_BG_BRIGHT: &str = "\x1b[105m";
pub const CYAN_FG: &str = "\x1b[36m";
pub const CYAN_BG: &str = "\x1b[46m";
pub const CYAN_FG_BRIGHT: &str = "\x1b[96m";
pub const CYAN_BG_BRIGHT: &str = "\x1b[106m";
pub const WHITE_FG: &str = "\x1b[37m";
pub const WHITE_BG: &str = "\x1b[47m";
pub const WHITE_FG_BRIGHT: &str = "\x1b[97m";
pub const WHITE_BG_BRIGHT: &str = "\x1b[107m";

// Controls
pub const CLEAR_SCREEN: &str = "\x1b[2J";
pub const CLEAR_SCREEN_DOWN: &str = "\x1b[J";
pub const CLEAR_SCREEN_UP: &str = "\x1b[1J";
pub const CLEAR_LINE: &str = "\x1b[2K";
pub const CLEAR_LINE_RIGHT: &str = "\x1b[K";
pub const CLEAR_LINE_LEFT: &str = "\x1b[1K";
pub const HIDE_CURSOR: &str = "\x1b[?25l";
pub const SHOW_CURSOR: &str = "\x1b[?25h";
pub const CURSOR_UP_1: &str = "\x1b[1A";
pub const CURSOR_DOWN_1: &str = "\x1b[1B";
pub const CURSOR_RIGHT_1: &str = "\x1b[1C";
pub const CURSOR_LEFT_1: &str = "\x1b[1D";
pub const CURSOR_NEXT_LINE_1: &str = "\x1b[1E";
pub const CURSOR_PREV_LINE_1: &str = "\x1b[1F";
pub const CURSOR_COLUMN_1: &str = "\x1b[1G";

// Character widths, used by visible_width and truncate
const ZERO_WIDTH_RANGES: &[(u32, u32)] = &[
    (0x0000, 0x001F),   // C0 controls
    (0x007F, 0x009F),   // Delete and C1 controls
    (0x0300, 0x036F),   // Combining Diacritical Marks
    (0x1AB0, 0x1AFF),   // Combining Diacritical Marks Extended
    (0x1DC0, 0x1DFF),   // Combining Diacritical Marks Supplement
    (0x200B, 0x200F),   // Zero width space, joiners and direction marks
    (0x2060, 0x2064),   // Word joiner and invisible operators
    (0x20D0, 0x20FF),   // Combining Diacritical Marks for Symbols
    (0xFE00, 0xFE0F),   // Variation Selectors
    (0xFE20, 0xFE2F),   // Combining Half Marks
    (0xFEFF, 0xFEFF),   // Zero width no-break space
    (0xE0000, 0xE007F), // Tags
    (0xE0100, 0xE01EF), // Variation Selectors Supplement
];
const WIDE_RANGES: &[(u32, u32)] = &[
    (0x1100, 0x115F),   // Hangul Jamo initial consonants
    (0x231A, 0x231B),   // Watch and hourglass
    (0x2329, 0x232A),   // Angle brackets
    (0x23E9, 0x23EC),   // Fast-forward and rewind buttons
    (0x23F0, 0x23F0),   // Alarm clock
    (0x23F3, 0x23F3),   // Hourglass with flowing sand
    (0x25FD, 0x25FE),   // Medium small squares
    (0x2614, 0x2615),   // Umbrella with rain and hot beverage
    (0x2648, 0x2653),   // Zodiac signs
    (0x267F, 0x267F),   // Wheelchair symbol
    (0x2693, 0x2693),   // Anchor
    (0x26A1, 0x26A1),   // High voltage
    (0x26AA, 0x26AB),   // Medium circles
    (0x26BD, 0x26BE),   // Soccer ball and baseball
    (0x26C4, 0x26C5),   // Snowman and sun behind cloud
    (0x26CE, 0x26CE),   // Ophiuchus
    (0x26D4, 0x26D4),   // No entry
    (0x26EA, 0x26EA),   // Church
    (0x26F2, 0x26F3),   // Fountain and flag in hole
    (0x26F5, 0x26F5),   // Sailboat
    (0x26FA, 0x26FA),   // Tent
    (0x26FD, 0x26FD),   // Fuel pump
    (0x2705, 0x2705),   // Check mark button
    (0x270A, 0x270B),   // Raised fist and hand
    (0x2728, 0x2728),   // Sparkles
    (0x274C, 0x274C),   // Cross mark
    (0x274E, 0x274E),   // Cross mark button
    (0x2753, 0x2755),   // Question and exclamation marks
    (0x2757, 0x2757),   // Exclamation mark
    (0x2795, 0x2797),   // Plus, minus and divide signs
    (0x27B0, 0x27B0),   // Curly loop
    (0x27BF, 0x27BF),   // Double curly loop
    (0x2B1B, 0x2B1C),   // Large squares
    (0x2B50, 0x2B50),   // Star
    (0x2B55, 0x2B55),   // Hollow red circle
    (0x2E80, 0x303E),   // CJK radicals, symbols and punctuation
    (0x3040, 0xA4CF),   // Kana, Bopomofo, Hangul compatibility jamo, CJK and Yi
    (0xA960, 0xA97F),   // Hangul Jamo Extended-A
    (0xAC00, 0xD7A3),   // Hangul Syllables
    (0xF900, 0xFAFF),   // CJK Compatibility Ideographs
    (0xFE10, 0xFE19),   // Vertical Forms
    (0xFE30, 0xFE6F),   // CJK Compatibility Forms and Small Form Variants
    (0xFF00, 0xFF60),   // Fullwidth Forms
    (0xFFE0, 0xFFE6),   // Fullwidth signs
    (0x16FE0, 0x18DFF), // Ideographic symbols, Tangut and Khitan
    (0x1AFF0, 0x1B2FF), // Kana supplements and Nushu
    (0x1F004, 0x1F004), // Mahjong tile red dragon
    (0x1F0CF, 0x1F0CF), // Joker
    (0x1F18E, 0x1F18E), // AB button
    (0x1F191, 0x1F19A), // Squared words
    (0x1F200, 0x1F2FF), // Enclosed Ideographic Supplement
    (0x1F300, 0x1F64F), // Miscellaneous Symbols and Pictographs, Emoticons
    (0x1F680, 0x1F6FF), // Transport and Map Symbols
    (0x1F7E0, 0x1F7FF), // Large circles and squares
    (0x1F900, 0x1F9FF), // Supplemental Symbols and Pictographs
    (0x1FA70, 0x1FAFF), // Symbols and Pictographs Extended-A
    (0x20000, 0x2FFFD), // Supplementary Ideographic Plane
    (0x30000, 0x3FFFD), // Tertiary Ideographic Plane
];

// GEN END

// Cursor movement

/// Moves the cursor up by `n` lines. Returns an empty string when `n` is 0
pub fn cursor_up(n: u32) -> String {
    move_by(CURSOR_UP_1, n)
}

/// Moves the cursor down by `n` lines. Returns an empty string when `n` is 0
pub fn cursor_down(n: u32) -> String {
    move_by(CURSOR_DOWN_1, n)
}

/// Moves the cursor right by `n` columns. Returns an empty string when `n` is 0
pub fn cursor_right(n: u32) -> String {
    move_by(CURSOR_RIGHT_1, n)
}

/// Moves the cursor left by `n` columns. Returns an empty string when `n` is 0
pub fn cursor_left(n: u32) -> String {
    move_by(CURSOR_LEFT_1, n)
}

/// Moves the cursor to the start of the line `n` lines down. Returns an empty string when `n` is 0
pub fn cursor_next_line(n: u32) -> String {
    move_by(CURSOR_NEXT_LINE_1, n)
}

/// Moves the cursor to the start of the line `n` lines up. Returns an empty string when `n` is 0
pub fn cursor_prev_line(n: u32) -> String {
    move_by(CURSOR_PREV_LINE_1, n)
}

/// Moves the cursor to column `n` of the current line. Columns are 1-based
pub fn cursor_column(n: u32) -> String {
    CURSOR_COLUMN_1.replace("[1", &format!("[{}", n))
}

/// Moves the cursor to the given row and column. Both are 1-based
pub fn cursor_position(row: u32, col: u32) -> String {
    format!("\x1b[{};{}H", row, col)
}

// Sets the count of a relative movement, whose constant moves by 1.
// In CSI sequences a count of 0 means the default of 1, so 0 returns an empty string instead
fn move_by(code_1: &str, n: u32) -> String {
    if n == 0 {
        return String::new();
    }
    code_1.replace("[1", &format!("[{}", n))
}

// Styled text

const ESC: u8 = 0x1b;

/// Returns `s` without its escape sequences.
///
/// It removes CSI sequences (styles, colors, cursor movement), OSC sequences (such as hyperlinks)
/// and the other two-character escape sequences
pub fn strip(s: &str) -> String {
    let bytes = s.as_bytes();
    let mut out = String::with_capacity(s.len());
    let mut i = 0;
    while i < bytes.len() {
        if bytes[i] == ESC {
            i += escape_len(bytes, i);
            continue;
        }
        let ch = char_at(s, i);
        out.push(ch);
        i += ch.len_utf8();
    }
    out
}

/// Returns the number of terminal columns that `s` takes, ignoring its escape sequences.
///
/// Wide characters, such as CJK and most emoji, take two columns, and so does a narrow character
/// followed by the emoji presentation selector (U+FE0F). Combining marks, zero-width characters
/// and control characters, including tabs and newlines, take none. The characters of a joined
/// emoji (ZWJ sequences, skin tones) are counted one by one, so it can come out wider than the
/// terminal draws it
pub fn visible_width(s: &str) -> usize {
    let bytes = s.as_bytes();
    let (mut width, mut prev) = (0, 0);
    let mut i = 0;
    while i < bytes.len() {
        if bytes[i] == ESC {
            i += escape_len(bytes, i);
            continue;
        }
        let ch = char_at(s, i);
        prev = char_width(ch, prev);
        width += prev;
        i += ch.len_utf8();
    }
    width
}

/// Cuts `s` to at most `width` terminal columns, ending it with `tail` when it is cut.
///
/// The tail counts towards the width, and is left out when it is wider than `width` on its own.
/// Only the visible characters after the cut are dropped. Every escape sequence is kept, so the
/// styles and hyperlinks that `s` closes are still closed. `s` comes back unchanged when it fits,
/// and a `width` of 0 returns an empty string
pub fn truncate(s: &str, width: usize, tail: &str) -> String {
    if width == 0 {
        return String::new();
    }
    if visible_width(s) <= width {
        return s.to_string();
    }
    let mut tail = tail;
    let mut tail_width = visible_width(tail);
    if tail_width > width {
        tail = "";
        tail_width = 0;
    }
    let room = width - tail_width;

    let bytes = s.as_bytes();
    let mut out = String::with_capacity(s.len());
    let (mut used, mut prev, mut cut) = (0, 0, false);
    let mut i = 0;
    while i < bytes.len() {
        if bytes[i] == ESC {
            let n = escape_len(bytes, i);
            out.push_str(&s[i..i + n]);
            i += n;
            continue;
        }
        let ch = char_at(s, i);
        if !cut {
            let w = char_width(ch, prev);
            if used + w <= room {
                out.push(ch);
                used += w;
                prev = w;
            } else {
                out.push_str(tail);
                cut = true;
            }
        }
        i += ch.len_utf8();
    }
    out
}

// Returns the length of the escape sequence that starts at s[i], which is ESC.
// - CSI: ESC [, then parameter and intermediate bytes, then a final byte
// - OSC: ESC ], then anything up to BEL or ESC \
// - Others: ESC, then intermediate bytes, then a final byte
// A malformed sequence ends at the first byte that cannot belong to it, an unterminated OSC
// runs to the end of s, and a lone ESC has a length of 1.
// Every byte that a sequence ends on is ASCII, so the length always ends on a char boundary
fn escape_len(s: &[u8], i: usize) -> usize {
    let mut j = i + 1;
    if j == s.len() {
        return 1;
    }
    match s[j] {
        b'[' => {
            j += 1;
            while j < s.len() && (0x20..=0x3F).contains(&s[j]) {
                j += 1;
            }
            if j < s.len() && (0x40..=0x7E).contains(&s[j]) {
                j += 1;
            }
        }
        b']' => {
            j += 1;
            while j < s.len() {
                if s[j] == 0x07 {
                    return j + 1 - i;
                }
                if s[j] == ESC {
                    return if s.get(j + 1) == Some(&b'\\') {
                        j + 2 - i
                    } else {
                        j - i
                    };
                }
                j += 1;
            }
        }
        _ => {
            while j < s.len() && (0x20..=0x2F).contains(&s[j]) {
                j += 1;
            }
            if j < s.len() && (0x30..=0x7E).contains(&s[j]) {
                j += 1;
            }
        }
    }
    j - i
}

// Returns the char that starts at byte i of s
fn char_at(s: &str, i: usize) -> char {
    s[i..].chars().next().unwrap()
}

// Returns the number of terminal columns that ch takes, after a character that takes prev columns.
// The emoji presentation selector (U+FE0F) widens a narrow character, such as ⚠, to two columns
fn char_width(ch: char, prev: usize) -> usize {
    let cp = ch as u32;
    if (0x20..0x7F).contains(&cp) {
        return 1;
    }
    if cp == 0xFE0F && prev == 1 {
        return 1;
    }
    if in_ranges(cp, ZERO_WIDTH_RANGES) {
        return 0;
    }
    if in_ranges(cp, WIDE_RANGES) {
        return 2;
    }
    1
}

fn in_ranges(cp: u32, ranges: &[(u32, u32)]) -> bool {
    ranges
        .iter()
        .any(|&(start, end)| (start..=end).contains(&cp))
}

#[cfg(test)]
mod tests {
    use super::*;

    const LINK: &str = "\x1b]8;;https://example.com\x07";
    const LINK_END: &str = "\x1b]8;;\x07";

    #[test]
    fn controls() {
        assert_eq!(CLEAR_SCREEN_DOWN, "\x1b[J");
        assert_eq!(CLEAR_SCREEN_UP, "\x1b[1J");
        assert_eq!(CLEAR_LINE_RIGHT, "\x1b[K");
        assert_eq!(CLEAR_LINE_LEFT, "\x1b[1K");
        assert_eq!(HIDE_CURSOR, "\x1b[?25l");
        assert_eq!(SHOW_CURSOR, "\x1b[?25h");
    }

    #[test]
    fn cursor_movement() {
        let moves: [(&str, fn(u32) -> String); 6] = [
            ("A", cursor_up),
            ("B", cursor_down),
            ("C", cursor_right),
            ("D", cursor_left),
            ("E", cursor_next_line),
            ("F", cursor_prev_line),
        ];
        for (letter, move_cursor) in moves {
            assert_eq!(move_cursor(3), format!("\x1b[3{}", letter));
            assert_eq!(move_cursor(0), "");
        }

        // Absolute positions are 1-based and are not changed for 0
        assert_eq!(cursor_column(0), "\x1b[0G");
        assert_eq!(cursor_position(2, 5), "\x1b[2;5H");
    }

    #[test]
    fn strip_codes() {
        let cases = [
            (String::new(), ""),
            ("hello".to_string(), "hello"),
            (format!("{RED_FG}he{BOLD}llo{RESET}"), "hello"),
            (format!("a{HIDE_CURSOR}b{}c", cursor_position(3, 4)), "abc"),
            (format!("{LINK}link{LINK_END}"), "link"),
            (
                "\x1b]8;;https://example.com\x1b\\link\x1b]8;;\x1b\\".to_string(),
                "link",
            ),
            ("a\x1b7b\x1b8c".to_string(), "abc"),
            ("a\x1b(Bb".to_string(), "ab"),
            ("abc\x1b".to_string(), "abc"),
            ("abc\x1b[31".to_string(), "abc"),
            (format!("日本{GREEN_FG}語{RESET}"), "日本語"),
        ];
        for (s, want) in cases {
            assert_eq!(strip(&s), want, "strip({:?})", s);
        }
    }

    #[test]
    fn visible_widths() {
        let cases = [
            (String::new(), 0),
            ("hello".to_string(), 5),
            (format!("{RED_FG}hello{RESET}"), 5),
            (
                format!("{RED_FG}{BOLD}{RESET}{CURSOR_UP_1}{CLEAR_SCREEN_DOWN}"),
                0,
            ),
            (format!("{LINK}link{LINK_END}"), 4),
            ("⠹ 12 passed".to_string(), 11),
            ("日本".to_string(), 4),
            ("✅ ok".to_string(), 5),
            ("🎉".to_string(), 2),
            ("e\u{0301}".to_string(), 1),
            ("a\u{200b}b".to_string(), 2),
            ("a\tb\n".to_string(), 2),
            ("⚠".to_string(), 1),
            ("⚠\u{FE0F} ok".to_string(), 5),
            ("✅\u{FE0F}".to_string(), 2),
        ];
        for (s, want) in cases {
            assert_eq!(visible_width(&s), want, "visible_width({:?})", s);
        }
    }

    #[test]
    fn truncation() {
        let styled = |s: &str| format!("{RED_FG}{s}{RESET}");
        let cases = [
            ("fits", "hello".to_string(), 5, "…", "hello".to_string()),
            ("styled and fits", styled("hello"), 5, "…", styled("hello")),
            (
                "plain cut",
                "hello world".to_string(),
                8,
                "…",
                "hello w…".to_string(),
            ),
            (
                "no tail",
                "hello world".to_string(),
                3,
                "",
                "hel".to_string(),
            ),
            (
                "code in the middle",
                format!("hello {}", styled("world")),
                8,
                "…",
                format!("hello {}", styled("w…")),
            ),
            (
                "cut inside a styled span",
                styled("hello world"),
                6,
                "…",
                styled("hello…"),
            ),
            (
                "tail as wide as the width",
                styled("hello"),
                3,
                "...",
                styled("..."),
            ),
            (
                "tail wider than the width",
                "hello".to_string(),
                2,
                "...",
                "he".to_string(),
            ),
            ("width 0", "hello".to_string(), 0, "…", String::new()),
            (
                "emoji selector",
                "⚠\u{FE0F}ok".to_string(),
                3,
                "…",
                "⚠\u{FE0F}…".to_string(),
            ),
            (
                "emoji selector at the cut",
                "⚠\u{FE0F}ok".to_string(),
                2,
                "…",
                "⚠…".to_string(),
            ),
            ("only codes", styled(""), 3, "…", styled("")),
            (
                "wide characters",
                "日本語".to_string(),
                5,
                "…",
                "日本…".to_string(),
            ),
            (
                "wide character at the cut",
                "日本語".to_string(),
                4,
                "…",
                "日…".to_string(),
            ),
            ("emoji", "🎉🎉🎉".to_string(), 5, "…", "🎉🎉…".to_string()),
            (
                "hyperlink",
                format!("{LINK}click here{LINK_END}"),
                6,
                "…",
                format!("{LINK}click…{LINK_END}"),
            ),
        ];
        for (name, s, width, tail, want) in cases {
            let got = truncate(&s, width, tail);
            assert_eq!(got, want, "{}", name);
            assert!(
                visible_width(&got) <= width,
                "{}: {:?} is wider than {}",
                name,
                got,
                width
            );
        }
    }
}
