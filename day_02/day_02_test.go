package day_02

import (
	"strings"
	"testing"
)

func TestCheckSum(t *testing.T) {
	reader := strings.NewReader(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`)
	if CheckSum(reader) != 12 {
		t.Error("Wrong checksum")
	}
}

func TestFindCorrectId(t *testing.T) {
	reader := strings.NewReader(`abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`)
	if FindTheBoxesFullOfPrototypeFabric(reader) != "fgij" {
		t.Error("Wrong checksum")
	}
}
