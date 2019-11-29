// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
//
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

// Next returns the ajancent next element.
func (element *Element) Next() *Element {
	return element.next[0]
}

// NextLevel returns next element at a specific level.
func (element *Element) NextLevel(level int) *Element {
	if level >= len(element.next) || level < 0 {
		panic("invalid argument to NextLevel")
	}

	return element.next[level]
}

// Key returns the key of element.
func (element *Element) Key() interface{} {
	return element.key
}
