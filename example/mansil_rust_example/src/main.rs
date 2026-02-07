use std::thread;
use std::time::Duration;
use mansil;

fn main() {
    println!("One");
    thread::sleep(Duration::from_secs(2));
    println!("{}{}{}", mansil::CURSOR_UP_1, mansil::CLEAR_LINE, "two");
}
