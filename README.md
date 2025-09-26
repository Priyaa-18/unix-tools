# Unix Tools

A collection of classic Unix command-line utilities reimplemented in **Go**.  
This project is designed to deepen my understanding of Go (file I/O, string processing, concurrency, testing) while building a practical toolkit.

Currently implemented:
- `wc` (line count version)

Planned tools:
- `wc` (full: lines, words, bytes, characters)
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
Counts the number of lines in a file.  
*Current version:* counts lines only.  
*Next milestone:* add word, byte, and character counts.

Usage:
```bash
go run ./cmd/wc <filename>
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
- Initialize repo with `go.mod`
- Add `wc` (line count)
- Extend `wc` to word, byte, char counts
- Add `cut` (delimiter + field extraction)
- Add `sort` (lexicographic, numeric, reverse)
- Add `uniq` (deduplicate, count)
- Add automated testing and GitHub Actions CI
- Package as installable binaries

---

## Contributing
Contributions are welcome!  
For new tools or feature requests, open an issue or submit a pull request.

---

## License

This project is licensed under the MIT License.

---