import time
from mansil import CURSOR_UP_1, CLEAR_LINE


def main():
    print("One")
    time.sleep(2)
    print(CURSOR_UP_1 + CLEAR_LINE + "two")


if __name__ == "__main__":
    main()
