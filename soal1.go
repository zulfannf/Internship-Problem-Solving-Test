package main

import "fmt"

func main(){
	var nilai int
	fmt.Println("Silahkan masukan nilai:")
	_,err := fmt.Scanf("%d", &nilai)
	if err != nil {
		fmt.Println("Error, silahkan masukan nilai berupa angka")
		return
	}

	deret := make([]int, nilai)
	hasil := 1
	for i := 0; i < nilai; i++ {
		hasil += i
		deret[i] = hasil
	}
	fmt.Println("Output: ", deret)
}