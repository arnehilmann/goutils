package goutils

import "strings"

func Rot13(s string) string {
	return strings.Map(rot13char, s)
}

func rot13char(c rune) rune {
	if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
		return c + 13
	} else if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
		return c - 13
	}
	return c
}
