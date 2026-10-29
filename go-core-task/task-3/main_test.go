package main

import "testing"

func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	if m.m == nil {
		t.Fatal("ожидалась инициализированная карта, получен nil")
	}
	if len(m.m) != 0 {
		t.Fatalf("ожидалась пустая карта, получено %d элементов", len(m.m))
	}
}

func TestStringIntMapAdd(t *testing.T) {
	tests := []struct {
		name  string
		pairs map[string]int
		key   string
		value int
	}{
		{
			name:  "добавление в пустую карту",
			pairs: map[string]int{},
			key:   "one",
			value: 1,
		},
		{
			name:  "добавление нового ключа",
			pairs: map[string]int{"one": 1, "two": 2},
			key:   "three",
			value: 3,
		},
		{
			name:  "перезапись существующего ключа",
			pairs: map[string]int{"one": 1},
			key:   "one",
			value: 100,
		},
		{
			name:  "добавление пустого ключа",
			pairs: map[string]int{},
			key:   "",
			value: 0,
		},
		{
			name:  "отрицательное значение",
			pairs: map[string]int{},
			key:   "negative",
			value: -42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			for key, value := range tt.pairs {
				m.Add(key, value)
			}

			m.Add(tt.key, tt.value)

			got, ok := m.Get(tt.key)
			if !ok {
				t.Fatalf("ключ %q отсутствует после добавления", tt.key)
			}
			if got != tt.value {
				t.Fatalf("Get(%q) = %d, ожидалось %d", tt.key, got, tt.value)
			}

			wantLength := len(tt.pairs) + 1
			if keyExists(tt.pairs, tt.key) {
				wantLength = len(tt.pairs)
			}
			if len(m.m) != wantLength {
				t.Fatalf("ожидалось %d элементов, получено %d", wantLength, len(m.m))
			}
		})
	}
}

func TestStringIntMapRemove(t *testing.T) {
	tests := []struct {
		name       string
		pairs      map[string]int
		key        string
		wantExists bool
		wantLength int
	}{
		{
			name:       "удаление существующего ключа",
			pairs:      map[string]int{"one": 1, "two": 2},
			key:        "one",
			wantExists: false,
			wantLength: 1,
		},
		{
			name:       "удаление несуществующего ключа",
			pairs:      map[string]int{"one": 1},
			key:        "missing",
			wantExists: false,
			wantLength: 1,
		},
		{
			name:       "удаление из пустой карты",
			pairs:      map[string]int{},
			key:        "one",
			wantExists: false,
			wantLength: 0,
		},
		{
			name:       "удаление единственного элемента",
			pairs:      map[string]int{"one": 1},
			key:        "one",
			wantExists: false,
			wantLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			for key, value := range tt.pairs {
				m.Add(key, value)
			}

			m.Remove(tt.key)

			if got := m.Exists(tt.key); got != tt.wantExists {
				t.Fatalf("Exists(%q) = %v, ожидалось %v", tt.key, got, tt.wantExists)
			}
			if len(m.m) != tt.wantLength {
				t.Fatalf("ожидалось %d элементов, получено %d", tt.wantLength, len(m.m))
			}
		})
	}
}

func TestStringIntMapExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)

	tests := []struct {
		name string
		key  string
		want bool
	}{
		{name: "существующий ключ", key: "one", want: true},
		{name: "несуществующий ключ", key: "two", want: false},
		{name: "пустой ключ", key: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.Exists(tt.key); got != tt.want {
				t.Fatalf("Exists(%q) = %v, ожидалось %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestStringIntMapGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("zero", 0)
	m.Add("negative", -5)

	tests := []struct {
		name      string
		key       string
		wantValue int
		wantOK    bool
	}{
		{name: "существующий ключ", key: "one", wantValue: 1, wantOK: true},
		{name: "значение по умолчанию", key: "zero", wantValue: 0, wantOK: true},
		{name: "отрицательное значение", key: "negative", wantValue: -5, wantOK: true},
		{name: "несуществующий ключ", key: "two", wantValue: 0, wantOK: false},
		{name: "пустой ключ", key: "", wantValue: 0, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOK := m.Get(tt.key)
			if gotValue != tt.wantValue {
				t.Fatalf("Get(%q) value = %d, ожидалось %d", tt.key, gotValue, tt.wantValue)
			}
			if gotOK != tt.wantOK {
				t.Fatalf("Get(%q) ok = %v, ожидалось %v", tt.key, gotOK, tt.wantOK)
			}
		})
	}
}

func TestStringIntMapGetFromEmptyMap(t *testing.T) {
	m := NewStringIntMap()

	gotValue, gotOK := m.Get("missing")

	if gotValue != 0 {
		t.Fatalf("ожидалось значение 0, получено %d", gotValue)
	}
	if gotOK {
		t.Fatal("ожидалось ok = false для пустой карты")
	}
}

func TestStringIntMapCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	copied := m.Copy()

	if len(copied) != 3 {
		t.Fatalf("ожидалось 3 элемента в копии, получено %d", len(copied))
	}
	for key, want := range map[string]int{"one": 1, "two": 2, "three": 3} {
		got, ok := copied[key]
		if !ok {
			t.Fatalf("ключ %q отсутствует в копии", key)
		}
		if got != want {
			t.Fatalf("копия[%q] = %d, ожидалось %d", key, got, want)
		}
	}
}

func TestStringIntMapCopyEmpty(t *testing.T) {
	m := NewStringIntMap()

	copied := m.Copy()

	if copied == nil {
		t.Fatal("ожидалась не nil карта")
	}
	if len(copied) != 0 {
		t.Fatalf("ожидалась пустая копия, получено %d элементов", len(copied))
	}
}

func TestStringIntMapCopyIsIndependent(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	copied := m.Copy()

	copied["one"] = 100
	copied["new"] = 3
	delete(copied, "two")

	if got, _ := m.Get("one"); got != 1 {
		t.Fatalf("изменение копии повлияло на оригинал: Get(\"one\") = %d, ожидалось 1", got)
	}
	if m.Exists("new") {
		t.Fatal("новый ключ копии появился в оригинале")
	}
	if !m.Exists("two") {
		t.Fatal("удаление ключа из копии повлияло на оригинал")
	}

	m.Add("original", 4)
	if _, ok := copied["original"]; ok {
		t.Fatal("изменение оригинала повлияло на копию")
	}
}

func keyExists(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
