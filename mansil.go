package mansil

import (
	"fmt"
	"strconv"
	"strings"
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
const ClearLine = "\033[2K"
const CursorUp1 = "\033[1A"
const CursorDown1 = "\033[1B"
const CursorRight1 = "\033[1C"
const CursorLeft1 = "\033[1D"
const CursorNextLine1 = "\033[1E"
const CursorPrevLine1 = "\033[1F"
const CursorColumn1 = "\033[1G"

// GEN END

func CursorUp(n int) string {
	return strings.Replace(CursorUp1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorDown(n int) string {
	return strings.Replace(CursorDown1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorRight(n int) string {
	return strings.Replace(CursorRight1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorLeft(n int) string {
	return strings.Replace(CursorLeft1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorNextLine(n int) string {
	return strings.Replace(CursorNextLine1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorPrevLine(n int) string {
	return strings.Replace(CursorPrevLine1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorColumn(n int) string {
	return strings.Replace(CursorColumn1, "[1", "["+strconv.Itoa(n), 1)
}

func CursorPosition(row, col int) string {
	return fmt.Sprintf("\033[%d;%dH", row, col)
}
