package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}
	t.Run("in dictionary", func(t *testing.T) {
		res, err := dict.Search("test")
		if err != nil {
			t.Fatal("should find word", err)
		}
		expect := "this is just a test"
		assertStrings(t, res, expect)
	})
	t.Run("not in dictionary", func(t *testing.T) {
		_, err := dict.Search("unknown")
		expect := ErrNotFound
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, expect)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		def := "this is just a test"
		d.Add(word, def)
		assertDef(t, d, word, def)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d := Dictionary{word: def}
		err := d.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDef(t, d, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		newDef := "new definition"
		err := dict.Update(word, newDef)
		assertError(t, err, nil)
		assertDef(t, dict, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{}
		err := dict.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	w := "test"
	d := "word to be deleted"
	dict := Dictionary{w: d}
	dict.Delete(w)
	_, err := dict.Search(w)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", w)
	}
}

func assertDef(t testing.TB, d Dictionary, word, def string) {
	t.Helper()
	res, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, res, def)

}

func assertStrings(t testing.TB, res, expect string) {
	t.Helper()
	if res != expect {
		t.Errorf("res %q, expect %q, given %q", res, expect, "test")
	}
}

func assertError(t testing.TB, res, expect error) {
	t.Helper()
	if res != expect {
		t.Errorf("res %q, expect %q, given %q", res, expect, "test")
	}
}
