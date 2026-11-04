# urlencode and urldecode

Two Go commands for URL query/form encoding and decoding, using only the standard
library.

## Install with Homebrew

Tap this repository and install both commands:

```sh
brew tap argium/urlencode-decode https://github.com/argium/go-urlencode-urldecode
brew install urlencode-decode
```

The `urlencode-decode` package installs both `urlencode` and `urldecode` on your
PATH. It builds a versioned release from source, verified with a SHA-256 checksum.
Homebrew installs Go as a build dependency.

`argium/urlencode-decode` is the tap name; `urlencode-decode` is the package name.
Homebrew does not support `brew install argium/urlencode-decode`.

Source archives and checksums are also available on the
[releases page](https://github.com/argium/go-urlencode-urldecode/releases).

## Build

Alternatively, clone the repository and build locally:

```sh
git clone https://github.com/argium/go-urlencode-urldecode.git
cd go-urlencode-urldecode
go build -o bin/urlencode ./cmd/urlencode
go build -o bin/urldecode ./cmd/urldecode
export PATH="$PWD/bin:$PATH"
```

To install both commands into your Go binary directory:

```sh
go install ./cmd/...
```

## Usage

Pass one argument, or omit it to read stdin until EOF:

```sh
urlencode 'hello world&name=Go'
# hello+world%26name%3DGo

urldecode 'hello+world%26name%3DGo'
# hello world&name=Go

printf '%s' 'https://example.com/a?x=1&y=two words' | urlencode
# https%3A%2F%2Fexample.com%2Fa%3Fx%3D1%26y%3Dtwo+words

printf '%s' 'hello%20world' | urldecode
# hello world
```

- Encoding uses query/form rules: spaces become `+`, and literal `+` becomes `%2B`.
  This encodes a query value, not an entire URL while preserving its structure.
- Decoding accepts percent escapes and converts `+` to a space.
- An argument takes precedence over stdin, including an empty argument.
- All input bytes are preserved, including whitespace and trailing newlines.
  Use `printf` rather than `echo` when you do not want a newline in the input.
- Results go to stdout without an added newline, so piping through both commands
  reproduces the original input exactly.
- More than one argument, malformed percent escapes, and I/O failures produce an
  error on stderr and a non-zero exit status. Quote arguments containing spaces.
- Arguments are literal input; there are no flags.

## Publishing a release

Push a tag in the form `vMAJOR.MINOR.PATCH` to trigger the
[release workflow](.github/workflows/release.yml). For example:

```sh
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```

The workflow builds and checks both commands, then publishes a GitHub release
with a source archive and `checksums.txt`. After it succeeds, update the URL and
SHA-256 in [the Homebrew formula](Formula/urlencode-decode.rb) to that release and
push the change to `main`.
