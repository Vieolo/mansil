package mansil

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// GEN START

// Styles
const Reset = "\033[0m"
const Bold = "\033[1m"
const Dim = "\033[2m"
const Italic = "\033[3m"
const Underline = "\033[4m"
const Blink = "\033[5m"
const Inverse = "\033[7m"
const Hidden = "\033[8m"
const Strikethrough = "\033[9m"

// Colors
const BlackFG = "\033[30m"
const BlackBG = "\033[40m"
const BlackFGBright = "\033[90m"
const BlackBGBright = "\033[100m"
const RedFG = "\033[31m"
const RedBG = "\033[41m"
const RedFGBright = "\033[91m"
const RedBGBright = "\033[101m"
const GreenFG = "\033[32m"
const GreenBG = "\033[42m"
const GreenFGBright = "\033[92m"
const GreenBGBright = "\033[102m"
const YellowFG = "\033[33m"
const YellowBG = "\033[43m"
const YellowFGBright = "\033[93m"
const YellowBGBright = "\033[103m"
const BlueFG = "\033[34m"
const BlueBG = "\033[44m"
const BlueFGBright = "\033[94m"
const BlueBGBright = "\033[104m"
const MagentaFG = "\033[35m"
const MagentaBG = "\033[45m"
const MagentaFGBright = "\033[95m"
const MagentaBGBright = "\033[105m"
const CyanFG = "\033[36m"
const CyanBG = "\033[46m"
const CyanFGBright = "\033[96m"
const CyanBGBright = "\033[106m"
const WhiteFG = "\033[37m"
const WhiteBG = "\033[47m"
const WhiteFGBright = "\033[97m"
const WhiteBGBright = "\033[107m"

// Controls
const ClearScreen = "\033[2J"
const ClearScreenDown = "\033[J"
const ClearScreenUp = "\033[1J"
const ClearLine = "\033[2K"
const ClearLineRight = "\033[K"
const ClearLineLeft = "\033[1K"
const HideCursor = "\033[?25l"
const ShowCursor = "\033[?25h"
const CursorUp1 = "\033[1A"
const CursorDown1 = "\033[1B"
const CursorRight1 = "\033[1C"
const CursorLeft1 = "\033[1D"
const CursorNextLine1 = "\033[1E"
const CursorPrevLine1 = "\033[1F"
const CursorColumn1 = "\033[1G"

// Character widths, used by VisibleWidth and Truncate
var zeroWidthRanges = [][2]rune{
	{0x0000, 0x001F},   // C0 controls
	{0x007F, 0x009F},   // Delete and C1 controls
	{0x0300, 0x036F},   // Combining Diacritical Marks
	{0x1AB0, 0x1AFF},   // Combining Diacritical Marks Extended
	{0x1DC0, 0x1DFF},   // Combining Diacritical Marks Supplement
	{0x200B, 0x200F},   // Zero width space, joiners and direction marks
	{0x2060, 0x2064},   // Word joiner and invisible operators
	{0x20D0, 0x20FF},   // Combining Diacritical Marks for Symbols
	{0xFE00, 0xFE0F},   // Variation Selectors
	{0xFE20, 0xFE2F},   // Combining Half Marks
	{0xFEFF, 0xFEFF},   // Zero width no-break space
	{0xE0000, 0xE007F}, // Tags
	{0xE0100, 0xE01EF}, // Variation Selectors Supplement
}
var wideRanges = [][2]rune{
	{0x1100, 0x115F},   // Hangul Jamo initial consonants
	{0x231A, 0x231B},   // Watch and hourglass
	{0x2329, 0x232A},   // Angle brackets
	{0x23E9, 0x23EC},   // Fast-forward and rewind buttons
	{0x23F0, 0x23F0},   // Alarm clock
	{0x23F3, 0x23F3},   // Hourglass with flowing sand
	{0x25FD, 0x25FE},   // Medium small squares
	{0x2614, 0x2615},   // Umbrella with rain and hot beverage
	{0x2648, 0x2653},   // Zodiac signs
	{0x267F, 0x267F},   // Wheelchair symbol
	{0x2693, 0x2693},   // Anchor
	{0x26A1, 0x26A1},   // High voltage
	{0x26AA, 0x26AB},   // Medium circles
	{0x26BD, 0x26BE},   // Soccer ball and baseball
	{0x26C4, 0x26C5},   // Snowman and sun behind cloud
	{0x26CE, 0x26CE},   // Ophiuchus
	{0x26D4, 0x26D4},   // No entry
	{0x26EA, 0x26EA},   // Church
	{0x26F2, 0x26F3},   // Fountain and flag in hole
	{0x26F5, 0x26F5},   // Sailboat
	{0x26FA, 0x26FA},   // Tent
	{0x26FD, 0x26FD},   // Fuel pump
	{0x2705, 0x2705},   // Check mark button
	{0x270A, 0x270B},   // Raised fist and hand
	{0x2728, 0x2728},   // Sparkles
	{0x274C, 0x274C},   // Cross mark
	{0x274E, 0x274E},   // Cross mark button
	{0x2753, 0x2755},   // Question and exclamation marks
	{0x2757, 0x2757},   // Exclamation mark
	{0x2795, 0x2797},   // Plus, minus and divide signs
	{0x27B0, 0x27B0},   // Curly loop
	{0x27BF, 0x27BF},   // Double curly loop
	{0x2B1B, 0x2B1C},   // Large squares
	{0x2B50, 0x2B50},   // Star
	{0x2B55, 0x2B55},   // Hollow red circle
	{0x2E80, 0x303E},   // CJK radicals, symbols and punctuation
	{0x3040, 0xA4CF},   // Kana, Bopomofo, Hangul compatibility jamo, CJK and Yi
	{0xA960, 0xA97F},   // Hangul Jamo Extended-A
	{0xAC00, 0xD7A3},   // Hangul Syllables
	{0xF900, 0xFAFF},   // CJK Compatibility Ideographs
	{0xFE10, 0xFE19},   // Vertical Forms
	{0xFE30, 0xFE6F},   // CJK Compatibility Forms and Small Form Variants
	{0xFF00, 0xFF60},   // Fullwidth Forms
	{0xFFE0, 0xFFE6},   // Fullwidth signs
	{0x16FE0, 0x18DFF}, // Ideographic symbols, Tangut and Khitan
	{0x1AFF0, 0x1B2FF}, // Kana supplements and Nushu
	{0x1F004, 0x1F004}, // Mahjong tile red dragon
	{0x1F0CF, 0x1F0CF}, // Joker
	{0x1F18E, 0x1F18E}, // AB button
	{0x1F191, 0x1F19A}, // Squared words
	{0x1F200, 0x1F2FF}, // Enclosed Ideographic Supplement
	{0x1F300, 0x1F64F}, // Miscellaneous Symbols and Pictographs, Emoticons
	{0x1F680, 0x1F6FF}, // Transport and Map Symbols
	{0x1F7E0, 0x1F7FF}, // Large circles and squares
	{0x1F900, 0x1F9FF}, // Supplemental Symbols and Pictographs
	{0x1FA70, 0x1FAFF}, // Symbols and Pictographs Extended-A
	{0x20000, 0x2FFFD}, // Supplementary Ideographic Plane
	{0x30000, 0x3FFFD}, // Tertiary Ideographic Plane
}

// GEN END

// Cursor movement

// CursorUp moves the cursor up by n lines. It returns an empty string when n <= 0
func CursorUp(n int) string {
	return moveBy(CursorUp1, n)
}

// CursorDown moves the cursor down by n lines. It returns an empty string when n <= 0
func CursorDown(n int) string {
	return moveBy(CursorDown1, n)
}

// CursorRight moves the cursor right by n columns. It returns an empty string when n <= 0
func CursorRight(n int) string {
	return moveBy(CursorRight1, n)
}

// CursorLeft moves the cursor left by n columns. It returns an empty string when n <= 0
func CursorLeft(n int) string {
	return moveBy(CursorLeft1, n)
}

// CursorNextLine moves the cursor to the start of the line n lines down. It returns an empty string when n <= 0
func CursorNextLine(n int) string {
	return moveBy(CursorNextLine1, n)
}

// CursorPrevLine moves the cursor to the start of the line n lines up. It returns an empty string when n <= 0
func CursorPrevLine(n int) string {
	return moveBy(CursorPrevLine1, n)
}

// CursorColumn moves the cursor to column n of the current line. Columns are 1-based
func CursorColumn(n int) string {
	return strings.Replace(CursorColumn1, "[1", "["+strconv.Itoa(n), 1)
}

// CursorPosition moves the cursor to the given row and column. Both are 1-based
func CursorPosition(row, col int) string {
	return fmt.Sprintf("\033[%d;%dH", row, col)
}

// Sets the count of a relative movement, whose constant moves by 1.
// In CSI sequences a count of 0 means the default of 1, so n <= 0 returns an empty string instead
func moveBy(code1 string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Replace(code1, "[1", "["+strconv.Itoa(n), 1)
}

// Styled text

const esc = 0x1b

// Strip returns s without its escape sequences.
// It removes CSI sequences (styles, colors, cursor movement), OSC sequences (such as hyperlinks)
// and the other two-character escape sequences
func Strip(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); {
		if s[i] == esc {
			i += escapeLen(s, i)
			continue
		}
		sb.WriteByte(s[i])
		i++
	}
	return sb.String()
}

// VisibleWidth returns the number of terminal columns that s takes, ignoring its escape sequences.
//
// Wide characters, such as CJK and most emoji, take two columns, and so does a narrow character
// followed by the emoji presentation selector (U+FE0F). Combining marks, zero-width characters
// and control characters, including tabs and newlines, take none. The characters of a joined
// emoji (ZWJ sequences, skin tones) are counted one by one, so it can come out wider than the
// terminal draws it
func VisibleWidth(s string) int {
	width, prev := 0, 0
	for i := 0; i < len(s); {
		if s[i] == esc {
			i += escapeLen(s, i)
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		prev = runeWidth(r, prev)
		width += prev
		i += size
	}
	return width
}

// Truncate cuts s to at most width terminal columns, ending it with tail when it is cut.
//
// The tail counts towards the width, and is left out when it is wider than width on its own.
// Only the visible characters after the cut are dropped. Every escape sequence is kept, so the
// styles and hyperlinks that s closes are still closed. s comes back unchanged when it fits,
// and a width <= 0 returns an empty string
func Truncate(s string, width int, tail string) string {
	if width <= 0 {
		return ""
	}
	if VisibleWidth(s) <= width {
		return s
	}
	tailWidth := VisibleWidth(tail)
	if tailWidth > width {
		tail, tailWidth = "", 0
	}
	room := width - tailWidth

	var sb strings.Builder
	used, prev, cut := 0, 0, false
	for i := 0; i < len(s); {
		if s[i] == esc {
			n := escapeLen(s, i)
			sb.WriteString(s[i : i+n])
			i += n
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if !cut {
			if w := runeWidth(r, prev); used+w <= room {
				sb.WriteString(s[i : i+size])
				used, prev = used+w, w
			} else {
				sb.WriteString(tail)
				cut = true
			}
		}
		i += size
	}
	return sb.String()
}

// Returns the length of the escape sequence that starts at s[i], which is ESC.
//   - CSI: ESC [, then parameter and intermediate bytes, then a final byte
//   - OSC: ESC ], then anything up to BEL or ESC \
//   - Others: ESC, then intermediate bytes, then a final byte
//
// A malformed sequence ends at the first byte that cannot belong to it, an unterminated OSC
// runs to the end of s, and a lone ESC has a length of 1
func escapeLen(s string, i int) int {
	j := i + 1
	if j == len(s) {
		return 1
	}
	switch s[j] {
	case '[':
		j++
		for j < len(s) && s[j] >= 0x20 && s[j] <= 0x3F {
			j++
		}
		if j < len(s) && s[j] >= 0x40 && s[j] <= 0x7E {
			j++
		}
	case ']':
		for j++; j < len(s); j++ {
			if s[j] == 0x07 {
				return j + 1 - i
			}
			if s[j] == esc {
				if j+1 < len(s) && s[j+1] == '\\' {
					return j + 2 - i
				}
				return j - i
			}
		}
	default:
		for j < len(s) && s[j] >= 0x20 && s[j] <= 0x2F {
			j++
		}
		if j < len(s) && s[j] >= 0x30 && s[j] <= 0x7E {
			j++
		}
	}
	return j - i
}

// Returns the number of terminal columns that r takes, after a character that takes prev columns.
// The emoji presentation selector (U+FE0F) widens a narrow character, such as ⚠, to two columns
func runeWidth(r rune, prev int) int {
	if r >= 0x20 && r < 0x7F {
		return 1
	}
	if r == 0xFE0F && prev == 1 {
		return 1
	}
	if inRanges(r, zeroWidthRanges) {
		return 0
	}
	if inRanges(r, wideRanges) {
		return 2
	}
	return 1
}

func inRanges(r rune, ranges [][2]rune) bool {
	for _, rg := range ranges {
		if r >= rg[0] && r <= rg[1] {
			return true
		}
	}
	return false
}
