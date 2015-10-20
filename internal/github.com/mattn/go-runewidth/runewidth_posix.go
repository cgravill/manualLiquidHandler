// internal/github.com/mattn/go-runewidth/runewidth_posix.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
// 
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
// 
// For more information relating to the software or licensing issues please
// contact license@antha-lang.org or write to the Antha team c/o 
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

// +build !windows,!js

package runewidth

import (
	"os"
	"regexp"
	"strings"
)

var reLoc = regexp.MustCompile(`^[a-z][a-z][a-z]?(?:_[A-Z][A-Z])?\.(.+)`)

func IsEastAsian() bool {
	locale := os.Getenv("LC_CTYPE")
	if locale == "" {
		locale = os.Getenv("LANG")
	}

	// ignore C locale
	if locale == "POSIX" || locale == "C" {
		return false
	}
	if len(locale) > 1 && locale[0] == 'C' && (locale[1] == '.' || locale[1] == '-') {
		return false
	}

	charset := strings.ToLower(locale)
	r := reLoc.FindStringSubmatch(locale)
	if len(r) == 2 {
		charset = strings.ToLower(r[1])
	}

	if strings.HasSuffix(charset, "@cjk_narrow") {
		return false
	}

	for pos, b := range []byte(charset) {
		if b == '@' {
			charset = charset[:pos]
			break
		}
	}

	mbc_max := 1
	switch charset {
	case "utf-8", "utf8":
		mbc_max = 6
	case "jis":
		mbc_max = 8
	case "eucjp":
		mbc_max = 3
	case "euckr", "euccn":
		mbc_max = 2
	case "sjis", "cp932", "cp51932", "cp936", "cp949", "cp950":
		mbc_max = 2
	case "big5":
		mbc_max = 2
	case "gbk", "gb2312":
		mbc_max = 2
	}

	if mbc_max > 1 && (charset[0] != 'u' ||
		strings.HasPrefix(locale, "ja") ||
		strings.HasPrefix(locale, "ko") ||
		strings.HasPrefix(locale, "zh")) {
		return true
	}
	return false
}