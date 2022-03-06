package main

import (
	"strings"
	"testing"
)

func TestBoxIDChecksum(t *testing.T) {
	r := strings.NewReader(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`)
	got, err := BoxIDChecksum(r)
	if err != nil {
		t.Errorf("BoxIDChecksum() error = %v", err)
		return
	}
	if want := 12; got != want {
		t.Errorf("BoxIDChecksum() got = %v, want %v", got, want)
	}
}

func TestBoxIDCommonLetters(t *testing.T) {
	r := strings.NewReader(`abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`)
	got, err := BoxIDCommonLetters(r)
	if err != nil {
		t.Errorf("BoxIDCommonLetters() error = %v", err)
		return
	}
	if want := "fgij"; got != want {
		t.Errorf("BoxIDCommonLetters() got = %v, want %v", got, want)
	}
}
