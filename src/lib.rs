//! Mansil ANSI Codes

// Styles
pub const RESET: &str = "\x1b[0m";
pub const BOLD: &str = "\x1b[1m";
pub const DIM: &str = "\x1b[2m";
pub const ITALIC: &str = "\x1b[3m";
pub const UNDERLINE: &str = "\x1b[4m";
pub const BLINK: &str = "\x1b[5m";
pub const INVERSE: &str = "\x1b[7m";
pub const HIDDEN: &str = "\x1b[8m";
pub const STRIKETHROUGH: &str = "\x1b[9m";

// Colors
pub const BLACK_FG: &str = "\x1b[30m";
pub const BLACK_BG: &str = "\x1b[40m";
pub const BLACK_FG_BRIGHT: &str = "\x1b[90m";
pub const BLACK_BG_BRIGHT: &str = "\x1b[100m";
pub const RED_FG: &str = "\x1b[31m";
pub const RED_BG: &str = "\x1b[41m";
pub const RED_FG_BRIGHT: &str = "\x1b[91m";
pub const RED_BG_BRIGHT: &str = "\x1b[101m";
pub const GREEN_FG: &str = "\x1b[32m";
pub const GREEN_BG: &str = "\x1b[42m";
pub const GREEN_FG_BRIGHT: &str = "\x1b[92m";
pub const GREEN_BG_BRIGHT: &str = "\x1b[102m";
pub const YELLOW_FG: &str = "\x1b[33m";
pub const YELLOW_BG: &str = "\x1b[43m";
pub const YELLOW_FG_BRIGHT: &str = "\x1b[93m";
pub const YELLOW_BG_BRIGHT: &str = "\x1b[103m";
pub const BLUE_FG: &str = "\x1b[34m";
pub const BLUE_BG: &str = "\x1b[44m";
pub const BLUE_FG_BRIGHT: &str = "\x1b[94m";
pub const BLUE_BG_BRIGHT: &str = "\x1b[104m";
pub const MAGENTA_FG: &str = "\x1b[35m";
pub const MAGENTA_BG: &str = "\x1b[45m";
pub const MAGENTA_FG_BRIGHT: &str = "\x1b[95m";
pub const MAGENTA_BG_BRIGHT: &str = "\x1b[105m";
pub const CYAN_FG: &str = "\x1b[36m";
pub const CYAN_BG: &str = "\x1b[46m";
pub const CYAN_FG_BRIGHT: &str = "\x1b[96m";
pub const CYAN_BG_BRIGHT: &str = "\x1b[106m";
pub const WHITE_FG: &str = "\x1b[37m";
pub const WHITE_BG: &str = "\x1b[47m";
pub const WHITE_FG_BRIGHT: &str = "\x1b[97m";
pub const WHITE_BG_BRIGHT: &str = "\x1b[107m";

// Controls
pub const CLEAR_SCREEN: &str = "\x1b[2J";
pub const CLEAR_LINE: &str = "\x1b[2K";
pub fn cursor_up(n: u32) -> String {
    format!("\x1b[{}A", n)
}
pub fn cursor_down(n: u32) -> String {
    format!("\x1b[{}B", n)
}
pub fn cursor_right(n: u32) -> String {
    format!("\x1b[{}C", n)
}
pub fn cursor_left(n: u32) -> String {
    format!("\x1b[{}D", n)
}
pub fn cursor_next_line(n: u32) -> String {
    format!("\x1b[{}E", n)
}
pub fn cursor_prev_line(n: u32) -> String {
    format!("\x1b[{}F", n)
}
pub fn cursor_column(n: u32) -> String {
    format!("\x1b[{}G", n)
}
pub fn cursor_position(row: u32, col: u32) -> String {
    format!("\x1b[{};{}H", row, col)
}
