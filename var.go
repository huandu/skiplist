// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
// 
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

import "bytes"

const PROPABILITY = 0x3FFF

var (
    DefaultMaxLevel int = 32

    Byte GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(byte) > rhs.(byte)
    }
    ByteAscending               = Byte
    ByteDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(byte) < rhs.(byte)
    }

    Float32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(float32) > rhs.(float32)
    }
    Float32Ascending               = Float32
    Float32Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(float32) < rhs.(float32)
    }

    Float64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(float64) > rhs.(float64)
    }
    Float64Ascending               = Float64
    Float64Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(float64) < rhs.(float64)
    }

    Int GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int) > rhs.(int)
    }
    IntAscending               = Int
    IntDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int) < rhs.(int)
    }

    Int16 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int16) > rhs.(int16)
    }
    Int16Ascending               = Int16
    Int16Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int16) < rhs.(int16)
    }

    Int32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int32) > rhs.(int32)
    }
    Int32Ascending               = Int32
    Int32Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int32) < rhs.(int32)
    }

    Int64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int64) > rhs.(int64)
    }
    Int64Ascending               = Int64
    Int64Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int64) < rhs.(int64)
    }

    Int8 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int8) > rhs.(int8)
    }
    Int8Ascending               = Int8
    Int8Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(int8) < rhs.(int8)
    }

    Rune GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(rune) > rhs.(rune)
    }
    RuneAscending               = Rune
    RuneDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(rune) < rhs.(rune)
    }

    String GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(string) > rhs.(string)
    }
    StringAscending               = String
    StringDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(string) < rhs.(string)
    }

    Uint GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint) > rhs.(uint)
    }
    UintAscending               = Uint
    UintDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint) < rhs.(uint)
    }

    Uint16 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint16) > rhs.(uint16)
    }
    Uint16Ascending               = Uint16
    Uint16Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint16) < rhs.(uint16)
    }

    Uint32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint32) > rhs.(uint32)
    }
    Uint32Ascending               = Uint32
    Uint32Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint32) < rhs.(uint32)
    }

    Uint64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint64) > rhs.(uint64)
    }
    Uint64Ascending               = Uint64
    Uint64Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint64) < rhs.(uint64)
    }

    Uint8 GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint8) > rhs.(uint8)
    }
    Uint8Ascending               = Uint8
    Uint8Descending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uint8) < rhs.(uint8)
    }

    Uintptr GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uintptr) > rhs.(uintptr)
    }
    UintptrAscending               = Uintptr
    UintptrDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return lhs.(uintptr) < rhs.(uintptr)
    }

    // the type []byte.
    Bytes GreaterThanFunc = func(lhs, rhs interface{}) bool {
        return bytes.Compare(lhs.([]byte), rhs.([]byte)) > 0
    }
    BytesAscending = Bytes
    // the type []byte. reversed order.
    BytesDescending LessThanFunc = func(lhs, rhs interface{}) bool {
        return bytes.Compare(lhs.([]byte), rhs.([]byte)) < 0
    }
)
