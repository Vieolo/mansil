library;

class Mansil {
  // GEN START

  // Styles
  static const String reset = "\u001b[0m";
  static const String bold = "\u001b[1m";
  static const String dim = "\u001b[2m";
  static const String italic = "\u001b[3m";
  static const String underline = "\u001b[4m";
  static const String blink = "\u001b[5m";
  static const String inverse = "\u001b[7m";
  static const String hidden = "\u001b[8m";
  static const String strikethrough = "\u001b[9m";

  // Colors
  static const String blackFg = "\u001b[30m";
  static const String blackBg = "\u001b[40m";
  static const String blackFgBright = "\u001b[90m";
  static const String blackBgBright = "\u001b[100m";
  static const String redFg = "\u001b[31m";
  static const String redBg = "\u001b[41m";
  static const String redFgBright = "\u001b[91m";
  static const String redBgBright = "\u001b[101m";
  static const String greenFg = "\u001b[32m";
  static const String greenBg = "\u001b[42m";
  static const String greenFgBright = "\u001b[92m";
  static const String greenBgBright = "\u001b[102m";
  static const String yellowFg = "\u001b[33m";
  static const String yellowBg = "\u001b[43m";
  static const String yellowFgBright = "\u001b[93m";
  static const String yellowBgBright = "\u001b[103m";
  static const String blueFg = "\u001b[34m";
  static const String blueBg = "\u001b[44m";
  static const String blueFgBright = "\u001b[94m";
  static const String blueBgBright = "\u001b[104m";
  static const String magentaFg = "\u001b[35m";
  static const String magentaBg = "\u001b[45m";
  static const String magentaFgBright = "\u001b[95m";
  static const String magentaBgBright = "\u001b[105m";
  static const String cyanFg = "\u001b[36m";
  static const String cyanBg = "\u001b[46m";
  static const String cyanFgBright = "\u001b[96m";
  static const String cyanBgBright = "\u001b[106m";
  static const String whiteFg = "\u001b[37m";
  static const String whiteBg = "\u001b[47m";
  static const String whiteFgBright = "\u001b[97m";
  static const String whiteBgBright = "\u001b[107m";

  // Controls
  static const String clearScreen = "\u001b[2J";
  static const String clearScreenDown = "\u001b[J";
  static const String clearScreenUp = "\u001b[1J";
  static const String clearLine = "\u001b[2K";
  static const String clearLineRight = "\u001b[K";
  static const String clearLineLeft = "\u001b[1K";
  static const String hideCursor = "\u001b[?25l";
  static const String showCursor = "\u001b[?25h";
  static const String cursorUp1 = "\u001b[1A";
  static const String cursorDown1 = "\u001b[1B";
  static const String cursorRight1 = "\u001b[1C";
  static const String cursorLeft1 = "\u001b[1D";
  static const String cursorNextLine1 = "\u001b[1E";
  static const String cursorPrevLine1 = "\u001b[1F";
  static const String cursorColumn1 = "\u001b[1G";

  // Character widths, used by visibleWidth and truncate
  static const List<(int, int)> _zeroWidthRanges = [
    (0x0000, 0x001F), // C0 controls
    (0x007F, 0x009F), // Delete and C1 controls
    (0x0300, 0x036F), // Combining Diacritical Marks
    (0x1AB0, 0x1AFF), // Combining Diacritical Marks Extended
    (0x1DC0, 0x1DFF), // Combining Diacritical Marks Supplement
    (0x200B, 0x200F), // Zero width space, joiners and direction marks
    (0x2060, 0x2064), // Word joiner and invisible operators
    (0x20D0, 0x20FF), // Combining Diacritical Marks for Symbols
    (0xFE00, 0xFE0F), // Variation Selectors
    (0xFE20, 0xFE2F), // Combining Half Marks
    (0xFEFF, 0xFEFF), // Zero width no-break space
    (0xE0000, 0xE007F), // Tags
    (0xE0100, 0xE01EF), // Variation Selectors Supplement
  ];
  static const List<(int, int)> _wideRanges = [
    (0x1100, 0x115F), // Hangul Jamo initial consonants
    (0x231A, 0x231B), // Watch and hourglass
    (0x2329, 0x232A), // Angle brackets
    (0x23E9, 0x23EC), // Fast-forward and rewind buttons
    (0x23F0, 0x23F0), // Alarm clock
    (0x23F3, 0x23F3), // Hourglass with flowing sand
    (0x25FD, 0x25FE), // Medium small squares
    (0x2614, 0x2615), // Umbrella with rain and hot beverage
    (0x2648, 0x2653), // Zodiac signs
    (0x267F, 0x267F), // Wheelchair symbol
    (0x2693, 0x2693), // Anchor
    (0x26A1, 0x26A1), // High voltage
    (0x26AA, 0x26AB), // Medium circles
    (0x26BD, 0x26BE), // Soccer ball and baseball
    (0x26C4, 0x26C5), // Snowman and sun behind cloud
    (0x26CE, 0x26CE), // Ophiuchus
    (0x26D4, 0x26D4), // No entry
    (0x26EA, 0x26EA), // Church
    (0x26F2, 0x26F3), // Fountain and flag in hole
    (0x26F5, 0x26F5), // Sailboat
    (0x26FA, 0x26FA), // Tent
    (0x26FD, 0x26FD), // Fuel pump
    (0x2705, 0x2705), // Check mark button
    (0x270A, 0x270B), // Raised fist and hand
    (0x2728, 0x2728), // Sparkles
    (0x274C, 0x274C), // Cross mark
    (0x274E, 0x274E), // Cross mark button
    (0x2753, 0x2755), // Question and exclamation marks
    (0x2757, 0x2757), // Exclamation mark
    (0x2795, 0x2797), // Plus, minus and divide signs
    (0x27B0, 0x27B0), // Curly loop
    (0x27BF, 0x27BF), // Double curly loop
    (0x2B1B, 0x2B1C), // Large squares
    (0x2B50, 0x2B50), // Star
    (0x2B55, 0x2B55), // Hollow red circle
    (0x2E80, 0x303E), // CJK radicals, symbols and punctuation
    (0x3040, 0xA4CF), // Kana, Bopomofo, Hangul compatibility jamo, CJK and Yi
    (0xA960, 0xA97F), // Hangul Jamo Extended-A
    (0xAC00, 0xD7A3), // Hangul Syllables
    (0xF900, 0xFAFF), // CJK Compatibility Ideographs
    (0xFE10, 0xFE19), // Vertical Forms
    (0xFE30, 0xFE6F), // CJK Compatibility Forms and Small Form Variants
    (0xFF00, 0xFF60), // Fullwidth Forms
    (0xFFE0, 0xFFE6), // Fullwidth signs
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

  /// Moves the cursor up by [n] lines. Returns an empty string when [n] <= 0
  static String cursorUp(int n) => _moveBy(cursorUp1, n);

  /// Moves the cursor down by [n] lines. Returns an empty string when [n] <= 0
  static String cursorDown(int n) => _moveBy(cursorDown1, n);

  /// Moves the cursor right by [n] columns. Returns an empty string when [n] <= 0
  static String cursorRight(int n) => _moveBy(cursorRight1, n);

  /// Moves the cursor left by [n] columns. Returns an empty string when [n] <= 0
  static String cursorLeft(int n) => _moveBy(cursorLeft1, n);

  /// Moves the cursor to the start of the line [n] lines down. Returns an empty string when [n] <= 0
  static String cursorNextLine(int n) => _moveBy(cursorNextLine1, n);

  /// Moves the cursor to the start of the line [n] lines up. Returns an empty string when [n] <= 0
  static String cursorPrevLine(int n) => _moveBy(cursorPrevLine1, n);

  /// Moves the cursor to column [n] of the current line. Columns are 1-based
  static String cursorColumn(int n) => cursorColumn1.replaceFirst("[1", "[$n");

  /// Moves the cursor to [row] and [col]. Both are 1-based
  static String cursorPosition(int row, int col) => "\u001b[$row;${col}H";

  // Sets the count of a relative movement, whose constant moves by 1.
  // In CSI sequences a count of 0 means the default of 1, so n <= 0 returns an empty string instead
  static String _moveBy(String code1, int n) =>
      n <= 0 ? "" : code1.replaceFirst("[1", "[$n");

  // Styled text

  static const int _esc = 0x1b;

  /// Returns [s] without its escape sequences.
  ///
  /// It removes CSI sequences (styles, colors, cursor movement), OSC sequences (such as hyperlinks)
  /// and the other two-character escape sequences
  static String strip(String s) {
    final out = StringBuffer();
    for (var i = 0; i < s.length;) {
      if (s.codeUnitAt(i) == _esc) {
        i += _escapeLength(s, i);
        continue;
      }
      out.write(s[i]);
      i++;
    }
    return out.toString();
  }

  /// Returns the number of terminal columns that [s] takes, ignoring its escape sequences.
  ///
  /// Wide characters, such as CJK and most emoji, take two columns, and so does a narrow character
  /// followed by the emoji presentation selector (U+FE0F). Combining marks, zero-width characters
  /// and control characters, including tabs and newlines, take none. The characters of a joined
  /// emoji (ZWJ sequences, skin tones) are counted one by one, so it can come out wider than the
  /// terminal draws it
  static int visibleWidth(String s) {
    var width = 0;
    var prev = 0;
    for (var i = 0; i < s.length;) {
      if (s.codeUnitAt(i) == _esc) {
        i += _escapeLength(s, i);
        continue;
      }
      final (cp, size) = _codePointAt(s, i);
      prev = _charWidth(cp, prev);
      width += prev;
      i += size;
    }
    return width;
  }

  /// Cuts [s] to at most [width] terminal columns, ending it with [tail] when it is cut.
  ///
  /// The tail counts towards the width, and is left out when it is wider than [width] on its own.
  /// Only the visible characters after the cut are dropped. Every escape sequence is kept, so the
  /// styles and hyperlinks that [s] closes are still closed. [s] comes back unchanged when it fits,
  /// and a [width] <= 0 returns an empty string
  static String truncate(String s, int width, [String tail = ""]) {
    if (width <= 0) {
      return "";
    }
    if (visibleWidth(s) <= width) {
      return s;
    }
    var tailWidth = visibleWidth(tail);
    if (tailWidth > width) {
      tail = "";
      tailWidth = 0;
    }
    final room = width - tailWidth;

    final out = StringBuffer();
    var used = 0;
    var prev = 0;
    var cut = false;
    for (var i = 0; i < s.length;) {
      if (s.codeUnitAt(i) == _esc) {
        final n = _escapeLength(s, i);
        out.write(s.substring(i, i + n));
        i += n;
        continue;
      }
      final (cp, size) = _codePointAt(s, i);
      if (!cut) {
        final w = _charWidth(cp, prev);
        if (used + w <= room) {
          out.write(s.substring(i, i + size));
          used += w;
          prev = w;
        } else {
          out.write(tail);
          cut = true;
        }
      }
      i += size;
    }
    return out.toString();
  }

  // Returns the length of the escape sequence that starts at s[i], which is ESC.
  // - CSI: ESC [, then parameter and intermediate bytes, then a final byte
  // - OSC: ESC ], then anything up to BEL or ESC \
  // - Others: ESC, then intermediate bytes, then a final byte
  // A malformed sequence ends at the first character that cannot belong to it, an unterminated OSC
  // runs to the end of s, and a lone ESC has a length of 1
  static int _escapeLength(String s, int i) {
    var j = i + 1;
    if (j == s.length) {
      return 1;
    }
    bool inRange(int start, int end) =>
        j < s.length && s.codeUnitAt(j) >= start && s.codeUnitAt(j) <= end;

    final kind = s.codeUnitAt(j);
    if (kind == 0x5b) {
      // [
      j++;
      while (inRange(0x20, 0x3f)) {
        j++;
      }
      if (inRange(0x40, 0x7e)) {
        j++;
      }
    } else if (kind == 0x5d) {
      // ]
      for (j++; j < s.length; j++) {
        final c = s.codeUnitAt(j);
        if (c == 0x07) {
          return j + 1 - i;
        }
        if (c == _esc) {
          final terminated = j + 1 < s.length && s.codeUnitAt(j + 1) == 0x5c;
          return terminated ? j + 2 - i : j - i;
        }
      }
    } else {
      while (inRange(0x20, 0x2f)) {
        j++;
      }
      if (inRange(0x30, 0x7e)) {
        j++;
      }
    }
    return j - i;
  }

  // Returns the code point at s[i] and the number of UTF-16 code units it takes
  static (int, int) _codePointAt(String s, int i) {
    final high = s.codeUnitAt(i);
    if (high >= 0xD800 && high <= 0xDBFF && i + 1 < s.length) {
      final low = s.codeUnitAt(i + 1);
      if (low >= 0xDC00 && low <= 0xDFFF) {
        return (0x10000 + ((high - 0xD800) << 10) + (low - 0xDC00), 2);
      }
    }
    return (high, 1);
  }

  // Returns the number of terminal columns that the code point takes, after a character that takes prev columns.
  // The emoji presentation selector (U+FE0F) widens a narrow character, such as ⚠, to two columns
  static int _charWidth(int cp, int prev) {
    if (cp >= 0x20 && cp < 0x7f) {
      return 1;
    }
    if (cp == 0xFE0F && prev == 1) {
      return 1;
    }
    if (_zeroWidthRanges.any((r) => cp >= r.$1 && cp <= r.$2)) {
      return 0;
    }
    if (_wideRanges.any((r) => cp >= r.$1 && cp <= r.$2)) {
      return 2;
    }
    return 1;
  }
}
