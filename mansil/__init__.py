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
CLEAR_LINE = "\033[2K"
CURSOR_UP_1 = "\033[1A"
CURSOR_DOWN_1 = "\033[1B"
CURSOR_RIGHT_1 = "\033[1C"
CURSOR_LEFT_1 = "\033[1D"
CURSOR_NEXT_LINE_1 = "\033[1E"
CURSOR_PREV_LINE_1 = "\033[1F"
CURSOR_COLUMN_1 = "\033[1G"

# GEN END

def cursor_up(n: int) -> str:
    return CURSOR_UP_1.replace("[1", f"[{n}")

def cursor_down(n: int) -> str:
    return CURSOR_DOWN_1.replace("[1", f"[{n}")

def cursor_right(n: int) -> str:
    return CURSOR_RIGHT_1.replace("[1", f"[{n}")

def cursor_left(n: int) -> str:
    return CURSOR_LEFT_1.replace("[1", f"[{n}")

def cursor_next_line(n: int) -> str:
    return CURSOR_NEXT_LINE_1.replace("[1", f"[{n}")

def cursor_prev_line(n: int) -> str:
    return CURSOR_PREV_LINE_1.replace("[1", f"[{n}")

def cursor_column(n: int) -> str:
    return CURSOR_COLUMN_1.replace("[1", f"[{n}")

def cursor_position(row: int, col: int) -> str:
    return f"\033[{row};{col}H"