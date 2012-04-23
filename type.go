// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
// 
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

// return true if lhs greater than rhs
type GreaterThanFunc func(lhs, rhs interface{}) bool

type elementNode struct {
	next []*Element
}

type Element struct {
	elementNode
	key, Value interface{}
}

type SkipList struct {
	elementNode
	level   int
	length  int
	keyFunc GreaterThanFunc
}
