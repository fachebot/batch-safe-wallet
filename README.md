# batch-safe-wallet
批量生成以太坊多签合约地址

## 编译程序
```bash
go build
```

## 批量生成地址
使用 `batch` 命令批量生成地址，参数如下：
```
batch-safe-wallet batch --help

批量生成地址

Usage:
batch [flags] [count] [length] [maxOffset]

Args:
count      int    生成地址数量 (default: 1000)
length     int    连续字符长度 (default: 5)
maxOffset  int    最大起始位置偏移 (default: 1)
```

示例：
```bash
./batch-safe-wallet batch 100000 6 0
```
程序会在后台自动生成以太坊账户，并将数据保存到 `keys.db` 文件，在全部生成后程序退出。

## 筛选靓号地址
使用 `filter` 命令从 `keys.db` 文件搜索符合条件的以太坊账户，参数如下：
```
batch-safe-wallet filter --help

搜索靓号地址

Usage:
  filter [flags] [type] [length] [maxOffset]

Args:
  type       string    地址类型(address/contract) (default: address)
  length     int       连续字符长度 (default: 5)
  maxOffset  int       最大起始位置偏移 (default: 1)

Flags:
  -h, --help     display help
```

示例：
```bash
./batch-safe-wallet filter contract 8 0
+--------------------------------------------+--------------------------------------------+--------------------------------------------------------------------+
|                  账户地址                  |                  合约地址                  |                              账户私钥                              |
+--------------------------------------------+--------------------------------------------+--------------------------------------------------------------------+
| 0x43a40387d42b2Dd939EC756764E63b82c3009301 | 0x777777772E10bCF4Ed7C41d77E03ce534911Cd1A | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0xBe5f6522E88c3e709D6bfD9E61a9c17bF7733A2a | 0x2222222267c81550565B3C4bf1ee8bc9Ca8A81b5 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0xdF35992e1fd3675C1cdffb697F6d1E15cfe8D5c4 | 0xcccccCCC20F1Eaa77a60c7A505683B6BD6C331ad | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0xa323Ac4006a331622926Ce51f0C67C81524a8A7c | 0x55555555AE39A9aa79f3Dc0E8161AD0ca29E2e74 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0x0678de3De7072A6A683532d29e8090FEED82e8d4 | 0x44444444b427c7B3c7731b41FF4Bb4f7C53a3752 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0xdFf553a07D55E47859AF7ba8C8E34F44535346A1 | 0xcCCCcCccbc3E326689052D6330e9959908BAd312 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0xeAEA0b762746fde15ae49C5ecF0C6B97DE0AE9A0 | 0x1111111124Bd583307E7218DAdcF5fdEa9247990 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0x2696B6dcb7209d3b1904C7e0C90DCf401DE1E56e | 0xeEeeeeEE04F7f3bDBb18eC5d0b269489D3CD0b89 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0x459C62a8a898d42F6d0485fb90E3B7cC2B80DFF6 | 0x222222227DDFbCEbdDb21935eACB896cd58C46a6 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0x7351bf6D92aD496055CD870A6987A6A2A9Ad6bc7 | 0x999999991C928D8159D2FB96aB421C27C748Fdf9 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0xfCe75B76D05974914E9F8e40636328Bb81815709 | 0x333333338f322f40E732E46eE1B9CAACD9F34424 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| 0x68180Db96696678F7d06aE8cf37625A47190B5d1 | 0xCCCcCCcC73005CFCc2E45986f70a201d0A364b63 | 0x0000000000000000000000000000000000000000000000000000000000000000 |
+--------------------------------------------+--------------------------------------------+--------------------------------------------------------------------+
```