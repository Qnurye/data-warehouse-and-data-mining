[English](README.md) | 简体中文

# 关联规则挖掘算法 Go 实现

## FP-Growth


### 快速开始

```shell
go build -o build/fp ./cmd/fp/main.go
./build/fp -s 0.01 -p retail.dat
```

`fp` 命令行参数：
- `-h`: 显示帮助信息
- `-s`: 最小支持度比值，默认 `0.01`
- `-c`: 最小计数支持度
- `-p`: 数据集文件路径，默认 `retail.dat`

`-s` 优于 `-c`，如果两个都设置，只使用 `-s`。

## Apriori

Apriori 我实现得太 Fancy 了，反而很慢，不推荐使用，后续不优化，因为 FP-Growth 降维打击。

### 快速开始

```shell
go build -o build/apriori ./cmd/apriori/main.go
./build/apriori -s 0.01 -p retail.dat
```

`apriori` 命令行参数：
- `-h`: 显示帮助信息
- `-s`: 最小支持度比值，默认 `0.01`
- `-c`: 最小计数支持度
- `-p`: 数据集文件路径，默认 `retail.dat`

`-s` 优于 `-c`，如果两个都设置，只使用 `-s`。
