package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"gorm.io/gorm"

	"github.com/desertbit/grumble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/olekukonko/tablewriter"
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
		Name: "batch",
		Help: "批量生成地址",
		Args: func(a *grumble.Args) {
			a.Int("count", "生成地址数量", grumble.Default(1000))
			a.Int("length", "连续字符长度", grumble.Default(5))
			a.Int("maxOffset", "最大起始位置偏移", grumble.Default(1))
		},
		Flags: func(f *grumble.Flags) {
			f.IntL("batchSize", 100, "批次大小")
			f.StringL("type", "evm", "地址类型(evm/tron)")
		},
		Run: BatchCreateAccounts,
	})

	app.AddCommand(&grumble.Command{
		Name: "create2",
		Help: "使用Create2生成地址",
		Args: func(a *grumble.Args) {
			a.String("deployer", "部署器地址", grumble.Default(""))
			a.String("initHash", "部署合约代码哈希值", grumble.Default(""))
			a.Uint64("saltNonce", "起始nonce位置", grumble.Default(uint64(0)))
		},
		Flags: func(f *grumble.Flags) {
			f.IntL("batchSize", 100, "批次大小")
			f.IntL("count", 1000, "生成地址数量")
			f.IntL("length", 5, "连续字符长度")
			f.IntL("maxOffset", 1, "最大起始位置偏移")
		},
		Run: BatchCreate2Accounts,
	})

	app.AddCommand(&grumble.Command{
		Name: "filter",
		Help: "搜索靓号地址",
		Args: func(a *grumble.Args) {
			a.String("type", "地址类型(address/contract)", grumble.Default("address"))
			a.Int("length", "连续字符长度", grumble.Default(5))
			a.Int("maxOffset", "最大起始位置偏移", grumble.Default(1))
		},
		Flags: func(f *grumble.Flags) {
			f.BoolL("create2", false, "create2方式创建的地址")
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
	keyType := c.Flags.String("type")
	batchSize := c.Flags.Int("batchSize")

	// 生成地址
	var stopped int32
	keysChan := make(chan Key, batchSize)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for atomic.LoadInt32(&stopped) == 0 {
				key, err := NewKey(keyType)
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
			if len(keys) < batchSize && size < count {
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

// BatchCreate2Accounts 批量创建账户，使用create2方式
func BatchCreate2Accounts(c *grumble.Context) error {
	initHash := common.HexToHash(c.Args.String("initHash"))
	deployer := common.HexToAddress(c.Args.String("deployer"))
	saltNonce := c.Args.Uint64("saltNonce")

	count := c.Flags.Int("count")
	if count <= 0 {
		count = 1000
	}
	length := c.Flags.Int("length")
	if length <= 0 {
		length = 5
	}
	maxOffset := c.Flags.Int("maxOffset")
	if maxOffset < 0 {
		maxOffset = 1
	}
	batchSize := c.Flags.Int("batchSize")

	if saltNonce == 0 {
		var err error
		n, err := Create2Keys{}.LastNonce(deployer)
		if err == nil {
			saltNonce = n + 1
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// 生成地址
	var stopped int32
	keysChan := make(chan Create2Key, batchSize)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for atomic.LoadInt32(&stopped) == 0 {
				key, err := NewCreate2Key(deployer, initHash, atomic.AddUint64(&saltNonce, 1))
				if err != nil {
					panic(err)
				}

				if !IsBeautifulAddress(key.Address, length, true, maxOffset) &&
					!IsBeautifulAddress(key.Contract, length, true, maxOffset) {
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
		keys := make([]Create2Key, 0)
		for key := range keysChan {
			size++
			keys = append(keys, key)
			if len(keys) < batchSize && size < count {
				continue
			}

			err := Create2Keys{}.Save(keys)
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

	if !c.Flags.Bool("create2") {
		return renderAddresses(addressType, length, maxOffset)
	}

	return renderCreate2Addresses(addressType, length, maxOffset)
}

// LoadKeysDatabase 加载地址数据库
func LoadKeysDatabase(c *grumble.Context) error {
	openDatabase(c.Args.String("path"))
	return nil
}

func renderAddresses(addressType string, length, maxOffset int) error {
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
				if IsBeautifulAddress(key.Address, length, true, maxOffset) {
					table.Append([]string{key.Address, key.Contract, key.PrivateKey})
				}
			}
		} else {
			for _, key := range keys {
				if IsBeautifulAddress(key.Contract, length, true, maxOffset) {
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

func renderCreate2Addresses(addressType string, length, maxOffset int) error {
	offset := 0
	const limit = BatchSize
	table := tablewriter.NewWriter(os.Stdout)
	for {
		keys, err := Create2Keys{}.Scan(offset, limit)
		if err != nil {
			return err
		}

		if addressType == "address" {
			for _, key := range keys {
				if IsBeautifulAddress(key.Address, length, true, maxOffset) {
					table.Append([]string{key.Address, key.Contract, strconv.FormatUint(key.SaltNonce, 10), key.InitHash})
				}
			}
		} else {
			for _, key := range keys {
				if IsBeautifulAddress(key.Contract, length, true, maxOffset) {
					table.Append([]string{key.Address, key.Contract, strconv.FormatUint(key.SaltNonce, 10), key.InitHash})
				}
			}
		}

		if len(keys) < limit {
			break
		}
		offset += len(keys)
	}

	table.SetHeader([]string{"账户地址", "合约地址", "Salt Nonce", "InitHash"})
	table.Render()
	return nil
}
