package generator

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

// Main generator function to generate the constants for all languages
func Generate() error {
	src, err := ReadSource("ansi_codes.json")
	if err != nil {
		return fmt.Errorf("failed to read source: %w", err)
	}

	if err := generateGo(src, filepath.Join("mansil.go")); err != nil {
		return err
	}
	if err := generatePython(src, filepath.Join("mansil", "__init__.py")); err != nil {
		return err
	}
	if err := generateRust(src, filepath.Join("src", "lib.rs")); err != nil {
		return err
	}
	if err := generateTS(src, filepath.Join("npm", "index.ts")); err != nil {
		return err
	}
	if err := generateDart(src, filepath.Join("lib", "mansil.dart")); err != nil {
		return err
	}
	return nil
}

// Reads `ansi.json` as the source of the ANSI codes
func ReadSource(path string) (Source, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Source{}, err
	}
	var s Source
	if err := json.Unmarshal(b, &s); err != nil {
		return Source{}, err
	}
	if err := validateWidths(s.Widths); err != nil {
		return Source{}, err
	}
	return s, nil
}

// Makes sure that every width range is valid, and that no code point is in more than one range.
// A code point in two ranges would get the width of whichever range the helpers check first
func validateWidths(w Widths) error {
	all := slices.Concat(w.Zero, w.Wide)
	for _, r := range all {
		if r.To < r.From {
			return fmt.Errorf("width range %q ends before it starts", r.Name)
		}
	}
	slices.SortFunc(all, func(a, b CharRange) int { return cmp.Compare(a.From, b.From) })
	for i := 1; i < len(all); i++ {
		if all[i].From <= all[i-1].To {
			return fmt.Errorf("width ranges %q and %q overlap", all[i-1].Name, all[i].Name)
		}
	}
	return nil
}

// Helpers

func toSnake(s string) string {
	var sb strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			sb.WriteRune('_')
		}
		sb.WriteRune(unicode.ToUpper(r))
	}
	return sb.String()
}

func toLowerCamel(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func toUpperSnake(s string) string {
	return strings.ToUpper(toSnake(s))
}

// Writes one line per range: its bounds formatted with `boundsFmt`, then its name as a comment.
// With `align`, the comments line up in one column, the way gofmt and rustfmt format them
func writeRanges(sb *strings.Builder, ranges []CharRange, indent, boundsFmt, comment string, align bool) {
	bounds := make([]string, len(ranges))
	longest := 0
	for i, r := range ranges {
		bounds[i] = fmt.Sprintf(boundsFmt, r.From, r.To)
		longest = max(longest, len(bounds[i]))
	}
	for i, r := range ranges {
		pad := ""
		if align {
			pad = strings.Repeat(" ", longest-len(bounds[i]))
		}
		sb.WriteString(indent + bounds[i] + pad + comment + r.Name + "\n")
	}
}

// All constants, in all languages, will be placed in between the "GEN START" and "GEN END" comments
// and the rest of the files will remain untouched. This allows us to manually modify each file
// as needed without being overwritten by the generator
//
// This function will take the new content and insert it in the boundary of the "GEN *" comments
func replaceGenSection(path string, newContent string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	content := string(data)

	// Find GEN markers
	// Since the comments in each language might be differnt (looking at you Python!),
	// we will look at the comment value directly
	startIdx := strings.Index(content, "GEN START")
	endIdx := strings.Index(content, "GEN END")

	if startIdx == -1 {
		return fmt.Errorf("GEN START marker not found in %s", path)
	}
	if endIdx == -1 {
		return fmt.Errorf("GEN END marker not found in %s", path)
	}

	// Find the end of the line containing GEN START
	startLineEnd := strings.Index(content[startIdx:], "\n")
	if startLineEnd == -1 {
		return fmt.Errorf("no newline after GEN START in %s", path)
	}
	startLineEnd += startIdx

	// Find the start of the line containing GEN END (go back to find the line start)
	endLineStart := endIdx
	for endLineStart > 0 && content[endLineStart-1] != '\n' {
		endLineStart--
	}

	// Build the new file content
	var result strings.Builder
	result.WriteString(content[:startLineEnd+1]) // Everything up to and including GEN START line
	result.WriteString(newContent)               // New generated content
	result.WriteString(content[endLineStart:])   // GEN END line and everything after

	return os.WriteFile(path, []byte(result.String()), 0644)
}

// Generators

// Generates the exported Go code.
func generateGo(s Source, path string) error {
	fmt.Println("Generating for Go...")
	var sb strings.Builder

	sb.WriteString("\n// Styles\n")
	for _, st := range s.Styles {
		sb.WriteString(fmt.Sprintf("const %s = \"\\033[%sm\"\n", st.Name, st.Code))
	}

	sb.WriteString("\n// Colors\n")
	for _, c := range s.Colors {
		sb.WriteString(fmt.Sprintf("const %sFG = \"\\033[%sm\"\n", c.Name, c.Fg))
		sb.WriteString(fmt.Sprintf("const %sBG = \"\\033[%sm\"\n", c.Name, c.Bg))
		sb.WriteString(fmt.Sprintf("const %sFGBright = \"\\033[%sm\"\n", c.Name, c.FgBright))
		sb.WriteString(fmt.Sprintf("const %sBGBright = \"\\033[%sm\"\n", c.Name, c.BgBright))
	}

	sb.WriteString("\n// Controls\n")
	for _, c := range s.Controls {
		if c.Seq != "" {
			sb.WriteString(fmt.Sprintf("const %s = \"%s\"\n", c.Name, c.Seq))
		} else {
			// Only generate the _1 constant for controls with fmt (single parameter)
			count := strings.Count(c.Fmt, "%d")
			if count == 1 {
				singleFmt := strings.Replace(c.Fmt, "%d", "1", 1)
				sb.WriteString(fmt.Sprintf("const %s1 = \"%s\"\n", c.Name, singleFmt))
			}
		}
	}

	sb.WriteString("\n// Character widths, used by VisibleWidth and Truncate\n")
	sb.WriteString("var zeroWidthRanges = [][2]rune{\n")
	writeRanges(&sb, s.Widths.Zero, "\t", "{0x%04X, 0x%04X},", " // ", true)
	sb.WriteString("}\n")
	sb.WriteString("var wideRanges = [][2]rune{\n")
	writeRanges(&sb, s.Widths.Wide, "\t", "{0x%04X, 0x%04X},", " // ", true)
	sb.WriteString("}\n")

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

// Generates the exported python code
func generatePython(s Source, path string) error {
	fmt.Println("Generating for Python...")
	var sb strings.Builder

	sb.WriteString("\n# Styles\n")
	for _, st := range s.Styles {
		sb.WriteString(fmt.Sprintf("%s = \"\\033[%sm\"\n", toUpperSnake(st.Name), st.Code))
	}

	sb.WriteString("\n# Colors\n")
	for _, c := range s.Colors {
		sb.WriteString(fmt.Sprintf("%s_FG = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.Fg))
		sb.WriteString(fmt.Sprintf("%s_BG = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.Bg))
		sb.WriteString(fmt.Sprintf("%s_FG_BRIGHT = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.FgBright))
		sb.WriteString(fmt.Sprintf("%s_BG_BRIGHT = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.BgBright))
	}

	sb.WriteString("\n# Controls\n")
	for _, c := range s.Controls {
		name := toUpperSnake(c.Name)
		if c.Seq != "" {
			sb.WriteString(fmt.Sprintf("%s = \"%s\"\n", name, c.Seq))
		} else {
			count := strings.Count(c.Fmt, "%d")
			if count == 1 {
				singleFmt := strings.Replace(c.Fmt, "%d", "1", 1)
				sb.WriteString(fmt.Sprintf("%s_1 = \"%s\"\n", name, singleFmt))
			}
		}
	}

	sb.WriteString("\n# Character widths, used by visible_width and truncate\n")
	sb.WriteString("_ZERO_WIDTH_RANGES = (\n")
	writeRanges(&sb, s.Widths.Zero, "    ", "(0x%04X, 0x%04X),", "  # ", false)
	sb.WriteString(")\n")
	sb.WriteString("_WIDE_RANGES = (\n")
	writeRanges(&sb, s.Widths.Wide, "    ", "(0x%04X, 0x%04X),", "  # ", false)
	sb.WriteString(")\n")

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

// Generates the exported Rust code
func generateRust(s Source, path string) error {
	fmt.Println("Generating for Rust...")
	var sb strings.Builder

	sb.WriteString("\n// Styles\n")
	for _, st := range s.Styles {
		sb.WriteString(fmt.Sprintf("pub const %s: &str = \"\\x1b[%sm\";\n", toUpperSnake(st.Name), st.Code))
	}

	sb.WriteString("\n// Colors\n")
	for _, c := range s.Colors {
		sb.WriteString(fmt.Sprintf("pub const %s_FG: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.Fg))
		sb.WriteString(fmt.Sprintf("pub const %s_BG: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.Bg))
		sb.WriteString(fmt.Sprintf("pub const %s_FG_BRIGHT: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.FgBright))
		sb.WriteString(fmt.Sprintf("pub const %s_BG_BRIGHT: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.BgBright))
	}

	sb.WriteString("\n// Controls\n")
	for _, c := range s.Controls {
		name := toUpperSnake(c.Name)
		if c.Seq != "" {
			seq := strings.ReplaceAll(c.Seq, "\\033", "\\x1b")
			sb.WriteString(fmt.Sprintf("pub const %s: &str = \"%s\";\n", name, seq))
		} else {
			count := strings.Count(c.Fmt, "%d")
			if count == 1 {
				singleFmt := strings.Replace(c.Fmt, "%d", "1", 1)
				singleFmt = strings.ReplaceAll(singleFmt, "\\033", "\\x1b")
				sb.WriteString(fmt.Sprintf("pub const %s_1: &str = \"%s\";\n", name, singleFmt))
			}
		}
	}

	sb.WriteString("\n// Character widths, used by visible_width and truncate\n")
	sb.WriteString("const ZERO_WIDTH_RANGES: &[(u32, u32)] = &[\n")
	writeRanges(&sb, s.Widths.Zero, "    ", "(0x%04X, 0x%04X),", " // ", true)
	sb.WriteString("];\n")
	sb.WriteString("const WIDE_RANGES: &[(u32, u32)] = &[\n")
	writeRanges(&sb, s.Widths.Wide, "    ", "(0x%04X, 0x%04X),", " // ", true)
	sb.WriteString("];\n")

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

// Generates the exported TS code
func generateTS(s Source, path string) error {
	fmt.Println("Generating for TS...")
	var sb strings.Builder
	indent := "    " // 4 spaces

	sb.WriteString("\n" + indent + "// Styles\n")
	for _, st := range s.Styles {
		sb.WriteString(fmt.Sprintf("%sstatic readonly %s = \"\\u001b[%sm\";\n", indent, st.Name, st.Code))
	}

	sb.WriteString("\n" + indent + "// Colors\n")
	for _, c := range s.Colors {
		sb.WriteString(fmt.Sprintf("%sstatic readonly %sFg = \"\\u001b[%sm\";\n", indent, c.Name, c.Fg))
		sb.WriteString(fmt.Sprintf("%sstatic readonly %sBg = \"\\u001b[%sm\";\n", indent, c.Name, c.Bg))
		sb.WriteString(fmt.Sprintf("%sstatic readonly %sFgBright = \"\\u001b[%sm\";\n", indent, c.Name, c.FgBright))
		sb.WriteString(fmt.Sprintf("%sstatic readonly %sBgBright = \"\\u001b[%sm\";\n", indent, c.Name, c.BgBright))
	}

	sb.WriteString("\n" + indent + "// Controls\n")
	for _, c := range s.Controls {
		name := toLowerCamel(c.Name)
		if c.Seq != "" {
			seq := strings.ReplaceAll(c.Seq, "\\033", "\\u001b")
			sb.WriteString(fmt.Sprintf("%sstatic readonly %s = \"%s\";\n", indent, name, seq))
		} else {
			count := strings.Count(c.Fmt, "%d")
			if count == 1 {
				singleFmt := strings.Replace(c.Fmt, "%d", "1", 1)
				singleFmt = strings.ReplaceAll(singleFmt, "\\033", "\\u001b")
				sb.WriteString(fmt.Sprintf("%sstatic readonly %s1 = \"%s\";\n", indent, name, singleFmt))
			}
		}
	}

	sb.WriteString("\n" + indent + "// Character widths, used by visibleWidth and truncate\n")
	sb.WriteString(indent + "private static readonly zeroWidthRanges: [number, number][] = [\n")
	writeRanges(&sb, s.Widths.Zero, indent+indent, "[0x%04X, 0x%04X],", " // ", false)
	sb.WriteString(indent + "];\n")
	sb.WriteString(indent + "private static readonly wideRanges: [number, number][] = [\n")
	writeRanges(&sb, s.Widths.Wide, indent+indent, "[0x%04X, 0x%04X],", " // ", false)
	sb.WriteString(indent + "];\n")

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

// Generates the exported Dart code
func generateDart(s Source, path string) error {
	fmt.Println("Generating for Dart...")
	var sb strings.Builder
	indent := "  " // 2 spaces

	sb.WriteString("\n" + indent + "// Styles\n")
	for _, st := range s.Styles {
		sb.WriteString(fmt.Sprintf("%sstatic const String %s = \"\\u001b[%sm\";\n", indent, toLowerCamel(st.Name), st.Code))
	}

	sb.WriteString("\n" + indent + "// Colors\n")
	for _, c := range s.Colors {
		sb.WriteString(fmt.Sprintf("%sstatic const String %sFg = \"\\u001b[%sm\";\n", indent, toLowerCamel(c.Name), c.Fg))
		sb.WriteString(fmt.Sprintf("%sstatic const String %sBg = \"\\u001b[%sm\";\n", indent, toLowerCamel(c.Name), c.Bg))
		sb.WriteString(fmt.Sprintf("%sstatic const String %sFgBright = \"\\u001b[%sm\";\n", indent, toLowerCamel(c.Name), c.FgBright))
		sb.WriteString(fmt.Sprintf("%sstatic const String %sBgBright = \"\\u001b[%sm\";\n", indent, toLowerCamel(c.Name), c.BgBright))
	}

	sb.WriteString("\n" + indent + "// Controls\n")
	for _, c := range s.Controls {
		name := toLowerCamel(c.Name)
		if c.Seq != "" {
			seq := strings.ReplaceAll(c.Seq, "\\033", "\\u001b")
			sb.WriteString(fmt.Sprintf("%sstatic const String %s = \"%s\";\n", indent, name, seq))
		} else {
			count := strings.Count(c.Fmt, "%d")
			if count == 1 {
				singleFmt := strings.Replace(c.Fmt, "%d", "1", 1)
				singleFmt = strings.ReplaceAll(singleFmt, "\\033", "\\u001b")
				sb.WriteString(fmt.Sprintf("%sstatic const String %s1 = \"%s\";\n", indent, name, singleFmt))
			}
		}
	}

	sb.WriteString("\n" + indent + "// Character widths, used by visibleWidth and truncate\n")
	sb.WriteString(indent + "static const List<(int, int)> _zeroWidthRanges = [\n")
	writeRanges(&sb, s.Widths.Zero, indent+indent, "(0x%04X, 0x%04X),", " // ", false)
	sb.WriteString(indent + "];\n")
	sb.WriteString(indent + "static const List<(int, int)> _wideRanges = [\n")
	writeRanges(&sb, s.Widths.Wide, indent+indent, "(0x%04X, 0x%04X),", " // ", false)
	sb.WriteString(indent + "];\n")

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}
