package main

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func IsBeautifulAddress(address common.Address, long int, strict bool) bool {
	curr := 1
	count := 0
	str := address.Hex()
	if !strict {
		str = strings.ToLower(str)
	}

	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			curr++
		} else {
			if curr > count && curr >= 2 {
				count = curr
			}
			curr = 1
		}
	}
	return count >= long
}
