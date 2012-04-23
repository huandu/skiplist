// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
// 
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

import (
    "testing"
    "math/rand"
)

func checkSanity(list *SkipList, t *testing.T) {
    // each level must be correctly ordered
    for k, v := range(list.next) {
        if v == nil {
            continue
        }

        if k > len(v.next) {
            t.Fatal("first node's level must be no less than current level")
        }

        prev, next := v, v.next[k]

        for next != nil {
            if !list.keyFunc(next.key, prev.key) {
                t.Fatal("next key value must be greater than prev key value")
            }

            if k > len(next.next) {
                t.Fatal("node's level must be no less than current level")
            }

            prev, next = next, next.next[k]
        }
    }
}

func TestBasicCRUD(t *testing.T) {
    list := New(String)
    checkSanity(list, t)

    list.Set("A", 1)
    list.Set("golang", 2)
    list.Set("Skip", 3)
    list.Set("List", 4)
    list.Set("Implementation", 5)
    checkSanity(list, t)

    list.Set("List", 9)
    checkSanity(list, t)

    list.Remove("a")
    list.Remove("List")
    checkSanity(list, t)

    v1 := list.Get("A")
    v2, ok2 := list.GetValue("golang")
    v3, ok3 := list.GetValue("Skip")
    v4, ok4 := list.GetValue("List")
    v5, ok5 := list.GetValue("Implementation")
    v6, ok6 := list.GetValue("not-exist")

    if v1 == nil || v1.Value.(int) != 1 || v1.Key().(string) != "A" {
        t.Fatal(`wrong "A" value`)
    }

    if v2.(int) != 2 || !ok2 {
        t.Fatal(`wrong "golang" value`)
    }

    if v3.(int) != 3 || !ok3 {
        t.Fatal(`wrong "Skip" value`)
    }

    if v4 != nil || ok4 {
        t.Fatal(`wrong "List" value`)
    }

    if v5.(int) != 5 || !ok5 {
        t.Fatal(`wrong "Implementation" value`)
    }

    if v6 != nil || ok6 {
        t.Fatal(`wrong "not-exist" value`)
    }
}

func TestChangeLevel(t *testing.T) {
    DefaultMaxLevel = 10
    list := New(IntReversed)

    if list.MaxLevel() != 10 {
        t.Fatal("max level must equal default max value")
    }

    for i := 0; i <= 200; i += 4 {
        list.Set(i, i * 10)
    }

    checkSanity(list, t)

    list.SetMaxLevel(20)
    checkSanity(list, t)

    for i := 1; i <= 201; i += 4 {
        list.Set(i, i * 10)
    }

    list.SetMaxLevel(4)
    checkSanity(list, t)

    if list.Len() != 102 {
        t.Fatal("wrong list element number", list.Len())
    }

    for c := list.Front(); c != nil; c = c.Next() {
        if c.Key().(int) * 10 != c.Value.(int) {
            t.Fatal("wrong list element value")
        }
    }
}

func BenchmarkWorstInserts(b *testing.B) {
    b.StopTimer()
    list := New(Int)
    list.SetMaxLevel(96)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        list.Set(i, i)
    }
}

func BenchmarkBestInserts(b *testing.B) {
    b.StopTimer()
    list := New(IntReversed)
    list.SetMaxLevel(96)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        list.Set(i, i)
    }
}

func BenchmarkRandomSelect(b *testing.B) {
    b.StopTimer()
    list := New(IntReversed)
    list.SetMaxLevel(96)

    for i := 0; i < b.N; i++ {
        list.Set(i, i)
    }

    keys := make([]int, b.N)

    for i := 0; i < b.N; i++ {
        keys[i] = rand.Intn(b.N)
    }

    b.StartTimer()
    for k := range(keys) {
        list.Get(k)
    }
}
