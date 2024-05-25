package utils

import (
	"math"
	"strings"
	"unicode"
)

func DetectEncoding(s string) string {
	if IsASCII(s) {
		return "Text"
	}
	if IsUnicode(s) {
		return "Unicode"
	}

	return ""
}

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

// isUnicode checks if the string contains any Unicode characters.
func IsUnicode(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}

func GetCharacterCount(s string) int {
	return strings.Count(s, "")
}

func GetGSMCount(s string) int {
	gsmChars := "@£$¥èéùìòÇ\nØø\rÅåΔ_ΦΓΛΩΠΨΣΘΞ\x1BÆæßÉ \"!\"#¤%&'()*+,-./0123456789:;<=>?¡ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÑÜ§¿abcdefghijklmnopqrstuvwxyzäöñüà"
	gsmCharMap := make(map[rune]struct{})

	for _, char := range gsmChars {
		gsmCharMap[char] = struct{}{}
	}

	gsmCount := 0
	for _, char := range s {
		if _, exists := gsmCharMap[char]; exists {
			gsmCount++
		}
	}

	return gsmCount
}

func GetSMSCount(s string) int {
	if IsUnicode(s) {

		return int(math.Ceil(float64(GetCharacterCount(s)) / float64(70)))
	}

	return int(math.Ceil(float64(GetCharacterCount(s)) / float64(160)))
}

func GetUnicodeCount(s string) int {
	count := 0

	for _, char := range s {
		if char > unicode.MaxASCII {
			count++
		}
	}

	return count
}
