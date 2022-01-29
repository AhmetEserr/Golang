package error_handlingg

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {

		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}

	if tahmin < aklimdakiSayi {
		return "Daha küçük sayı gir", nil

	}

	if tahmin > aklimdakiSayi {
		return "Daha büyük sayı gir", nil

	}

	return "Bildiniz", nil
}

func Demo2() {
	mesaj, hata := TahminEt(51)
	fmt.Println(mesaj)
	fmt.Println(hata)
}
