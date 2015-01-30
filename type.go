// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
//
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

import (
	"math/rand"
)

// Return true if lhs greater than rhs
type GreaterThanFunc func(lhs, rhs interface{}) bool

// Return true if lhs less than rhs
type LessThanFunc GreaterThanFunc

type defaultRandSource struct{}

type Comparable interface {
	Descending() bool
	Compare(lhs, rhs interface{}) bool
}

type elementNode struct {
	next []*Element
}

type Element struct {
	elementNode
	key, Value interface{}
	score      float64
}

type SkipList struct {
	elementNode
	level      int
	length     int
	keyFunc    Comparable
	randSource rand.Source
	reversed   bool

	prevNodesCache []*elementNode // a cache for Set/Remove
}

// It is used by skip list using customized key comparing function.
// For built-in functions, there is no need to care of this interface.
//
// Every skip list element with customized key must have a score value
// to indicate its sequence.
// For any two elements with key "k1" and "k2":
// - If Compare(k1, k2) is true, k1.Score() >= k2.Score() must be true.
// - If Compare(k1, k2) is false and k1 doesn't equal to k2, k1.Score() < k2.Score() must be true.
type Scorable interface {
	Score() float64
}

func (r defaultRandSource) Int63() int64 {
	return rand.Int63()
}

func (r defaultRandSource) Seed(seed int64) {
	rand.Seed(seed)
}

func (f GreaterThanFunc) Descending() bool {
	return false
}

func (f GreaterThanFunc) Compare(lhs, rhs interface{}) bool {
	return f(lhs, rhs)
}

func (f LessThanFunc) Descending() bool {
	return true
}

func (f LessThanFunc) Compare(lhs, rhs interface{}) bool {
	return f(lhs, rhs)
}
