/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// Copyright 2013 Alexander Tchitchigin

package main

import (
	"fmt"
	"math"
)


func palindrome(n uint64) bool {
	if n / 10 == 0 {
		return true
	}

	digits := make([]byte, 0, 10)
	for i := n; i > 0; i = i / 10 {
		digits = append(digits, byte(i%10))
	}

	m := len(digits)
	for i := 0; i < m/2; i++ {
		if digits[i] != digits[m-1-i] {
			return false
		}
	}

	return true
}

func fairAndSquare(n uint64) bool {
	if !palindrome(n) {
		return false
	}

	s := math.Sqrt( float64(n) )
	st := math.Trunc(s)
	if (s - st) > 0.000000001 {
		return false // non-integer square root
	}

	i := uint64(s)
	if !palindrome(i) {
		return false
	}

	return true
}

func processCase(c int) {
	var min, max uint64 = 0, 0
	_, err := fmt.Scanln(&min, &max)
	if err != nil {
		fmt.Printf("Wrong format, case %d.\n", c)
		return
	}
	//fmt.Println(min, max)

	var count uint64 = 0
	var i uint64
	for i = min; i <= max; i++ {
		if fairAndSquare(i) {
			count++
			//fmt.Println(i)
		}
	}

	fmt.Printf("Case #%d: %d\n", c, count)
}

func main() {
	ntests := 0
	_, err := fmt.Scanln(&ntests)
	if err != nil {
		fmt.Printf("Wrong format.\n")
		return
	}

	for i := 1; i <= ntests; i++ {
		processCase(i)
	}
}
