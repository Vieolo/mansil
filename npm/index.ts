export default class Mansil {
    // GEN START

    // Styles
    static readonly Reset = "\u001b[0m";
    static readonly Bold = "\u001b[1m";
    static readonly Dim = "\u001b[2m";
    static readonly Italic = "\u001b[3m";
    static readonly Underline = "\u001b[4m";
    static readonly Blink = "\u001b[5m";
    static readonly Inverse = "\u001b[7m";
    static readonly Hidden = "\u001b[8m";
    static readonly Strikethrough = "\u001b[9m";

    // Colors
    static readonly BlackFg = "\u001b[30m";
    static readonly BlackBg = "\u001b[40m";
    static readonly BlackFgBright = "\u001b[90m";
    static readonly BlackBgBright = "\u001b[100m";
    static readonly RedFg = "\u001b[31m";
    static readonly RedBg = "\u001b[41m";
    static readonly RedFgBright = "\u001b[91m";
    static readonly RedBgBright = "\u001b[101m";
    static readonly GreenFg = "\u001b[32m";
    static readonly GreenBg = "\u001b[42m";
    static readonly GreenFgBright = "\u001b[92m";
    static readonly GreenBgBright = "\u001b[102m";
    static readonly YellowFg = "\u001b[33m";
    static readonly YellowBg = "\u001b[43m";
    static readonly YellowFgBright = "\u001b[93m";
    static readonly YellowBgBright = "\u001b[103m";
    static readonly BlueFg = "\u001b[34m";
    static readonly BlueBg = "\u001b[44m";
    static readonly BlueFgBright = "\u001b[94m";
    static readonly BlueBgBright = "\u001b[104m";
    static readonly MagentaFg = "\u001b[35m";
    static readonly MagentaBg = "\u001b[45m";
    static readonly MagentaFgBright = "\u001b[95m";
    static readonly MagentaBgBright = "\u001b[105m";
    static readonly CyanFg = "\u001b[36m";
    static readonly CyanBg = "\u001b[46m";
    static readonly CyanFgBright = "\u001b[96m";
    static readonly CyanBgBright = "\u001b[106m";
    static readonly WhiteFg = "\u001b[37m";
    static readonly WhiteBg = "\u001b[47m";
    static readonly WhiteFgBright = "\u001b[97m";
    static readonly WhiteBgBright = "\u001b[107m";

    // Controls
    static readonly clearScreen = "\u001b[2J";
    static readonly clearScreenDown = "\u001b[J";
    static readonly clearScreenUp = "\u001b[1J";
    static readonly clearLine = "\u001b[2K";
    static readonly clearLineRight = "\u001b[K";
    static readonly clearLineLeft = "\u001b[1K";
    static readonly hideCursor = "\u001b[?25l";
    static readonly showCursor = "\u001b[?25h";
    static readonly cursorUp1 = "\u001b[1A";
    static readonly cursorDown1 = "\u001b[1B";
    static readonly cursorRight1 = "\u001b[1C";
    static readonly cursorLeft1 = "\u001b[1D";
    static readonly cursorNextLine1 = "\u001b[1E";
    static readonly cursorPrevLine1 = "\u001b[1F";
    static readonly cursorColumn1 = "\u001b[1G";

    // Character widths, used by visibleWidth and truncate
    private static readonly zeroWidthRanges: [number, number][] = [
        [0x0000, 0x001F], // C0 controls
        [0x007F, 0x009F], // Delete and C1 controls
        [0x0300, 0x036F], // Combining Diacritical Marks
        [0x1AB0, 0x1AFF], // Combining Diacritical Marks Extended
        [0x1DC0, 0x1DFF], // Combining Diacritical Marks Supplement
        [0x200B, 0x200F], // Zero width space, joiners and direction marks
        [0x2060, 0x2064], // Word joiner and invisible operators
        [0x20D0, 0x20FF], // Combining Diacritical Marks for Symbols
        [0xFE00, 0xFE0F], // Variation Selectors
        [0xFE20, 0xFE2F], // Combining Half Marks
        [0xFEFF, 0xFEFF], // Zero width no-break space
        [0xE0000, 0xE007F], // Tags
        [0xE0100, 0xE01EF], // Variation Selectors Supplement
    ];
    private static readonly wideRanges: [number, number][] = [
        [0x1100, 0x115F], // Hangul Jamo initial consonants
        [0x231A, 0x231B], // Watch and hourglass
        [0x2329, 0x232A], // Angle brackets
        [0x23E9, 0x23EC], // Fast-forward and rewind buttons
        [0x23F0, 0x23F0], // Alarm clock
        [0x23F3, 0x23F3], // Hourglass with flowing sand
        [0x25FD, 0x25FE], // Medium small squares
        [0x2614, 0x2615], // Umbrella with rain and hot beverage
        [0x2648, 0x2653], // Zodiac signs
        [0x267F, 0x267F], // Wheelchair symbol
        [0x2693, 0x2693], // Anchor
        [0x26A1, 0x26A1], // High voltage
        [0x26AA, 0x26AB], // Medium circles
        [0x26BD, 0x26BE], // Soccer ball and baseball
        [0x26C4, 0x26C5], // Snowman and sun behind cloud
        [0x26CE, 0x26CE], // Ophiuchus
        [0x26D4, 0x26D4], // No entry
        [0x26EA, 0x26EA], // Church
        [0x26F2, 0x26F3], // Fountain and flag in hole
        [0x26F5, 0x26F5], // Sailboat
        [0x26FA, 0x26FA], // Tent
        [0x26FD, 0x26FD], // Fuel pump
        [0x2705, 0x2705], // Check mark button
        [0x270A, 0x270B], // Raised fist and hand
        [0x2728, 0x2728], // Sparkles
        [0x274C, 0x274C], // Cross mark
        [0x274E, 0x274E], // Cross mark button
        [0x2753, 0x2755], // Question and exclamation marks
        [0x2757, 0x2757], // Exclamation mark
        [0x2795, 0x2797], // Plus, minus and divide signs
        [0x27B0, 0x27B0], // Curly loop
        [0x27BF, 0x27BF], // Double curly loop
        [0x2B1B, 0x2B1C], // Large squares
        [0x2B50, 0x2B50], // Star
        [0x2B55, 0x2B55], // Hollow red circle
        [0x2E80, 0x303E], // CJK radicals, symbols and punctuation
        [0x3040, 0xA4CF], // Kana, Bopomofo, Hangul compatibility jamo, CJK and Yi
        [0xA960, 0xA97F], // Hangul Jamo Extended-A
        [0xAC00, 0xD7A3], // Hangul Syllables
        [0xF900, 0xFAFF], // CJK Compatibility Ideographs
        [0xFE10, 0xFE19], // Vertical Forms
        [0xFE30, 0xFE6F], // CJK Compatibility Forms and Small Form Variants
        [0xFF00, 0xFF60], // Fullwidth Forms
        [0xFFE0, 0xFFE6], // Fullwidth signs
        [0x16FE0, 0x18DFF], // Ideographic symbols, Tangut and Khitan
        [0x1AFF0, 0x1B2FF], // Kana supplements and Nushu
        [0x1F004, 0x1F004], // Mahjong tile red dragon
        [0x1F0CF, 0x1F0CF], // Joker
        [0x1F18E, 0x1F18E], // AB button
        [0x1F191, 0x1F19A], // Squared words
        [0x1F200, 0x1F2FF], // Enclosed Ideographic Supplement
        [0x1F300, 0x1F64F], // Miscellaneous Symbols and Pictographs, Emoticons
        [0x1F680, 0x1F6FF], // Transport and Map Symbols
        [0x1F7E0, 0x1F7FF], // Large circles and squares
        [0x1F900, 0x1F9FF], // Supplemental Symbols and Pictographs
        [0x1FA70, 0x1FAFF], // Symbols and Pictographs Extended-A
        [0x20000, 0x2FFFD], // Supplementary Ideographic Plane
        [0x30000, 0x3FFFD], // Tertiary Ideographic Plane
    ];

    // GEN END

    // Cursor movement

    /** Moves the cursor up by n lines. Returns an empty string when n <= 0 */
    static cursorUp(n: number): string {
        return Mansil.moveBy(Mansil.cursorUp1, n);
    }

    /** Moves the cursor down by n lines. Returns an empty string when n <= 0 */
    static cursorDown(n: number): string {
        return Mansil.moveBy(Mansil.cursorDown1, n);
    }

    /** Moves the cursor right by n columns. Returns an empty string when n <= 0 */
    static cursorRight(n: number): string {
        return Mansil.moveBy(Mansil.cursorRight1, n);
    }

    /** Moves the cursor left by n columns. Returns an empty string when n <= 0 */
    static cursorLeft(n: number): string {
        return Mansil.moveBy(Mansil.cursorLeft1, n);
    }

    /** Moves the cursor to the start of the line n lines down. Returns an empty string when n <= 0 */
    static cursorNextLine(n: number): string {
        return Mansil.moveBy(Mansil.cursorNextLine1, n);
    }

    /** Moves the cursor to the start of the line n lines up. Returns an empty string when n <= 0 */
    static cursorPrevLine(n: number): string {
        return Mansil.moveBy(Mansil.cursorPrevLine1, n);
    }

    /** Moves the cursor to column n of the current line. Columns are 1-based */
    static cursorColumn(n: number): string {
        return Mansil.cursorColumn1.replace("[1", `[${n}`);
    }

    /** Moves the cursor to the given row and column. Both are 1-based */
    static cursorPosition(row: number, col: number): string {
        return `\u001b[${row};${col}H`;
    }

    // Sets the count of a relative movement, whose constant moves by 1.
    // In CSI sequences a count of 0 means the default of 1, so n <= 0 returns an empty string instead
    private static moveBy(code1: string, n: number): string {
        if (n <= 0) {
            return "";
        }
        return code1.replace("[1", `[${n}`);
    }

    // Styled text

    private static readonly esc = 0x1b;

    /**
     * Returns s without its escape sequences.
     *
     * It removes CSI sequences (styles, colors, cursor movement), OSC sequences (such as hyperlinks)
     * and the other two-character escape sequences
     */
    static strip(s: string): string {
        let out = "";
        for (let i = 0; i < s.length; ) {
            if (s.charCodeAt(i) === Mansil.esc) {
                i += Mansil.escapeLength(s, i);
                continue;
            }
            out += s[i];
            i++;
        }
        return out;
    }

    /**
     * Returns the number of terminal columns that s takes, ignoring its escape sequences.
     *
     * Wide characters, such as CJK and most emoji, take two columns, and so does a narrow character
     * followed by the emoji presentation selector (U+FE0F). Combining marks, zero-width characters
     * and control characters, including tabs and newlines, take none. The characters of a joined
     * emoji (ZWJ sequences, skin tones) are counted one by one, so it can come out wider than the
     * terminal draws it
     */
    static visibleWidth(s: string): number {
        let width = 0;
        let prev = 0;
        for (let i = 0; i < s.length; ) {
            if (s.charCodeAt(i) === Mansil.esc) {
                i += Mansil.escapeLength(s, i);
                continue;
            }
            const cp = s.codePointAt(i)!;
            prev = Mansil.charWidth(cp, prev);
            width += prev;
            i += cp > 0xffff ? 2 : 1;
        }
        return width;
    }

    /**
     * Cuts s to at most width terminal columns, ending it with tail when it is cut.
     *
     * The tail counts towards the width, and is left out when it is wider than width on its own.
     * Only the visible characters after the cut are dropped. Every escape sequence is kept, so the
     * styles and hyperlinks that s closes are still closed. s comes back unchanged when it fits,
     * and a width <= 0 returns an empty string
     */
    static truncate(s: string, width: number, tail = ""): string {
        if (width <= 0) {
            return "";
        }
        if (Mansil.visibleWidth(s) <= width) {
            return s;
        }
        let tailWidth = Mansil.visibleWidth(tail);
        if (tailWidth > width) {
            tail = "";
            tailWidth = 0;
        }
        const room = width - tailWidth;

        let out = "";
        let used = 0;
        let prev = 0;
        let cut = false;
        for (let i = 0; i < s.length; ) {
            if (s.charCodeAt(i) === Mansil.esc) {
                const n = Mansil.escapeLength(s, i);
                out += s.slice(i, i + n);
                i += n;
                continue;
            }
            const cp = s.codePointAt(i)!;
            const size = cp > 0xffff ? 2 : 1;
            if (!cut) {
                const w = Mansil.charWidth(cp, prev);
                if (used + w <= room) {
                    out += s.slice(i, i + size);
                    used += w;
                    prev = w;
                } else {
                    out += tail;
                    cut = true;
                }
            }
            i += size;
        }
        return out;
    }

    // Returns the length of the escape sequence that starts at s[i], which is ESC.
    // - CSI: ESC [, then parameter and intermediate bytes, then a final byte
    // - OSC: ESC ], then anything up to BEL or ESC \
    // - Others: ESC, then intermediate bytes, then a final byte
    // A malformed sequence ends at the first character that cannot belong to it, an unterminated OSC
    // runs to the end of s, and a lone ESC has a length of 1
    private static escapeLength(s: string, i: number): number {
        let j = i + 1;
        if (j === s.length) {
            return 1;
        }
        const kind = s.charCodeAt(j);
        if (kind === 0x5b) {
            // [
            j++;
            while (j < s.length && s.charCodeAt(j) >= 0x20 && s.charCodeAt(j) <= 0x3f) {
                j++;
            }
            if (j < s.length && s.charCodeAt(j) >= 0x40 && s.charCodeAt(j) <= 0x7e) {
                j++;
            }
        } else if (kind === 0x5d) {
            // ]
            for (j++; j < s.length; j++) {
                const c = s.charCodeAt(j);
                if (c === 0x07) {
                    return j + 1 - i;
                }
                if (c === Mansil.esc) {
                    return s.charCodeAt(j + 1) === 0x5c ? j + 2 - i : j - i;
                }
            }
        } else {
            while (j < s.length && s.charCodeAt(j) >= 0x20 && s.charCodeAt(j) <= 0x2f) {
                j++;
            }
            if (j < s.length && s.charCodeAt(j) >= 0x30 && s.charCodeAt(j) <= 0x7e) {
                j++;
            }
        }
        return j - i;
    }

    // Returns the number of terminal columns that the code point takes, after a character that takes prev columns.
    // The emoji presentation selector (U+FE0F) widens a narrow character, such as ⚠, to two columns
    private static charWidth(cp: number, prev: number): number {
        if (cp >= 0x20 && cp < 0x7f) {
            return 1;
        }
        if (cp === 0xfe0f && prev === 1) {
            return 1;
        }
        if (Mansil.zeroWidthRanges.some(([start, end]) => cp >= start && cp <= end)) {
            return 0;
        }
        if (Mansil.wideRanges.some(([start, end]) => cp >= start && cp <= end)) {
            return 2;
        }
        return 1;
    }
}
