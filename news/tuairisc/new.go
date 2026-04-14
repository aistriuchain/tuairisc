package tuairisc

import (
	_ "embed"
	"fmt"
	"strconv"
	"time"
	"unicode/utf16"
)

//go:embed logo.jpg
var Logo []byte

func NewTuairisc(oldArticleTitles []string) *tuairisc {
	return &tuairisc{
		oldArticleTitles: oldArticleTitles,
	}
}

func (a *tuairisc) GetLogo() []byte {
	return Logo
}

func (a *tuairisc) GetCopyright() []uint16 {
	copyrightString := fmt.Sprintf("© %s Copyright Tuairisc.ie - Gach ceart ar cosaint.", strconv.Itoa(time.Now().Year()))
	return utf16.Encode([]rune(copyrightString))
}
