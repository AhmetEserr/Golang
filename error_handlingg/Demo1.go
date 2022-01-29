package error_handlingg

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")

	if err != nil {
		if pEer, ok := err.(*os.PathError); ok {
			fmt.Println(" Dosya Bulunamadı : ", pEer.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadı !!")
			return
		}

	} else {
		fmt.Println(f.Name())
	}

}
