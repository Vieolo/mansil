"""Regenerates the codes from ansi_codes.json, then runs the checks and tests of every language.

Usage:
    uv run --no-project agent.py              # every step
    uv run --no-project agent.py go dart      # only the steps of the given groups

Groups: generate, go, python, rust, ts, dart
"""

import subprocess
import sys
import time
from pathlib import Path

ROOT = Path(__file__).resolve().parent

# (group, name, command, whether any output means failure)
STEPS = [
    ("generate", "Generate the codes", ["go", "run", "./cmd/main.go", "generate"], False),
    ("go", "gofmt", ["gofmt", "-l", "mansil.go", "mansil_test.go", "cmd"], True),
    ("go", "go vet", ["go", "vet", ".", "./cmd/..."], False),
    ("go", "go test", ["go", "test", "-count=1", ".", "./cmd/..."], False),
    ("python", "unittest", [sys.executable, "-m", "unittest", "discover", "-s", "test"], False),
    ("rust", "cargo fmt", ["cargo", "fmt", "--check"], False),
    ("rust", "cargo test", ["cargo", "test", "--quiet"], False),
    ("ts", "tsc", ["npx", "--no-install", "tsc", "--noEmit"], False),
    ("ts", "node test", ["node", "--test", "test/mansil.test.ts"], False),
    ("dart", "dart format", ["dart", "format", "--output=none", "--set-exit-if-changed", "lib", "test"], False),
    ("dart", "dart analyze", ["dart", "analyze", "lib", "test"], False),
    ("dart", "dart test", ["dart", "test"], False),
]


def run(name, command, fail_on_output):
    start = time.monotonic()
    try:
        result = subprocess.run(command, cwd=ROOT, capture_output=True, text=True)
        output = (result.stdout + result.stderr).strip()
        passed = result.returncode == 0 and not (fail_on_output and output)
    except FileNotFoundError:
        output = f"{command[0]} is not installed"
        passed = False
    print(f"{'PASS' if passed else 'FAIL'}  {name} ({time.monotonic() - start:.1f}s)")
    if not passed:
        print("      " + output.replace("\n", "\n      "))
    return passed


def main():
    groups = set(sys.argv[1:])
    unknown = groups - {step[0] for step in STEPS}
    if unknown:
        sys.exit(f"Unknown groups: {', '.join(sorted(unknown))}")

    failed = [
        name
        for group, name, command, fail_on_output in STEPS
        if (not groups or group in groups) and not run(name, command, fail_on_output)
    ]
    if failed:
        sys.exit(f"\n{len(failed)} failed: {', '.join(failed)}")
    print("\nAll passed")


if __name__ == "__main__":
    main()
