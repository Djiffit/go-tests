package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("pekka")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get error")
		}

		assertError(t, want, err)
	})

}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{}

	t.Run("add word to dictionary", func(t *testing.T) {
		key := "test"
		want := "just a test"
		dictionary.Add(key, want)
		got, err := dictionary.Search(key)
		if err != nil {
			t.Fatal("should find: ", want)
		}
		assertStrings(t, got, want)
	})

	t.Run("error on duplicate key", func(t *testing.T) {
		key := "test"
		want := "just a test"
		dictionary.Add(key, want)
		err := dictionary.Add(key, "something else")
		assertError(t, err, ErrDuplicate)

		got, err := dictionary.Search(key)
		assertStrings(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		key := "test"
		want := "just a test"
		err := dictionary.Update(key, want)

		if err != nil {
			t.Fatal("should not get error on update ")
		}

		got, err2 := dictionary.Search("test")

		if err2 != nil {
			t.Fatal("should not get error on found item ")
		}

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		key := "testers"
		err := dictionary.Update(key, "pekka")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get error")
		}

		assertError(t, want, err)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, wanted error %q", got, want)
	}
}
