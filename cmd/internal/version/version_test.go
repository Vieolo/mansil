package version

import (
	"os"
	"path/filepath"
	"testing"
)

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
	path := filepath.Join(t.TempDir(), "Cargo.lock")
	if err := os.WriteFile(path, []byte(lock), 0644); err != nil {
		t.Fatal(err)
	}

	if err := updateCargoLockVersion(path, "0.2.0"); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestUpdateCargoLockVersionWithoutMansil(t *testing.T) {
	path := filepath.Join(t.TempDir(), "Cargo.lock")
	if err := os.WriteFile(path, []byte("version = 4\n"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := updateCargoLockVersion(path, "0.2.0"); err == nil {
		t.Error("expected an error for a lock file without the mansil package")
	}
}
