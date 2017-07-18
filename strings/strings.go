// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Package ds provides useless string functions which (s)c(h)ould be in "strings" but are not (yet) there.
//  Note: `ds` is shorthand for `do strings`.
package ds // do strings

import (
	"errors"
	"strings"
)

/*
func Index(s, sep string) int
    Index returns the index of the first instance of sep in s, or -1 if sep is
    not present in s.
*/

// Extract all text bracketed by left & right
//  TODO: return error, if some right found before next left
func Extract(text, left, right string) (string, error) {
	l, r := len(left), len(right)
	var s, t string

	if l < 1 {
		return s, errors.New("Left delimiter must not be empty")
	}
	if r < 1 {
		return s, errors.New("Right delimiter must not be empty")
	}

	t = text
	for {
		if t == "" {
			return s, nil
		}
		i := strings.Index(t, left)
		if i < 0 {
			return s, nil
		}
		t = string(t[(i + l):])
		j := strings.Index(t, right)
		if j < 0 {
			return s, errors.New("No matching right delimiter")
		}
		s = s + string(t[:j])
		t = string(t[(j + r):]) // len(t)-j-r-1
	}
}

// Join2 (a convenience for strings.Join) only joins, if both args are non-empty
func Join2(a, b string, sep string) string {
	S := []string{}
	if len(a) > 0 {
		S = append(S, a)
	}
	if len(b) > 0 {
		S = append(S, b)
	}
	return strings.Join(S, sep)
}

/*
func TrimLeft(s string, cutset string) string
	TrimLeft returns a slice of the string s with all leading Unicode code
	points contained in cutset removed.

func TrimRight(s string, cutset string) string
	TrimRight returns a slice of the string s, with all trailing Unicode code
	points contained in cutset removed.
*/

// UnBracket returns what it found between bo & bc
func UnBracket(in, bo, bc string) string {
	return strings.TrimLeft(strings.TrimRight(in, bc), bo)
}

// SplitAtFirst separates `text` at first `sep` into `head` and `tail`.
// If there is no `sep`, all `text` becomes `tail`.
func SplitAtFirst(text, sep string) (head, tail string) {
	tail = text
	if len(sep) > 0 {
		pos := strings.Index(text, sep)
		if pos > -1 {
			head, tail = text[:pos], text[pos+len(sep):]
		}
	}
	return head, tail
}

// SplitAtLast separates `text` at last `sep` into `head` and `tail`.
// If there is no `sep`, all `text` becomes `head`.
func SplitAtLast(text, sep string) (head, tail string) {
	head = text
	if len(sep) > 0 {
		pos := strings.LastIndex(text, sep)
		if pos > -1 {
			head, tail = text[:pos], text[pos+len(sep):]
		}
	}
	return head, tail
}

// SplitAllPrefixe looks for all occurrences of prefix and
// returns them as head, and text less them as tail.
// Invariant: len(text) == len(head) + len(tail)
func SplitAllPrefixe(text, prefix string) (head, tail string) {
	tail = text

	lenp := len(prefix)
	if lenp > 0 {
		for pos := lenp - 1; pos < len(tail); {
			if tail[:pos] == prefix {
				tail = tail[pos:]
				head = head + prefix
			} else {
				break
			}
		}
	}
	return head, tail
}

// SplitAllSuffixe looks for all occurrences of suffix and
// returns them as tail, and text less them as head.
// Invariant: len(text) == len(head) + len(tail)
func SplitAllSuffixe(text, suffix string) (head, tail string) {
	head = text

	lens := len(suffix)
	if lens > 0 {
		for pos := len(text) - lens; pos >= 0; pos = len(head) - lens {
			if head[pos:] == suffix {
				head = head[:pos]
				tail = tail + suffix
			} else {
				break
			}
		}
	}
	return head, tail
}

// SplitAtFirstChar splits text immediately following the first Separator,
// separating it into a head and tail component.
// If there is no Separator in text, head set to text and an empty tail are returned.
// The returned values have the property that text = head+tail.
// head hasSuffix sep if tail is not empty or if text hasSuffix sep and no more sep
//
// SplitAtFirstChar panics if len(sep) != 1
func SplitAtFirstChar(text, sep string) (head, tail string) {
	if len(sep) != 1 {
		panic("SplitAtFirstChar: Separator must be a one-char string!")
	}

	i := 0
	for i < len(text) && string(text[i]) != sep {
		i++
	}
	return text[:i+1], text[i+1:]
}

// SplitAtLastChar splits text immediately following the last Separator,
// separating it into a head and tail component.
// If there is no Separator in text, an empty head and tail set to text are returned.
// The returned values have the property that text = head+tail.
// tail hasSuffix sep if head is not empty or if text hasPrefix sep and no more sep
//
// SplitAtLastChar panics if len(sep) != 1
func SplitAtLastChar(text, sep string) (head, tail string) {
	if len(sep) != 1 {
		panic("SplitAtLastChar: Separator must be a one-char string!")
	}

	i := len(text) - 1
	for i >= 0 && string(text[i]) != sep {
		i--
	}
	return text[:i+1], text[i+1:]
}
