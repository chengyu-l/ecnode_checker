# ecnode_checker

**ecnode_checker** 目前支持3个功能：
1. 校验器：验证指定 **EcPartition** 所有 **Extent** 文件的数据是否正常。基于 EC 算法 [klauspost/reedsolomon](https://github.com/klauspost/reedsolomon) 提供的 *verify* 方法进行验证。
2. 修复器：修复指定 **EcPartition** 指定 **Extent** 文件（需要指定extentId参数）或者所有 **Extent** 文件。
3. 数据静默损坏模拟器：故意修改指定文件的数据内容。支持修改数据的offset以及修改数据的长度。

> **注意:**
> EcNode 目前只支持针对Extent文件丢失或者下线EcNode两种情况数据的修复。
> 不支持数据静默损坏等部分内容发生变化的数据修复。这是由于使用的EC库[klauspost/reedsolomon](https://github.com/klauspost/reedsolomon) 不支持这种情况。

## 编译
执行build.sh文件，可以编译生成 **ecnode_checker** 可执行文件，默认编译的是linux amd64平台的可执行文件。

## 使用方法

1. 在执行 **ecnode_checker** 前，必须先在ecnode_checker可执行文件同级目录下，创建配置文件 *ecnode_checker.json*。
配置文件用于配置ChubaoFS相关信息，例如Master地址，其中Master地址为必填项。

*ecnode_checker.json* 示例如下：
```json
{
  "masterAddr": [
    "127.0.0.1:34027"
  ]
}
```

| Master 测试环境地址 |
| :-----|
| test.chubaofs.jd.local | 


2. 准备好配置文件后，执行 *./ecnode_checker -h* 可以 **ecnode_checker** 的帮助信息以及具体参数。如下所示：

```shell
# ./ecnode_checker -h
Usage:
  ecnode_checker [command]

Available Commands:
  cm          
  ecmonkey    deliberately damage a Extent file on EcNode
  help        Help about any command
  repair      repair EcExtent of EcPartition on EcNode
  validate    validate EcExtent of EcPartition on EcNode

Flags:
  -h, --help   help for ecnode_checker

Use "ecnode_checker [command] --help" for more information about a command.

```

### 1. 校验器

```shell
# ./ecnode_checker validate
Error: required flag(s) "partitionId" not set
Usage:
  ecnode_checker validate [flags]

Flags:
  -h, --help                 help for validate
      --partitionId string   partitionId

checker error: required flag(s) "partitionId" not set
```

### 2. 修复器

```shell
# ./ecnode_checker repair  
Error: required flag(s) "partitionId" not set
Usage:
  ecnode_checker repair [flags]

Flags:
      --extentId string      If set extentId, it will only repair this extent, otherwise, repair all extents in this EcPartition
  -h, --help                 help for repair
      --partitionId string   partitionId

checker error: required flag(s) "partitionId" not set
```

### 3. 数据静默损坏模拟器

```shell
# ./ecnode_checker ecmonkey
Error: required flag(s) "file" not set
Usage:
  ecnode_checker ecmonkey [flags]

Flags:
      --file string     extent file path
  -h, --help            help for ecmonkey
      --offset string   damage the extent file from the offset (default "10")
      --size string     how many data size are damaged in the extent file (default "1")

checker error: required flag(s) "file" not set
```
