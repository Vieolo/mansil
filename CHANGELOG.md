## 0.2.0

2026-10-17

- Added the erase codes `ClearScreenDown`, `ClearScreenUp`, `ClearLineRight` and `ClearLineLeft`
- Added `HideCursor` and `ShowCursor`
- Added `Strip`, `VisibleWidth` and `Truncate`, to measure and cut styled text


**Breaking Changes**

- the relative cursor helpers (`CursorUp`, `CursorDown`, `CursorRight`, `CursorLeft`, `CursorNextLine` and `CursorPrevLine`) return an empty string for `n <= 0`. Before, `n = 0` moved the cursor by one, because a CSI parameter of 0 means 1
- [JS/TS] helpers no longer rely on `this`, so they also work when passed around unbound

## 0.1.2

2026-02-08
- Improved the structure of the package

## 0.1.0

2026-01-28
- Initial release
