package version

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// GoYaml represents the structure of go.yaml
type GoYaml struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

// BumpType represents the type of version bump
type BumpType string

const (
	BumpBuild BumpType = "build"
	BumpMinor BumpType = "minor"
	BumpMajor BumpType = "major"
)

// ParseBumpType converts a string to BumpType
func ParseBumpType(s string) (BumpType, error) {
	switch strings.ToLower(s) {
	case "build", "patch":
		return BumpBuild, nil
	case "minor":
		return BumpMinor, nil
	case "major":
		return BumpMajor, nil
	default:
		return "", fmt.Errorf("invalid bump type: %s (must be build, minor, or major)", s)
	}
}

// SemVer represents a semantic version
type SemVer struct {
	Major int
	Minor int
	Build int
}

// ParseSemVer parses a semver string
func ParseSemVer(s string) (SemVer, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return SemVer{}, fmt.Errorf("invalid semver format: %s", s)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	build, err := strconv.Atoi(parts[2])
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid build version: %s", parts[2])
	}

	return SemVer{Major: major, Minor: minor, Build: build}, nil
}

// String returns the semver as a string
func (v SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
}

// Bump increments the version based on bump type
func (v SemVer) Bump(bt BumpType) SemVer {
	switch bt {
	case BumpMajor:
		return SemVer{Major: v.Major + 1, Minor: 0, Build: 0}
	case BumpMinor:
		return SemVer{Major: v.Major, Minor: v.Minor + 1, Build: 0}
	case BumpBuild:
		return SemVer{Major: v.Major, Minor: v.Minor, Build: v.Build + 1}
	}
	return v
}

// ReadSourceVersion reads the version from go.yaml
func ReadSourceVersion() (string, error) {
	data, err := os.ReadFile("go.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to read go.yaml: %w", err)
	}

	var goYaml GoYaml
	if err := yaml.Unmarshal(data, &goYaml); err != nil {
		return "", fmt.Errorf("failed to parse go.yaml: %w", err)
	}

	return goYaml.Version, nil
}

// Bump reads the current version, bumps it, and updates all files
func Bump(bumpType BumpType) error {
	// Read current version from go.yaml
	currentVersion, err := ReadSourceVersion()
	if err != nil {
		return err
	}

	// Parse and bump version
	semver, err := ParseSemVer(currentVersion)
	if err != nil {
		return err
	}

	newVersion := semver.Bump(bumpType)
	fmt.Printf("Bumping version: %s -> %s\n", currentVersion, newVersion.String())

	// Update all files
	files := []struct {
		path    string
		updater func(string, string) error
	}{
		{"go.yaml", updateYamlVersion},
		{"package.json", updateJsonVersion},
		{"pubspec.yaml", updateYamlVersion},
		{"pyproject.toml", updateTomlVersion},
		{"Cargo.toml", updateTomlVersion},
	}

	for _, f := range files {
		fmt.Printf("Updating %s...\n", f.path)
		if err := f.updater(f.path, newVersion.String()); err != nil {
			return fmt.Errorf("failed to update %s: %w", f.path, err)
		}
	}

	fmt.Printf("Successfully bumped version to %s\n", newVersion.String())
	return nil
}

// updateYamlVersion updates version in YAML files (go.yaml, pubspec.yaml)
func updateYamlVersion(path string, newVersion string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Use regex to replace version line while preserving formatting
	re := regexp.MustCompile(`(?m)^version:\s*[\d.]+`)
	updated := re.ReplaceAllString(string(data), "version: "+newVersion)

	return os.WriteFile(path, []byte(updated), 0644)
}

// updateJsonVersion updates version in package.json
func updateJsonVersion(path string, newVersion string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Use regex to replace version in JSON
	re := regexp.MustCompile(`"version":\s*"[\d.]+"`)
	updated := re.ReplaceAllString(string(data), `"version": "`+newVersion+`"`)

	return os.WriteFile(path, []byte(updated), 0644)
}

// updateTomlVersion updates version in TOML files (pyproject.toml, Cargo.toml)
func updateTomlVersion(path string, newVersion string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Use regex to replace version in TOML
	re := regexp.MustCompile(`version\s*=\s*"[\d.]+"`)
	updated := re.ReplaceAllString(string(data), `version = "`+newVersion+`"`)

	return os.WriteFile(path, []byte(updated), 0644)
}
