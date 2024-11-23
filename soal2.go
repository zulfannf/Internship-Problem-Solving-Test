package main

import (
	"fmt"
	"sort"
)


	var peserta = []int {100, 100, 50, 40, 40, 20, 10}
	var gits = []int {5, 25, 50, 120}
	

func main(){
	
	
	peserta = append(peserta,gits...)
	sort.Sort(sort.Reverse(sort.IntSlice(peserta)))
	fmt.Println(peserta)

	rank := 1
	skorAwal := peserta[0]

	for _, skor := range peserta  {
		if skor != skorAwal {
			rank++
		}
		skorAwal = skor

		for _, gitsSkor := range gits{
			if skor == gitsSkor{
				fmt.Printf("%d: rank %d, ", gitsSkor, rank)
			}
		}
	}
}

