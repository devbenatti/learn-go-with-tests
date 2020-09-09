package maps

import "testing"

func compareStrings(t *testing.T, result, expect string) {
	t.Helper()
	if result != expect {
		t.Errorf("result '%s' - expect '%s', data '%s'", result, expect, "test")
	}
}

func compareError(t *testing.T, result, expect error) {
	t.Helper()
	if result != expect {
		t.Errorf("result error '%s' - expect '%s'", result, expect)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just one test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expect := "just one test"

		compareStrings(t, result, expect)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, result := dictionary.Search("unkown")
		compareError(t, result, ErrNotFound)
	})
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search("test")
	if err != nil {
		t.Fatal(err)
	}
	if definition != result {
		t.Errorf("result '%s' - expect '%s'", result, definition)
	}
}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "just one test"
		err := dictionary.Add(word, definition)
		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "just one test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		compareError(t, err, ErrWordExists)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "just one test"
		newDefinition := "new Definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)
		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "just one test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		compareError(t, err, ErrWordNonexistent)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "definition of test"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expect '%s' to be deleted", word)
	}
}
