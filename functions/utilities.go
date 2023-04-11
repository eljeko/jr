//Copyright © 2022 Ugo Landini <ugo.landini@gmail.com>
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package functions

import (
	"github.com/google/uuid"
)

// Counter creates a counter named c, starting from start and incrementing by step
func Counter(c string, start, step int) int {
	val, exists := JrContext.Counters[c]
	if exists {
		JrContext.Counters[c] = val + step
		return JrContext.Counters[c]
	} else {
		JrContext.Counters[c] = start
		return start
	}
}

// UniqueId returns a random uuid
func UniqueId() string {
	return uuid.New().String()
}

// RandomBool returns a random boolean
func RandomBool() string {
	b := Random.Intn(2)
	if b == 0 {
		return "false"
	} else {
		return "true"
	}
}

// YesOrNo returns a random yes or no
func YesOrNo() string {
	b := Random.Intn(2)
	if b == 0 {
		return "no"
	} else {
		return "yes"
	}
}

// Contains checks if the str string is present in an array of string []string
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
