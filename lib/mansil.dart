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
  static const String clearLine = "\u001b[2K";
  static const String cursorUp1 = "\u001b[1A";
  static const String cursorDown1 = "\u001b[1B";
  static const String cursorRight1 = "\u001b[1C";
  static const String cursorLeft1 = "\u001b[1D";
  static const String cursorNextLine1 = "\u001b[1E";
  static const String cursorPrevLine1 = "\u001b[1F";
  static const String cursorColumn1 = "\u001b[1G";

  // GEN END

  static String cursorUp(int n) => cursorUp1.replaceFirst("[1", "[$n");
  static String cursorDown(int n) => cursorDown1.replaceFirst("[1", "[$n");
  static String cursorRight(int n) => cursorRight1.replaceFirst("[1", "[$n");
  static String cursorLeft(int n) => cursorLeft1.replaceFirst("[1", "[$n");
  static String cursorNextLine(int n) => cursorNextLine1.replaceFirst("[1", "[$n");
  static String cursorPrevLine(int n) => cursorPrevLine1.replaceFirst("[1", "[$n");
  static String cursorColumn(int n) => cursorColumn1.replaceFirst("[1", "[$n");
  static String cursorPosition(int row, int col) => "\u001b[$row;${col}H";
}
