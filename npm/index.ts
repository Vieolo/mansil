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
    static readonly clearLine = "\u001b[2K";
    static readonly cursorUp1 = "\u001b[1A";
    static readonly cursorDown1 = "\u001b[1B";
    static readonly cursorRight1 = "\u001b[1C";
    static readonly cursorLeft1 = "\u001b[1D";
    static readonly cursorNextLine1 = "\u001b[1E";
    static readonly cursorPrevLine1 = "\u001b[1F";
    static readonly cursorColumn1 = "\u001b[1G";

    // GEN END

    static cursorUp(n: number): string {
        return this.cursorUp1.replace("[1", `[${n}`);
    }

    static cursorDown(n: number): string {
        return this.cursorDown1.replace("[1", `[${n}`);
    }

    static cursorRight(n: number): string {
        return this.cursorRight1.replace("[1", `[${n}`);
    }

    static cursorLeft(n: number): string {
        return this.cursorLeft1.replace("[1", `[${n}`);
    }

    static cursorNextLine(n: number): string {
        return this.cursorNextLine1.replace("[1", `[${n}`);
    }

    static cursorPrevLine(n: number): string {
        return this.cursorPrevLine1.replace("[1", `[${n}`);
    }

    static cursorColumn(n: number): string {
        return this.cursorColumn1.replace("[1", `[${n}`);
    }

    static cursorPosition(row: number, col: number): string {
        return `\u001b[${row};${col}H`;
    }
}
