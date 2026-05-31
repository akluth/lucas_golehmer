# Lucas-Lehmer

Lucas-Lehmer implementations for checking whether `2^p - 1` is a Mersenne prime.

The Go code builds a small command-line tool. The Rust code is packaged as a Cargo library.

## Go

```sh
go test ./...
go run . 19
```

Example output:

```text
2^19 - 1 is a Mersenne prime
```

## Rust

```sh
cargo test
```

## Notes

The input is the exponent `p`, not the Mersenne number itself. For example, `19`
checks whether `2^19 - 1` is prime.
