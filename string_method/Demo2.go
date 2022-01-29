package stringmethod

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng"))
	fmt.Println(s.HasSuffix(isim, "en"))
	fmt.Println(s.Index(isim, "n"))
	harfler := []string{"a", "h", "m", "e", "t"}
	sonuc := s.Join(harfler, "--")
	//fmt.Println(s.Join(harfler, "--"))
	fmt.Println(s.Replace(sonuc, "*", "+", 3))
	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 5))

}
