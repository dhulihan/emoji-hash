# emoji-hash

Generate a deterministic emoji based on a string.

It goes without saying, but DO NOT USE THIS FOR CRYPTOGRAPHY. This is a hobby project for inane uses.

## Installation

```sh
go install github.com/dhulihan/emoji-hash@latest
```

## CLI

```sh
$ emoji-hash "hello world"
🆓

$ emoji-hash "hello world"
🆓

$ go run main.go "hello jupiter"
🍠
```
