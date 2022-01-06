# okex
```go
package main

import (
	"log"

	"github.com/iaping/exgo/okex"
)

func main() {
	/*ok := okex.NewClient(&okex.Config{
		APIKey:     "your APIKey",
		SecretKey:  "your SecretKey",
		Passphrase: "your Passphrase",
		Simulated:  true,
	})*/
	ok := okex.New("your APIKey", "your SecretKey", "your Passphrase", false)
	data, err := ok.AssetCurrencies() // api AssetCurrencies()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
```

# 资金
API | 描述
--- | ---
AssetCurrencies() | 获取币种列表
AssetBalances() | 获取资金账户余额
AssetBills() | 获取资金流水

# 账户
API | 描述
--- | ---
AccountBalance() | 查看账户余额
AccountBills() | 账单流水查询（近七天）
AccountConfig() | 查看账户配置

# 子账户
API | 描述
--- | ---
AccountSubaccountBalances() | 获取子账户资产余额