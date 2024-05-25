package utils

import "testing"

// Generic for expected result
type TestCase[T any] struct {
	name     string
	s        string
	expected T
}

func TestDetectEncoding(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "ISO-8859-1 English",
			s:        "Hello, World",
			expected: "Text",
		},
		{
			name:     "UTF-8 Chinese",
			s:        "你好，世界",
			expected: "Unicode",
		},
		{
			name:     "UTF-8 Japanese",
			s:        "こんにちは世界",
			expected: "Unicode",
		},
		{
			name:     "UTF-8 Korean",
			s:        "안녕하세요 세계",
			expected: "Unicode",
		},
		{
			name:     "ISO-8859-1 French",
			s:        "Bonjour le monde",
			expected: "Text",
		},
		{
			name:     "ISO-8859-1 Spanish",
			s:        "Hola, Mundo",
			expected: "Text",
		},
		{
			name:     "Windows-1251 Russian",
			s:        "Привет, мир",
			expected: "Unicode",
		},
		{
			name:     "KOI8-R Russian",
			s:        "Привет, мир",
			expected: "Unicode",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run(tc.name, func(t *testing.T) {
				detected := DetectEncoding(tc.s)
				if detected != tc.expected {
					t.Errorf("Expected %s, but got %s", tc.expected, detected)
				}
			})
		})
	}
}

func TestGetLength(t *testing.T) {
	testCases := []TestCase[int]{
		{
			name:     "English Length",
			s:        "Hello, World",
			expected: 12,
		},
		{
			name:     "Chinese Length",
			s:        "你好，世界",
			expected: 6,
		},
		{
			name:     "Korean Length",
			s:        "안녕하세요 세계",
			expected: 5,
		},
		{
			name:     "French Length",
			s:        "Bonjour le monde",
			expected: 14,
		},
		{
			name:     "Spanish Length",
			s:        "Hola, Mundo",
			expected: 8,
		},
		{
			name:     "Russian Length",
			s:        "Привет, мир",
			expected: 7,
		},
		{
			name:     "Japanese Length",
			s:        "こんにちは世界",
			expected: 6,
		},
		{
			name:     "Bangla Length",
			s:        "আমি ভাল আছই",
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			length := GetCharacterCount(tc.s)
			if length != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, length)
			}
		})
	}
}
