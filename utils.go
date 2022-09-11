package main

import (
	"strings"
)

func IsBeautifulAddress(address string, length int, strict bool, maxOffset int) bool {
	max := 1
	count := 1
	if !strict {
		address = strings.ToLower(address)
	}
	if strings.HasPrefix(address, "0x") {
		address = address[2:]
	}

	for i := 0; i < len(address)-1; i++ {
		if address[i] == address[i+1] {
			count++
			continue
		}

		if (i+1)-count-maxOffset > 0 {
			return false
		} else if count-length >= 0 {
			return true
		}

		if count > max && count >= 2 {
			max = count
		}
		count = 1
	}
	return max >= length
}
