package mansil

import "fmt"

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
func CursorUp(n int) string {
	return fmt.Sprintf("\033[%dA", n)
}
const CursorUp1 = "\033[1A"
func CursorDown(n int) string {
	return fmt.Sprintf("\033[%dB", n)
}
const CursorDown1 = "\033[1B"
func CursorRight(n int) string {
	return fmt.Sprintf("\033[%dC", n)
}
const CursorRight1 = "\033[1C"
func CursorLeft(n int) string {
	return fmt.Sprintf("\033[%dD", n)
}
const CursorLeft1 = "\033[1D"
func CursorNextLine(n int) string {
	return fmt.Sprintf("\033[%dE", n)
}
const CursorNextLine1 = "\033[1E"
func CursorPrevLine(n int) string {
	return fmt.Sprintf("\033[%dF", n)
}
const CursorPrevLine1 = "\033[1F"
func CursorColumn(n int) string {
	return fmt.Sprintf("\033[%dG", n)
}
const CursorColumn1 = "\033[1G"
func CursorPosition(row, col int) string {
	return fmt.Sprintf("\033[%d;%dH", row, col)
}
