use std::thread;
use std::time::Duration;
use mansil::{CURSOR_UP_1, CLEAR_LINE};

fn main() {
    println!("One");
    thread::sleep(Duration::from_secs(2));
    println!("{}{}{}", CURSOR_UP_1, CLEAR_LINE, "two");
}
