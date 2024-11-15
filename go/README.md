English | [简体中文](README.zh_hans.md)

# Association Rule Mining Algorithms in Go

## FP-Growth

### Quick Start

```shell
go build -o build/fp ./cmd/fp/main.go
./build/fp -s 0.01 -p retail.dat
```

`fp` command-line arguments:
- `-h`: Displays help information
- `-s`: Minimum support ratio, default is `0.01`
- `-c`: Minimum count for support
- `-p`: Path to the dataset file, default is `retail.dat`

`-s` takes precedence over `-c`. If both are set, only `-s` is used.

## Apriori

The Apriori implementation is too fancy and consequently quite slow. Not recommended for use, and further optimization is not planned since FP-Growth performs dimensionality reduction effectively.

### Quick Start

```shell
go build -o build/apriori ./cmd/apriori/main.go
./build/apriori -s 0.01 -p retail.dat
```

`apriori` command-line arguments:
- `-h`: Displays help information
- `-s`: Minimum support ratio, default is `0.01`
- `-c`: Minimum count for support
- `-p`: Path to the dataset file, default is `retail.dat`

`-s` takes precedence over `-c`. If both are set, only `-s` is used.
