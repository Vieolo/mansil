package version

import (
	"os"
	"path/filepath"
	"testing"
)

// Writes content to a file in a temporary directory, runs update on it and returns the result
func runUpdater(t *testing.T, name, content string, update func(string, string) error) (string, error) {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	if err := update(path, "0.2.0"); err != nil {
		return "", err
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(got), nil
}

func TestUpdateTomlVersion(t *testing.T) {
	cases := []struct {
		name  string
		table string
		in    string
		want  string
	}{
		{
			name:  "Cargo.toml",
			table: "package",
			in: `[package]
name = "mansil"
version = "0.1.2"
rust-version = "1.70"

[dependencies]
foo = { version = "1.0.0" }

[dependencies.bar]
version = "0.1.2"
`,
			want: `[package]
name = "mansil"
version = "0.2.0"
rust-version = "1.70"

[dependencies]
foo = { version = "1.0.0" }

[dependencies.bar]
version = "0.1.2"
`,
		},
		{
			name:  "pyproject.toml",
			table: "project",
			in: `[project]
name = "mansil"
version = "0.1.2"

[tool.example]
version = "0.1.2"
`,
			want: `[project]
name = "mansil"
version = "0.2.0"

[tool.example]
version = "0.1.2"
`,
		},
	}
	for _, c := range cases {
		got, err := runUpdater(t, c.name, c.in, updateTomlVersion(c.table))
		if err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		if got != c.want {
			t.Errorf("%s: got:\n%s\nwant:\n%s", c.name, got, c.want)
		}
	}
}

func TestUpdateTomlVersionErrors(t *testing.T) {
	cases := map[string]string{
		"no table":   "[dependencies.bar]\nversion = \"0.1.2\"\n",
		"no version": "[package]\nname = \"mansil\"\n\n[dependencies.bar]\nversion = \"0.1.2\"\n",
	}
	for name, content := range cases {
		if _, err := runUpdater(t, "Cargo.toml", content, updateTomlVersion("package")); err == nil {
			t.Errorf("%s: expected an error", name)
		}
	}
}

func TestUpdateCargoLockVersion(t *testing.T) {
	lock := `version = 4

[[package]]
name = "mansil"
version = "0.1.2"

[[package]]
name = "other"
version = "0.1.2"
`
	want := `version = 4

[[package]]
name = "mansil"
version = "0.2.0"

[[package]]
name = "other"
version = "0.1.2"
`
	got, err := runUpdater(t, "Cargo.lock", lock, updateCargoLockVersion)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestUpdateCargoLockVersionWithoutMansil(t *testing.T) {
	if _, err := runUpdater(t, "Cargo.lock", "version = 4\n", updateCargoLockVersion); err == nil {
		t.Error("expected an error for a lock file without the mansil package")
	}
}
