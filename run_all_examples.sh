cd example

echo "-=-=-=-= Dart"
cd mansil_dart_example
dart run
cd ..
echo "~~~~~~~~~~~"

echo "-=-=-=-= Go"
cd mansil_go_example
go run .
cd ..
echo "~~~~~~~~~~~"

echo "-=-=-=-= TS"
cd mansil_ts_example
bun start
cd ..
echo "~~~~~~~~~~~"

echo "-=-=-=-= Python"
cd mansil_python_example
uv run python example_main.py
cd ..
echo "~~~~~~~~~~~"

echo "-=-=-=-= Rust"
cd mansil_rust_example
cargo run
cd ..
