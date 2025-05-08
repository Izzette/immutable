package immutable

import (
	"testing"
)

func TestSetsPut(t *testing.T) {
	s := NewSet[string](nil)
	s2 := s.Add("1").Add("1")
	s2.Add("2")
	if s.Len() != 0 {
		t.Fatalf("Unexpected mutation of set")
	}
	if s.Has("1") {
		t.Fatalf("Unexpected set element")
	}
	if s2.Len() != 1 {
		t.Fatalf("Unexpected non-mutation of set")
	}
	if !s2.Has("1") {
		t.Fatalf("Set element missing")
	}
	itr := s2.Iterator()
	counter := 0
	for !itr.Done() {
		i, v := itr.Next()
		t.Log(i, v)
		counter++
	}
	if counter != 1 {
		t.Fatalf("iterator wrong length")
	}
}

func TestSetsDelete(t *testing.T) {
	s := NewSet[string](nil)
	s2 := s.Add("1")
	s3 := s.Delete("1")
	if s2.Len() != 1 {
		t.Fatalf("Unexpected non-mutation of set")
	}
	if !s2.Has("1") {
		t.Fatalf("Set element missing")
	}
	if s3.Len() != 0 {
		t.Fatalf("Unexpected set length after delete")
	}
	if s3.Has("1") {
		t.Fatalf("Unexpected set element after delete")
	}
}

func TestSetBuilder(t *testing.T) {
	b := NewSetBuilder[string](nil)
	b.Set("test3")
	b.Set("test1")
	b.Set("test2")

	s := b.ToSet()
	if s.Len() != 3 {
		t.Fatalf("Set has wrong number of items")
	}
	if !s.Has("test1") {
		t.Fatalf("Missing item test1")
	}
	if !s.Has("test2") {
		t.Fatalf("Missing item test2")
	}
	if !s.Has("test3") {
		t.Fatalf("Missing item test3")
	}
	if s.Has("test4") {
		t.Fatalf("Unexpected item test4")
	}
}

func TestSortedSetsPut(t *testing.T) {
	s := NewSortedSet[string](nil)
	s2 := s.Add("1").Add("1").Add("0")
	if s.Len() != 0 {
		t.Fatalf("Unexpected mutation of set")
	}
	if s.Has("1") {
		t.Fatalf("Unexpected set element")
	}
	if s2.Len() != 2 {
		t.Fatalf("Unexpected non-mutation of set")
	}
	if !s2.Has("1") {
		t.Fatalf("Set element missing")
	}

	itr := s2.Iterator()
	counter := 0
	for !itr.Done() {
		i, v := itr.Next()
		t.Log(i, v)
		if counter == 0 && i != "0" {
			t.Fatalf("sort did not work for first el")
		}
		if counter == 1 && i != "1" {
			t.Fatalf("sort did not work for second el")
		}
		counter++
	}
	if counter != 2 {
		t.Fatalf("iterator wrong length")
	}
}

func TestSortedSetsDelete(t *testing.T) {
	s := NewSortedSet[string](nil)
	s2 := s.Add("1")
	s3 := s.Delete("1")
	if s2.Len() != 1 {
		t.Fatalf("Unexpected non-mutation of set")
	}
	if !s2.Has("1") {
		t.Fatalf("Set element missing")
	}
	if s3.Len() != 0 {
		t.Fatalf("Unexpected set length after delete")
	}
	if s3.Has("1") {
		t.Fatalf("Unexpected set element after delete")
	}
}

func TestSortedSetBuilder(t *testing.T) {
	b := NewSortedSetBuilder[string](nil)
	b.Set("test3")
	b.Set("test1")
	b.Set("test2")

	s := b.SortedSet()
	items := s.Items()

	if len(items) != 3 {
		t.Fatalf("Set has wrong number of items")
	}
	if items[0] != "test1" {
		t.Fatalf("First item incorrectly sorted")
	}
	if items[1] != "test2" {
		t.Fatalf("Second item incorrectly sorted")
	}
	if items[2] != "test3" {
		t.Fatalf("Third item incorrectly sorted")
	}
}
