# Contributing

For installation and usage, see the [README](README.md).

## Build from source

Install the Go version specified in [go.mod](go.mod), then clone and build:

```sh
git clone https://github.com/argium/urlencode-urldecode.git
cd urlencode-urldecode
go build -o bin/urlencode ./cmd/urlencode
go build -o bin/urldecode ./cmd/urldecode
export PATH="$PWD/bin:$PATH"
```

To install both commands into your Go binary directory:

```sh
go install ./cmd/...
```

## Development checks

Format changed Go files with `gofmt`, then run:

```sh
go vet ./...
go test ./...
go build ./cmd/...
```

There are currently no Go unit tests. The
[release workflow](.github/workflows/release.yml) also checks exact output for
argument-based encoding and stdin-based decoding.

## Publishing a release

Commit and push the changes to release, then push a tag in the form
`vMAJOR.MINOR.PATCH` to trigger the
[release workflow](.github/workflows/release.yml). For example:

```sh
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```

The workflow builds and checks both commands, then publishes a GitHub release
with a source archive and `checksums.txt`. After it succeeds, update the URL and
SHA-256 in [the Homebrew formula](Formula/urlencode-decode.rb) to that release and
push the change to `main`.
