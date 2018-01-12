// A golang Skip List Implementation.
// https://github.com/huandu/skiplist/
//
// Copyright 2011, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/skiplist/blob/master/LICENSE

package skiplist

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func checkSanity(list *SkipList, t *testing.T) {
	// each level must be correctly ordered
	for k, v := range list.next {
		//t.Log("Level", k)

		if v == nil {
			continue
		}

		if k > len(v.next) {
			t.Fatal("first node's level must be no less than current level")
		}

		next := v
		cnt := 1

		for next.next[k] != nil {
			if !list.keyFunc.Compare(next.next[k].key, next.key) {
				t.Fatalf("next key value must be greater than prev key value. [next:%v] [prev:%v]", next.next[k].key, next.key)
			}

			if next.score > next.next[k].score {
				t.Fatalf("next key score must be no less than prev key score. [next:%v] [prev:%v]", next.next[k].score, next.score)
			}

			if k > len(next.next) {
				t.Fatalf("node's level must be no less than current level. [cur:%v] [node:%v]", k, next.next)
			}

			//t.Log("TEST VALUE", next.key, next.score, next.Value)
			next = next.next[k]
			cnt++
		}

		if k == 0 {
			if cnt != list.Len() {
				t.Fatalf("list len must match the level 0 nodes count. [cur:%v] [level0:%v]", cnt, list.Len())
			}
		}
	}
}

func testBasicIntCRUD(t *testing.T, reversed bool) {
	var list *SkipList

	if reversed {
		list = New(IntDesc)
	} else {
		list = New(Int)
	}

	list.Set(10, 1)
	list.Set(60, 2)
	list.Set(30, 3)
	list.Set(20, 4)
	list.Set(90, 5)
	t.Log("inserted")
	checkSanity(list, t)

	list.Set(30, 9)
	t.Log("inserted duplicates")
	checkSanity(list, t)

	list.Remove(0)
	list.Remove(20)
	t.Log("removed")
	checkSanity(list, t)

	v1 := list.Get(10)
	v2, ok2 := list.GetValue(60)
	v3, ok3 := list.GetValue(30)
	v4, ok4 := list.GetValue(20)
	v5, ok5 := list.GetValue(90)
	v6, ok6 := list.GetValue(-1)

	if v1 == nil || v1.Value.(int) != 1 || v1.Key().(int) != 10 {
		t.Fatal(`wrong "10" value`, v1)
	}

	if v2 == nil || v2.(int) != 2 || !ok2 {
		t.Fatal(`wrong "60" value`)
	}

	if v3 == nil || v3.(int) != 9 || !ok3 {
		t.Fatal(`wrong "30" value`)
	}

	if v4 != nil || ok4 {
		t.Fatal(`wrong "20" value`)
	}

	if v5 == nil || v5.(int) != 5 || !ok5 {
		t.Fatal(`wrong "90" value`)
	}

	if v6 != nil || ok6 {
		t.Fatal(`wrong "-1" value`)
	}
}

func TestBasicIntCRUDNormal(t *testing.T) {
	testBasicIntCRUD(t, false)
}

func TestBasicIntCRUDDesc(t *testing.T) {
	testBasicIntCRUD(t, true)
}

func testBasicStringCRUD(t *testing.T, reversed bool) {
	var list *SkipList

	if reversed {
		list = New(StringDesc)
	} else {
		list = New(String)
	}

	list.Set("A", 1)
	list.Set("golang", 2)
	list.Set("Skip", 3)
	list.Set("List", 4)
	list.Set("Implementation", 5)
	t.Log("inserted")
	checkSanity(list, t)

	list.Set("List", 9)
	t.Log("inserted duplicates")
	checkSanity(list, t)

	list.Remove("a")
	list.Remove("List")
	t.Log("removed")
	checkSanity(list, t)

	v1 := list.Get("A")
	v2, ok2 := list.GetValue("golang")
	v3, ok3 := list.GetValue("Skip")
	v4, ok4 := list.GetValue("List")
	v5, ok5 := list.GetValue("Implementation")
	v6, ok6 := list.GetValue("not-exist")

	if v1 == nil || v1.Value.(int) != 1 || v1.Key().(string) != "A" {
		t.Fatal(`wrong "A" value`, v1)
	}

	if v2 == nil || v2.(int) != 2 || !ok2 {
		t.Fatal(`wrong "golang" value`, v2)
	}

	if v3 == nil || v3.(int) != 3 || !ok3 {
		t.Fatal(`wrong "Skip" value`, v3)
	}

	if v4 != nil || ok4 {
		t.Fatal(`wrong "List" value`, v4)
	}

	if v5 == nil || v5.(int) != 5 || !ok5 {
		t.Fatal(`wrong "Implementation" value`, v5)
	}

	if v6 != nil || ok6 {
		t.Fatal(`wrong "not-exist" value`, v6)
	}
}

func TestBasicStringCRUDNormal(t *testing.T) {
	testBasicStringCRUD(t, false)
}

func TestBasicStringCRUDDesc(t *testing.T) {
	testBasicStringCRUD(t, true)
}

func TestChangeLevel(t *testing.T) {
	DefaultMaxLevel = 10
	list := New(IntDesc)

	if list.MaxLevel() != 10 {
		t.Fatal("max level must equal default max value")
	}

	for i := 0; i <= 200; i += 4 {
		list.Set(i, i*10)
	}

	checkSanity(list, t)

	list.SetMaxLevel(20)
	checkSanity(list, t)

	for i := 1; i <= 201; i += 4 {
		list.Set(i, i*10)
	}

	list.SetMaxLevel(4)
	checkSanity(list, t)

	if list.Len() != 102 {
		t.Fatal("wrong list element number", list.Len())
	}

	for c := list.Front(); c != nil; c = c.Next() {
		if c.Key().(int)*10 != c.Value.(int) {
			t.Fatal("wrong list element value")
		}
	}

	DefaultMaxLevel = 32
}

func BenchmarkDefaultWorstInserts(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	list := New(Int)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		list.Set(i, i)
	}

	list = nil
}

func BenchmarkDefaultBestInserts(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	list := New(IntDesc)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		list.Set(i, i)
	}

	list = nil
}

func BenchmarkWorstInsertsWithLocalRandSource(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	list := New(Int)
	source := rand.NewSource(time.Now().UnixNano())
	list.SetRandSource(source)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		list.Set(i, i)
	}

	list = nil
}

func BenchmarkBestInsertsWithLocalRandSource(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	list := New(IntDesc)
	source := rand.NewSource(time.Now().UnixNano())
	list.SetRandSource(source)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		list.Set(i, i)
	}

	list = nil
}

func BenchmarkRandomSelect(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	list := New(IntDesc)

	for i := 0; i < b.N; i++ {
		list.Set(i, i)
	}

	keys := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		keys[i] = rand.Intn(b.N)
	}

	b.StartTimer()
	for _, k := range keys {
		list.Get(k)
	}

	list = nil
}
