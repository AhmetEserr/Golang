package stringmethod

import (
	"fmt"
	s "strings"
)

//Büyük küçük ifadeleri önemlidir.
func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "E"))    // E kaç tane var
	fmt.Println(s.Contains(isim, "n")) //n ifadesi var mı yok mu
	//fmt.Println(s.Index(isim, "i"))    //kaçıncı index de yer alıyor

	sonuc := s.Index(isim, "i")
	if sonuc != -1 {
		fmt.Println("i harfi var ") // i harfinin var olupolmaıdğını döndürür.
	} else {
		fmt.Println("i harfi yok ")
	}
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
