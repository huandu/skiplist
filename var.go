// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
// 
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

import "bytes"

var (
	DefaultMaxLevel int = 24

	Byte GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(byte) > rhs.(byte)
	}
	ByteReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(byte) < rhs.(byte)
	}

	Float32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float32) > rhs.(float32)
	}
	Float32Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float32) < rhs.(float32)
	}

	Float64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float64) > rhs.(float64)
	}
	Float64Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	}

	Int GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int) > rhs.(int)
	}
	IntReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int) < rhs.(int)
	}

	Int16 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int16) > rhs.(int16)
	}
	Int16Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int16) < rhs.(int16)
	}

	Int32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int32) > rhs.(int32)
	}
	Int32Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int32) < rhs.(int32)
	}

	Int64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int64) > rhs.(int64)
	}
	Int64Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int64) < rhs.(int64)
	}

	Int8 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int8) > rhs.(int8)
	}
	Int8Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int8) < rhs.(int8)
	}

	Rune GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(rune) > rhs.(rune)
	}
	RuneReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(rune) < rhs.(rune)
	}

	String GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(string) > rhs.(string)
	}
	StringReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(string) < rhs.(string)
	}

	Uint GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint) > rhs.(uint)
	}
	UintReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint) < rhs.(uint)
	}

	Uint16 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint16) > rhs.(uint16)
	}
	Uint16Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint16) < rhs.(uint16)
	}

	Uint32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint32) > rhs.(uint32)
	}
	Uint32Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint32) < rhs.(uint32)
	}

	Uint64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint64) > rhs.(uint64)
	}
	Uint64Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint64) < rhs.(uint64)
	}

	Uint8 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint8) > rhs.(uint8)
	}
	Uint8Reversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint8) < rhs.(uint8)
	}

	Uintptr GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uintptr) > rhs.(uintptr)
	}
	UintptrReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uintptr) < rhs.(uintptr)
	}

	// the type []byte.
	Bytes GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return bytes.Compare(lhs.([]byte), rhs.([]byte)) > 0
	}
	// the type []byte. reversed order.
	BytesReversed GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return bytes.Compare(lhs.([]byte), rhs.([]byte)) < 0
	}
)
