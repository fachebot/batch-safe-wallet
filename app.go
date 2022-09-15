package main

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/olekukonko/tablewriter"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var app = grumble.New(&grumble.Config{
	Name:        "safe",
	Description: "批量创建safe合约地址",

	Flags: func(f *grumble.Flags) {
		f.Bool("v", "verbose", false, "enable verbose mode")
	},
})

func init() {
	app.AddCommand(&grumble.Command{
		Name: "deploy",
		Help: "部署多签合约",
		Args: func(a *grumble.Args) {
			a.String("path", "数据库路径", grumble.Default("keys.db"))
		},
		Run: BatchCreateAccounts,
	})

	app.AddCommand(&grumble.Command{
		Name: "batch",
		Help: "批量生成地址",
		Args: func(a *grumble.Args) {
			a.Int("count", "生成地址数量", grumble.Default(1000))
			a.Int("length", "连续字符长度", grumble.Default(5))
			a.Int("maxOffset", "最大起始位置偏移", grumble.Default(1))
		},
		Run: BatchCreateAccounts,
	})

	app.AddCommand(&grumble.Command{
		Name: "filter",
		Help: "搜索靓号地址",
		Args: func(a *grumble.Args) {
			a.String("type", "地址类型(address/contract)", grumble.Default("address"))
			a.Int("length", "连续字符长度", grumble.Default(5))
			a.Int("maxOffset", "最大起始位置偏移", grumble.Default(1))
		},
		Run: FilterBeautifulAddresses,
	})

	app.AddCommand(&grumble.Command{
		Name: "load",
		Help: "加载地址数据库",
		Args: func(a *grumble.Args) {
			a.String("path", "数据库路径", grumble.Default("keys.db"))
		},
		Run: LoadKeysDatabase,
	})
}

// BatchCreateAccounts 批量创建账户
func BatchCreateAccounts(c *grumble.Context) error {
	count := c.Args.Int("count")
	if count <= 0 {
		count = 1000
	}
	length := c.Args.Int("length")
	if length <= 0 {
		length = 5
	}
	maxOffset := c.Args.Int("maxOffset")
	if maxOffset < 0 {
		maxOffset = 1
	}

	// 生成地址
	var stopped int32
	keysChan := make(chan Key, BatchSize)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for atomic.LoadInt32(&stopped) == 0 {
				key, err := NewKey()
				if err != nil {
					panic(err)
				}

				if !IsBeautifulAddress(key.Address, length, false, maxOffset) &&
					!IsBeautifulAddress(key.Contract, length, false, maxOffset) {
					continue
				}

				keysChan <- key
			}
		}()
	}

	// 保存账户
	start := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		var size int
		keys := make([]Key, 0)
		for key := range keysChan {
			size++
			keys = append(keys, key)
			if len(keys) < BatchSize && size < count {
				continue
			}

			err := Keys{}.Save(keys)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d accounts created, time: %v\n", size, time.Since(start))
			keys = keys[:0]

			if size >= count {
				atomic.StoreInt32(&stopped, 1)
				waitGroup.Done()
				return
			}
		}
	}()

	waitGroup.Wait()

	return nil
}

// FilterBeautifulAddresses 筛选靓号地址
func FilterBeautifulAddresses(c *grumble.Context) error {
	length := c.Args.Int("length")
	if length <= 0 {
		length = 5
	}
	maxOffset := c.Args.Int("maxOffset")
	if maxOffset < 0 {
		maxOffset = 1
	}
	addressType := c.Args.String("type")
	if addressType != "contract" {
		addressType = "address"
	}

	offset := 0
	const limit = BatchSize
	table := tablewriter.NewWriter(os.Stdout)
	for {
		keys, err := Keys{}.Scan(offset, limit)
		if err != nil {
			return err
		}

		if addressType == "address" {
			for _, key := range keys {
				if IsBeautifulAddress(key.Address, length, false, maxOffset) {
					table.Append([]string{key.Address, key.Contract, key.PrivateKey})
				}
			}
		} else {
			for _, key := range keys {
				if IsBeautifulAddress(key.Contract, length, false, maxOffset) {
					table.Append([]string{key.Address, key.Contract, key.PrivateKey})
				}
			}
		}

		if len(keys) < limit {
			break
		}
		offset += len(keys)
	}

	table.SetHeader([]string{"账户地址", "合约地址", "账户私钥"})
	table.Render()

	return nil
}

// LoadKeysDatabase 加载地址数据库
func LoadKeysDatabase(c *grumble.Context) error {
	openDatabase(c.Args.String("path"))
	return nil
}
