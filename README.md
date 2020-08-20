# ecnode_checker
ChubaoFS EcNode Checker

检查EcPartition所有Extent文件是否正常，对于不正常的文件，进行数据修复操作

| Master 测试环境地址 |
| :-----|
| test.chubaofs.jd.local | 
| 127.0.0.1:33845 |

## 使用方法
``` shell
Usage:
  ecnode_checker [command]

Available Commands:
  ecmonkey    deliberately damage a Extent file on EcNode
  help        Help about any command
  validate    validate EcExtent of EcPartition on EcNode

Flags:
  -h, --help   help for ecnode_checker

Use "ecnode_checker [command] --help" for more information about a command.

```


