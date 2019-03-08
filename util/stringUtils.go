package util

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

type ReplacerFunc func(index int, start int, end int, content string) (string, error)

func Left(str string, length int) string {
	if str == "" || length < 0 {
		return ""
	}
	strRune := []rune(str)
	if length < len(strRune) {
		return string(strRune[:length])
	} else {
		return str
	}
}

func Right(str string, length int) string {
	if str == "" || length < 0 {
		return ""
	}
	strRune := []rune(str)
	strLen := len(strRune)
	if length < strLen {
		return string(strRune[strLen-length:])
	} else {
		return str
	}
}

func Substr(str string, beginIndex int, endIndex int) (string, error) {
	if str == "" {
		return "", nil
	}
	if beginIndex < 0 {
		return "", fmt.Errorf("String index out of range: %d", beginIndex)
	}
	strRune := []rune(str)
	strLen := len(strRune)
	if endIndex > strLen {
		return "", fmt.Errorf("String index out of range: %d", endIndex)
	}
	subLen := endIndex - beginIndex
	if subLen < 0 {
		return "", fmt.Errorf("String index out of range: %d", subLen)
	}
	return string(strRune[beginIndex:endIndex]), nil
}

func ReplaceBetween(str string, open string, close string, replacer ReplacerFunc) (string, []string, error) {
	if str == "" {
		return "", nil, nil
	}
	strLen := utf8.RuneCountInString(str)
	openLen := len(open)
	closeLen := len(close)
	pos := 0
	index := 0
	var buffer bytes.Buffer
	contents := make([]string, 0, 10)
	for {
		if pos < strLen-closeLen {
			start := strings.Index(str[pos:], open)
			if start < 0 {
				break
			}
			start = pos + start + openLen
			end := strings.Index(str[start:], close)
			if end < 0 {
				break
			}
			end += start
			buffer.WriteString(str[pos : start-openLen])
			content := str[start:end]
			contents = append(contents, content)
			index++
			newContent, err := replacer(index, start-openLen, end, content)
			if err != nil {
				return "", nil, err
			}
			buffer.WriteString(newContent)
			pos = end + closeLen
		} else {
			break
		}
	}
	buffer.WriteString(str[pos:])
	return buffer.String(), contents, nil
}

func IndexOf(str string, substr string, fromIndex int) int {
	strLen := utf8.RuneCountInString(str)
	if fromIndex >= strLen {
		if substr == "" {
			return strLen
		}
		return -1
	}
	if fromIndex < 0 {
		fromIndex = 0
	}
	if substr == "" {
		return fromIndex
	}
	index := strings.Index(str[fromIndex:], substr)
	if index < 0 {
		return -1
	}
	return fromIndex + index
}
