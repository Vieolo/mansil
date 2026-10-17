# GEN START

# Styles
RESET = "\033[0m"
BOLD = "\033[1m"
DIM = "\033[2m"
ITALIC = "\033[3m"
UNDERLINE = "\033[4m"
BLINK = "\033[5m"
INVERSE = "\033[7m"
HIDDEN = "\033[8m"
STRIKETHROUGH = "\033[9m"

# Colors
BLACK_FG = "\033[30m"
BLACK_BG = "\033[40m"
BLACK_FG_BRIGHT = "\033[90m"
BLACK_BG_BRIGHT = "\033[100m"
RED_FG = "\033[31m"
RED_BG = "\033[41m"
RED_FG_BRIGHT = "\033[91m"
RED_BG_BRIGHT = "\033[101m"
GREEN_FG = "\033[32m"
GREEN_BG = "\033[42m"
GREEN_FG_BRIGHT = "\033[92m"
GREEN_BG_BRIGHT = "\033[102m"
YELLOW_FG = "\033[33m"
YELLOW_BG = "\033[43m"
YELLOW_FG_BRIGHT = "\033[93m"
YELLOW_BG_BRIGHT = "\033[103m"
BLUE_FG = "\033[34m"
BLUE_BG = "\033[44m"
BLUE_FG_BRIGHT = "\033[94m"
BLUE_BG_BRIGHT = "\033[104m"
MAGENTA_FG = "\033[35m"
MAGENTA_BG = "\033[45m"
MAGENTA_FG_BRIGHT = "\033[95m"
MAGENTA_BG_BRIGHT = "\033[105m"
CYAN_FG = "\033[36m"
CYAN_BG = "\033[46m"
CYAN_FG_BRIGHT = "\033[96m"
CYAN_BG_BRIGHT = "\033[106m"
WHITE_FG = "\033[37m"
WHITE_BG = "\033[47m"
WHITE_FG_BRIGHT = "\033[97m"
WHITE_BG_BRIGHT = "\033[107m"

# Controls
CLEAR_SCREEN = "\033[2J"
CLEAR_SCREEN_DOWN = "\033[J"
CLEAR_SCREEN_UP = "\033[1J"
CLEAR_LINE = "\033[2K"
CLEAR_LINE_RIGHT = "\033[K"
CLEAR_LINE_LEFT = "\033[1K"
HIDE_CURSOR = "\033[?25l"
SHOW_CURSOR = "\033[?25h"
CURSOR_UP_1 = "\033[1A"
CURSOR_DOWN_1 = "\033[1B"
CURSOR_RIGHT_1 = "\033[1C"
CURSOR_LEFT_1 = "\033[1D"
CURSOR_NEXT_LINE_1 = "\033[1E"
CURSOR_PREV_LINE_1 = "\033[1F"
CURSOR_COLUMN_1 = "\033[1G"

# Character widths, used by visible_width and truncate
_ZERO_WIDTH_RANGES = (
    (0x0000, 0x001F),  # C0 controls
    (0x007F, 0x009F),  # Delete and C1 controls
    (0x0300, 0x036F),  # Combining Diacritical Marks
    (0x1AB0, 0x1AFF),  # Combining Diacritical Marks Extended
    (0x1DC0, 0x1DFF),  # Combining Diacritical Marks Supplement
    (0x200B, 0x200F),  # Zero width space, joiners and direction marks
    (0x2060, 0x2064),  # Word joiner and invisible operators
    (0x20D0, 0x20FF),  # Combining Diacritical Marks for Symbols
    (0xFE00, 0xFE0F),  # Variation Selectors
    (0xFE20, 0xFE2F),  # Combining Half Marks
    (0xFEFF, 0xFEFF),  # Zero width no-break space
    (0xE0000, 0xE007F),  # Tags
    (0xE0100, 0xE01EF),  # Variation Selectors Supplement
)
_WIDE_RANGES = (
    (0x1100, 0x115F),  # Hangul Jamo initial consonants
    (0x231A, 0x231B),  # Watch and hourglass
    (0x2329, 0x232A),  # Angle brackets
    (0x23E9, 0x23EC),  # Fast-forward and rewind buttons
    (0x23F0, 0x23F0),  # Alarm clock
    (0x23F3, 0x23F3),  # Hourglass with flowing sand
    (0x25FD, 0x25FE),  # Medium small squares
    (0x2614, 0x2615),  # Umbrella with rain and hot beverage
    (0x2648, 0x2653),  # Zodiac signs
    (0x267F, 0x267F),  # Wheelchair symbol
    (0x2693, 0x2693),  # Anchor
    (0x26A1, 0x26A1),  # High voltage
    (0x26AA, 0x26AB),  # Medium circles
    (0x26BD, 0x26BE),  # Soccer ball and baseball
    (0x26C4, 0x26C5),  # Snowman and sun behind cloud
    (0x26CE, 0x26CE),  # Ophiuchus
    (0x26D4, 0x26D4),  # No entry
    (0x26EA, 0x26EA),  # Church
    (0x26F2, 0x26F3),  # Fountain and flag in hole
    (0x26F5, 0x26F5),  # Sailboat
    (0x26FA, 0x26FA),  # Tent
    (0x26FD, 0x26FD),  # Fuel pump
    (0x2705, 0x2705),  # Check mark button
    (0x270A, 0x270B),  # Raised fist and hand
    (0x2728, 0x2728),  # Sparkles
    (0x274C, 0x274C),  # Cross mark
    (0x274E, 0x274E),  # Cross mark button
    (0x2753, 0x2755),  # Question and exclamation marks
    (0x2757, 0x2757),  # Exclamation mark
    (0x2795, 0x2797),  # Plus, minus and divide signs
    (0x27B0, 0x27B0),  # Curly loop
    (0x27BF, 0x27BF),  # Double curly loop
    (0x2B1B, 0x2B1C),  # Large squares
    (0x2B50, 0x2B50),  # Star
    (0x2B55, 0x2B55),  # Hollow red circle
    (0x2E80, 0x303E),  # CJK radicals, symbols and punctuation
    (0x3040, 0xA4CF),  # Kana, Bopomofo, Hangul compatibility jamo, CJK and Yi
    (0xA960, 0xA97F),  # Hangul Jamo Extended-A
    (0xAC00, 0xD7A3),  # Hangul Syllables
    (0xF900, 0xFAFF),  # CJK Compatibility Ideographs
    (0xFE10, 0xFE19),  # Vertical Forms
    (0xFE30, 0xFE6F),  # CJK Compatibility Forms and Small Form Variants
    (0xFF00, 0xFF60),  # Fullwidth Forms
    (0xFFE0, 0xFFE6),  # Fullwidth signs
    (0x16FE0, 0x18DFF),  # Ideographic symbols, Tangut and Khitan
    (0x1AFF0, 0x1B2FF),  # Kana supplements and Nushu
    (0x1F004, 0x1F004),  # Mahjong tile red dragon
    (0x1F0CF, 0x1F0CF),  # Joker
    (0x1F18E, 0x1F18E),  # AB button
    (0x1F191, 0x1F19A),  # Squared words
    (0x1F200, 0x1F2FF),  # Enclosed Ideographic Supplement
    (0x1F300, 0x1F64F),  # Miscellaneous Symbols and Pictographs, Emoticons
    (0x1F680, 0x1F6FF),  # Transport and Map Symbols
    (0x1F7E0, 0x1F7FF),  # Large circles and squares
    (0x1F900, 0x1F9FF),  # Supplemental Symbols and Pictographs
    (0x1FA70, 0x1FAFF),  # Symbols and Pictographs Extended-A
    (0x20000, 0x2FFFD),  # Supplementary Ideographic Plane
    (0x30000, 0x3FFFD),  # Tertiary Ideographic Plane
)

# GEN END

# Cursor movement

def cursor_up(n: int) -> str:
    """Moves the cursor up by n lines. Returns an empty string when n <= 0"""
    return _move_by(CURSOR_UP_1, n)

def cursor_down(n: int) -> str:
    """Moves the cursor down by n lines. Returns an empty string when n <= 0"""
    return _move_by(CURSOR_DOWN_1, n)

def cursor_right(n: int) -> str:
    """Moves the cursor right by n columns. Returns an empty string when n <= 0"""
    return _move_by(CURSOR_RIGHT_1, n)

def cursor_left(n: int) -> str:
    """Moves the cursor left by n columns. Returns an empty string when n <= 0"""
    return _move_by(CURSOR_LEFT_1, n)

def cursor_next_line(n: int) -> str:
    """Moves the cursor to the start of the line n lines down. Returns an empty string when n <= 0"""
    return _move_by(CURSOR_NEXT_LINE_1, n)

def cursor_prev_line(n: int) -> str:
    """Moves the cursor to the start of the line n lines up. Returns an empty string when n <= 0"""
    return _move_by(CURSOR_PREV_LINE_1, n)

def cursor_column(n: int) -> str:
    """Moves the cursor to column n of the current line. Columns are 1-based"""
    return CURSOR_COLUMN_1.replace("[1", f"[{n}")

def cursor_position(row: int, col: int) -> str:
    """Moves the cursor to the given row and column. Both are 1-based"""
    return f"\033[{row};{col}H"

def _move_by(code_1: str, n: int) -> str:
    # Sets the count of a relative movement, whose constant moves by 1.
    # In CSI sequences a count of 0 means the default of 1, so n <= 0 returns an empty string instead
    if n <= 0:
        return ""
    return code_1.replace("[1", f"[{n}")

# Styled text

_ESC = "\033"

def strip(s: str) -> str:
    """Returns s without its escape sequences.

    It removes CSI sequences (styles, colors, cursor movement), OSC sequences (such as hyperlinks)
    and the other two-character escape sequences
    """
    out = []
    i = 0
    while i < len(s):
        if s[i] == _ESC:
            i += _escape_len(s, i)
            continue
        out.append(s[i])
        i += 1
    return "".join(out)

def visible_width(s: str) -> int:
    """Returns the number of terminal columns that s takes, ignoring its escape sequences.

    Wide characters, such as CJK and most emoji, take two columns, and so does a narrow character
    followed by the emoji presentation selector (U+FE0F). Combining marks, zero-width characters
    and control characters, including tabs and newlines, take none. The characters of a joined
    emoji (ZWJ sequences, skin tones) are counted one by one, so it can come out wider than the
    terminal draws it
    """
    width = prev = 0
    i = 0
    while i < len(s):
        if s[i] == _ESC:
            i += _escape_len(s, i)
            continue
        prev = _char_width(ord(s[i]), prev)
        width += prev
        i += 1
    return width

def truncate(s: str, width: int, tail: str = "") -> str:
    """Cuts s to at most width terminal columns, ending it with tail when it is cut.

    The tail counts towards the width, and is left out when it is wider than width on its own.
    Only the visible characters after the cut are dropped. Every escape sequence is kept, so the
    styles and hyperlinks that s closes are still closed. s comes back unchanged when it fits,
    and a width <= 0 returns an empty string
    """
    if width <= 0:
        return ""
    if visible_width(s) <= width:
        return s
    tail_width = visible_width(tail)
    if tail_width > width:
        tail, tail_width = "", 0
    room = width - tail_width

    out = []
    used, prev, cut = 0, 0, False
    i = 0
    while i < len(s):
        if s[i] == _ESC:
            n = _escape_len(s, i)
            out.append(s[i:i + n])
            i += n
            continue
        if not cut:
            w = _char_width(ord(s[i]), prev)
            if used + w <= room:
                out.append(s[i])
                used, prev = used + w, w
            else:
                out.append(tail)
                cut = True
        i += 1
    return "".join(out)

def _escape_len(s: str, i: int) -> int:
    # Returns the length of the escape sequence that starts at s[i], which is ESC.
    # - CSI: ESC [, then parameter and intermediate bytes, then a final byte
    # - OSC: ESC ], then anything up to BEL or ESC \
    # - Others: ESC, then intermediate bytes, then a final byte
    # A malformed sequence ends at the first character that cannot belong to it, an unterminated OSC
    # runs to the end of s, and a lone ESC has a length of 1
    j = i + 1
    if j == len(s):
        return 1
    if s[j] == "[":
        j += 1
        while j < len(s) and "\x20" <= s[j] <= "\x3f":
            j += 1
        if j < len(s) and "\x40" <= s[j] <= "\x7e":
            j += 1
    elif s[j] == "]":
        j += 1
        while j < len(s):
            if s[j] == "\x07":
                return j + 1 - i
            if s[j] == _ESC:
                return j + 2 - i if s[j + 1:j + 2] == "\\" else j - i
            j += 1
    else:
        while j < len(s) and "\x20" <= s[j] <= "\x2f":
            j += 1
        if j < len(s) and "\x30" <= s[j] <= "\x7e":
            j += 1
    return j - i

def _char_width(cp: int, prev: int) -> int:
    # Returns the number of terminal columns that the code point takes, after a character that takes prev columns.
    # The emoji presentation selector (U+FE0F) widens a narrow character, such as ⚠, to two columns
    if 0x20 <= cp < 0x7F:
        return 1
    if cp == 0xFE0F and prev == 1:
        return 1
    if any(start <= cp <= end for start, end in _ZERO_WIDTH_RANGES):
        return 0
    if any(start <= cp <= end for start, end in _WIDE_RANGES):
        return 2
    return 1