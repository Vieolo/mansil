// Styles
export const Reset = "\u001b[0m";
export const Bold = "\u001b[1m";
export const Dim = "\u001b[2m";
export const Italic = "\u001b[3m";
export const Underline = "\u001b[4m";
export const Blink = "\u001b[5m";
export const Inverse = "\u001b[7m";
export const Hidden = "\u001b[8m";
export const Strikethrough = "\u001b[9m";

// Colors
export const BlackFg = "\u001b[30m";
export const BlackBg = "\u001b[40m";
export const BlackFgBright = "\u001b[90m";
export const BlackBgBright = "\u001b[100m";
export const RedFg = "\u001b[31m";
export const RedBg = "\u001b[41m";
export const RedFgBright = "\u001b[91m";
export const RedBgBright = "\u001b[101m";
export const GreenFg = "\u001b[32m";
export const GreenBg = "\u001b[42m";
export const GreenFgBright = "\u001b[92m";
export const GreenBgBright = "\u001b[102m";
export const YellowFg = "\u001b[33m";
export const YellowBg = "\u001b[43m";
export const YellowFgBright = "\u001b[93m";
export const YellowBgBright = "\u001b[103m";
export const BlueFg = "\u001b[34m";
export const BlueBg = "\u001b[44m";
export const BlueFgBright = "\u001b[94m";
export const BlueBgBright = "\u001b[104m";
export const MagentaFg = "\u001b[35m";
export const MagentaBg = "\u001b[45m";
export const MagentaFgBright = "\u001b[95m";
export const MagentaBgBright = "\u001b[105m";
export const CyanFg = "\u001b[36m";
export const CyanBg = "\u001b[46m";
export const CyanFgBright = "\u001b[96m";
export const CyanBgBright = "\u001b[106m";
export const WhiteFg = "\u001b[37m";
export const WhiteBg = "\u001b[47m";
export const WhiteFgBright = "\u001b[97m";
export const WhiteBgBright = "\u001b[107m";

// Controls
export const clearScreen = "\u001b[2J";
export const clearLine = "\u001b[2K";
export function cursorUp(n: number): string {
    return `\u001b[${n}A`;
}
export function cursorDown(n: number): string {
    return `\u001b[${n}B`;
}
export function cursorRight(n: number): string {
    return `\u001b[${n}C`;
}
export function cursorLeft(n: number): string {
    return `\u001b[${n}D`;
}
export function cursorNextLine(n: number): string {
    return `\u001b[${n}E`;
}
export function cursorPrevLine(n: number): string {
    return `\u001b[${n}F`;
}
export function cursorColumn(n: number): string {
    return `\u001b[${n}G`;
}
export function cursorPosition(row: number, col: number): string {
    return `\u001b[${row};${col}H`;
}
