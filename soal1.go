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

var result[]int 

for i := 0; i < nilai ; i++ {
	// equation = n * (n+1)/2 +1
	current := i * (i+1)/2+1 
	result = append(result, current)
	}
	fmt.Println(result)

}

