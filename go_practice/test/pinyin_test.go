// :date 2021/12/27

package test

import (
	"fmt"
	"github.com/chain-zhang/pinyin"
	"testing"
	"unicode"
)

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func TestPinYin(t *testing.T) {
	res, err := pinyin.New("Server 广 告").Split(" ").Mode(pinyin.InitialsInCapitals).Convert()
	fmt.Println(res)
	if err != nil {
		t.Fatal(err)
	}

	//modArgs := pinyin.NewArgs()
	//fmt.Println(pinyin.Pinyin("Server 广 告", modArgs))
	//
	//fmt.Println(pinyin.Slug("Server 广 告", modArgs))

}
