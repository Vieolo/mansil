package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func ReadSource(path string) (*Source, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Source
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func Generate(root string) error {
	srcPath := filepath.Join(root, "ansi_codes.json")
	src, err := ReadSource(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read source: %w", err)
	}

	if err := generateGo(src, filepath.Join(root, "mansil.go")); err != nil {
		return err
	}
	if err := generatePython(src, filepath.Join(root, "mansil", "__init__.py")); err != nil {
		return err
	}
	if err := generateRust(src, filepath.Join(root, "src", "lib.rs")); err != nil {
		return err
	}
	if err := generateNPM(src, filepath.Join(root, "npm", "index.ts")); err != nil {
		return err
	}
	if err := generateDart(src, filepath.Join(root, "lib", "mansil.dart")); err != nil {
		return err
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

// replaceGenSection reads the file at path, finds the GEN START and GEN END markers,
// replaces the content between them with newContent, and writes the result back.
func replaceGenSection(path string, newContent string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	content := string(data)

	// Find markers - they might have different comment styles and indentation
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

func generateGo(s *Source, path string) error {
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

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

func generatePython(s *Source, path string) error {
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

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

func generateRust(s *Source, path string) error {
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

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

func generateNPM(s *Source, path string) error {
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

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}

func generateDart(s *Source, path string) error {
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

	sb.WriteString("\n")
	return replaceGenSection(path, sb.String())
}
