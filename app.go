package main

import (
	"fmt"
	"github.com/desertbit/grumble"
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
		Name: "batch",
		Help: "批量生成账户",
		Args: func(a *grumble.Args) {
			a.Int("count", "账户数量", grumble.Default(1000))
		},
		Run: BatchCreateAccounts,
	})
}

// BatchCreateAccounts 批量创建账户
func BatchCreateAccounts(c *grumble.Context) error {
	count := c.Args.Int("count")
	if count == 0 {
		count = 1000
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
