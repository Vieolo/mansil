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

func getArgs(count int) ([]string, []string) {
	if count == 1 {
		return []string{"n"}, []string{"n"}
	}
	var args []string
	var vars []string
	if count == 2 {
		return []string{"row", "col"}, []string{"row", "col"}
	}
	for i := 1; i <= count; i++ {
		v := fmt.Sprintf("n%d", i)
		args = append(args, v)
		vars = append(vars, v)
	}
	return args, vars
}

// Generators

func generateGo(s *Source, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "package mansil\n\nimport \"fmt\"\n\n")

	fmt.Fprintf(f, "// Styles\n")
	for _, st := range s.Styles {
		fmt.Fprintf(f, "const %s = \"\\033[%sm\"\n", st.Name, st.Code)
	}

	fmt.Fprintf(f, "\n// Colors\n")
	for _, c := range s.Colors {
		fmt.Fprintf(f, "const %sFG = \"\\033[%sm\"\n", c.Name, c.Fg)
		fmt.Fprintf(f, "const %sBG = \"\\033[%sm\"\n", c.Name, c.Bg)
		fmt.Fprintf(f, "const %sFGBright = \"\\033[%sm\"\n", c.Name, c.FgBright)
		fmt.Fprintf(f, "const %sBGBright = \"\\033[%sm\"\n", c.Name, c.BgBright)
	}

	fmt.Fprintf(f, "\n// Controls\n")
	for _, c := range s.Controls {
		if c.Seq != "" {
			fmt.Fprintf(f, "const %s = \"%s\"\n", c.Name, c.Seq)
		} else {
			count := strings.Count(c.Fmt, "%d")
			args, vars := getArgs(count)
			argDef := strings.Join(args, ", ") + " int"
			varList := strings.Join(vars, ", ")
			fmt.Fprintf(f, "func %s(%s) string {\n\treturn fmt.Sprintf(\"%s\", %s)\n}\n", c.Name, argDef, c.Fmt, varList)
		}
	}
	return nil
}

func generatePython(s *Source, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "\"\"\"Mansil ANSI Codes\"\"\"\n\n")

	fmt.Fprintf(f, "# Styles\n")
	for _, st := range s.Styles {
		fmt.Fprintf(f, "%s = \"\\033[%sm\"\n", toUpperSnake(st.Name), st.Code)
	}

	fmt.Fprintf(f, "\n# Colors\n")
	for _, c := range s.Colors {
		fmt.Fprintf(f, "%s_FG = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.Fg)
		fmt.Fprintf(f, "%s_BG = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.Bg)
		fmt.Fprintf(f, "%s_FG_BRIGHT = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.FgBright)
		fmt.Fprintf(f, "%s_BG_BRIGHT = \"\\033[%sm\"\n", toUpperSnake(c.Name), c.BgBright)
	}

	fmt.Fprintf(f, "\n# Controls\n")
	for _, c := range s.Controls {
		name := toUpperSnake(c.Name)
		if c.Seq != "" {
			fmt.Fprintf(f, "%s = \"%s\"\n", name, c.Seq)
		} else {
			count := strings.Count(c.Fmt, "%d")
			args, _ := getArgs(count)
			argDef := strings.Join(args, ": int, ") + ": int"

			pyFmt := c.Fmt
			for _, a := range args {
				pyFmt = strings.Replace(pyFmt, "%d", "{"+a+"}", 1)
			}

			funcName := strings.ToLower(name)
			fmt.Fprintf(f, "def %s(%s) -> str:\n    return f\"%s\"\n", funcName, argDef, pyFmt)
		}
	}

	return nil
}

func generateRust(s *Source, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "//! Mansil ANSI Codes\n\n")

	fmt.Fprintf(f, "// Styles\n")
	for _, st := range s.Styles {
		fmt.Fprintf(f, "pub const %s: &str = \"\\x1b[%sm\";\n", toUpperSnake(st.Name), st.Code)
	}

	fmt.Fprintf(f, "\n// Colors\n")
	for _, c := range s.Colors {
		fmt.Fprintf(f, "pub const %s_FG: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.Fg)
		fmt.Fprintf(f, "pub const %s_BG: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.Bg)
		fmt.Fprintf(f, "pub const %s_FG_BRIGHT: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.FgBright)
		fmt.Fprintf(f, "pub const %s_BG_BRIGHT: &str = \"\\x1b[%sm\";\n", toUpperSnake(c.Name), c.BgBright)
	}

	fmt.Fprintf(f, "\n// Controls\n")
	for _, c := range s.Controls {
		name := toUpperSnake(c.Name)
		if c.Seq != "" {
			fmt.Fprintf(f, "pub const %s: &str = \"%s\";\n", name, strings.ReplaceAll(c.Seq, "\\033", "\\x1b"))
		} else {
			count := strings.Count(c.Fmt, "%d")
			args, vars := getArgs(count)

			rustArgs := ""
			for _, a := range args {
				rustArgs += fmt.Sprintf("%s: u32, ", a)
			}
			rustArgs = strings.TrimSuffix(rustArgs, ", ")

			rustFmt := strings.Replace(c.Fmt, "%d", "{}", -1)
			rustFmt = strings.ReplaceAll(rustFmt, "\\033", "\\x1b")
			funcName := strings.ToLower(name)
			fmt.Fprintf(f, "pub fn %s(%s) -> String {\n    format!(\"%s\", %s)\n}\n", funcName, rustArgs, rustFmt, strings.Join(vars, ", "))
		}
	}
	return nil
}

func generateNPM(s *Source, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "// Styles\n")
	for _, st := range s.Styles {
		fmt.Fprintf(f, "export const %s = \"\\u001b[%sm\";\n", st.Name, st.Code)
	}

	fmt.Fprintf(f, "\n// Colors\n")
	for _, c := range s.Colors {
		fmt.Fprintf(f, "export const %sFg = \"\\u001b[%sm\";\n", c.Name, c.Fg)
		fmt.Fprintf(f, "export const %sBg = \"\\u001b[%sm\";\n", c.Name, c.Bg)
		fmt.Fprintf(f, "export const %sFgBright = \"\\u001b[%sm\";\n", c.Name, c.FgBright)
		fmt.Fprintf(f, "export const %sBgBright = \"\\u001b[%sm\";\n", c.Name, c.BgBright)
	}

	fmt.Fprintf(f, "\n// Controls\n")
	for _, c := range s.Controls {
		name := toLowerCamel(c.Name)
		if c.Seq != "" {
			fmt.Fprintf(f, "export const %s = \"%s\";\n", name, strings.ReplaceAll(c.Seq, "\\033", "\\u001b"))
		} else {
			count := strings.Count(c.Fmt, "%d")
			args, _ := getArgs(count)
			argDef := strings.Join(args, ": number, ") + ": number"

			jsFmt := strings.ReplaceAll(c.Fmt, "\\033", "\\u001b")
			for _, a := range args {
				jsFmt = strings.Replace(jsFmt, "%d", "${"+a+"}", 1)
			}

			fmt.Fprintf(f, "export function %s(%s): string {\n    return `%s`;\n}\n", name, argDef, jsFmt)
		}
	}
	return nil
}

func generateDart(s *Source, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "// Styles\n")
	for _, st := range s.Styles {
		fmt.Fprintf(f, "const String %s = \"\\u001b[%sm\";\n", toLowerCamel(st.Name), st.Code)
	}

	fmt.Fprintf(f, "\n// Colors\n")
	for _, c := range s.Colors {
		fmt.Fprintf(f, "const String %sFg = \"\\u001b[%sm\";\n", toLowerCamel(c.Name), c.Fg)
		fmt.Fprintf(f, "const String %sBg = \"\\u001b[%sm\";\n", toLowerCamel(c.Name), c.Bg)
		fmt.Fprintf(f, "const String %sFgBright = \"\\u001b[%sm\";\n", toLowerCamel(c.Name), c.FgBright)
		fmt.Fprintf(f, "const String %sBgBright = \"\\u001b[%sm\";\n", toLowerCamel(c.Name), c.BgBright)
	}

	fmt.Fprintf(f, "\n// Controls\n")
	for _, c := range s.Controls {
		name := toLowerCamel(c.Name)
		if c.Seq != "" {
			fmt.Fprintf(f, "const String %s = \"%s\";\n", name, strings.ReplaceAll(c.Seq, "\\033", "\\u001b"))
		} else {
			count := strings.Count(c.Fmt, "%d")
			args, _ := getArgs(count)
			argDef := strings.Join(args, ", int ")
			argDef = "int " + argDef

			dartStr := strings.ReplaceAll(c.Fmt, "\\033", "\\u001b")
			for _, a := range args {
				dartStr = strings.Replace(dartStr, "%d", "${"+a+"}", 1)
			}

			fmt.Fprintf(f, "String %s(%s) => \"%s\";\n", name, argDef, dartStr)
		}
	}
	return nil
}
