# Unix Tools

A collection of classic Unix command-line utilities reimplemented in **Go**.  
This project is designed to deepen my understanding of Go (file I/O, string processing, concurrency, testing) while building a practical toolkit.

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)

Currently implemented:
- `wc`

Planned tools:
- `cut`
- `sort`
- `uniq`

---

## Table of Contents
- [Quick Start](#quick-start)
- [Implemented Tools](#implemented-tools)
- [Architecture](#architecture)
- [Testing](#testing)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## Quick Start

**Prerequisites:**
- [Go 1.25+](https://go.dev/dl/)

**Clone the repo:**
```bash
git clone https://github.com/Priyaa-18/unix-tools.git
cd unix-tools
```

**Run `wc`:**
```bash
echo -e "hello\nworld\nhi" > test.txt
go run ./cmd/wc test.txt
# Expected output:
# 3 test.txt
```

---

## Implemented Tools
**`wc` (word count)**  
Supports:
- `-l` → count lines  
- `-w` → count words  
- `-c` → count bytes  
- `-m` → count characters  

Default (no flags): shows all counts.

Usage:
```bash
echo -e "hello\nworld\nhi" > test.txt

# All counts
go run ./cmd/wc test.txt
# Output: 3 3 15 15 test.txt

# Lines only
go run ./cmd/wc -l test.txt
# Output: 3 test.txt
```

---

## Architecture
- **Monorepo layout:** each tool lives under `cmd/<toolname>`.
- **Go modules:** project initialized with `go.mod`.
- **Standard library first:** minimal external dependencies (later `cobra` for CLI).

```text
unix-tools/
  cmd/
    wc/
      main.go
  go.mod
  LICENSE
  README.md
```

---

## Testing
Unit tests will be added with Go’s built-in testing framework:

```bash
go test ./...
```

---

## Roadmap
- ✅ Initialize repo with `go.mod`
- ✅ Add `wc` (full: lines, words, bytes, characters)
- ⏳ Add `cut` (delimiter + field extraction)
- ⏳ Add `sort` (lexicographic, numeric, reverse)
- ⏳ Add `uniq` (deduplicate, count)
- ⏳ Add automated testing and GitHub Actions CI
- ⏳ Package as installable binaries

---

## Contributing
Contributions are welcome!  

1. Fork the repo
2. Create a feature branch
3. Add tests for new features
4. Ensure all tests pass
5. Submit a PR

---

## License

This project is licensed under the MIT License.

---
