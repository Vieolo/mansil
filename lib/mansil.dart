// Styles
const String reset = "\u001b[0m";
const String bold = "\u001b[1m";
const String dim = "\u001b[2m";
const String italic = "\u001b[3m";
const String underline = "\u001b[4m";
const String blink = "\u001b[5m";
const String inverse = "\u001b[7m";
const String hidden = "\u001b[8m";
const String strikethrough = "\u001b[9m";

// Colors
const String blackFg = "\u001b[30m";
const String blackBg = "\u001b[40m";
const String blackFgBright = "\u001b[90m";
const String blackBgBright = "\u001b[100m";
const String redFg = "\u001b[31m";
const String redBg = "\u001b[41m";
const String redFgBright = "\u001b[91m";
const String redBgBright = "\u001b[101m";
const String greenFg = "\u001b[32m";
const String greenBg = "\u001b[42m";
const String greenFgBright = "\u001b[92m";
const String greenBgBright = "\u001b[102m";
const String yellowFg = "\u001b[33m";
const String yellowBg = "\u001b[43m";
const String yellowFgBright = "\u001b[93m";
const String yellowBgBright = "\u001b[103m";
const String blueFg = "\u001b[34m";
const String blueBg = "\u001b[44m";
const String blueFgBright = "\u001b[94m";
const String blueBgBright = "\u001b[104m";
const String magentaFg = "\u001b[35m";
const String magentaBg = "\u001b[45m";
const String magentaFgBright = "\u001b[95m";
const String magentaBgBright = "\u001b[105m";
const String cyanFg = "\u001b[36m";
const String cyanBg = "\u001b[46m";
const String cyanFgBright = "\u001b[96m";
const String cyanBgBright = "\u001b[106m";
const String whiteFg = "\u001b[37m";
const String whiteBg = "\u001b[47m";
const String whiteFgBright = "\u001b[97m";
const String whiteBgBright = "\u001b[107m";

// Controls
const String clearScreen = "\u001b[2J";
const String clearLine = "\u001b[2K";
String cursorUp(int n) => "\u001b[${n}A";
String cursorDown(int n) => "\u001b[${n}B";
String cursorRight(int n) => "\u001b[${n}C";
String cursorLeft(int n) => "\u001b[${n}D";
String cursorNextLine(int n) => "\u001b[${n}E";
String cursorPrevLine(int n) => "\u001b[${n}F";
String cursorColumn(int n) => "\u001b[${n}G";
String cursorPosition(int row, int col) => "\u001b[${row};${col}H";
