package main

import "fmt"

var peserta = []int{100, 100, 50, 40, 40, 20, 10}
var gits = []int{5, 25, 50, 120}

func main() {

	peserta = append(peserta, gits...)
	// gitsVal := make(map[int]bool, 0)
	pesertaUnqiue := make(map[int]bool, 0)
	// for _, v := range gits {
	// 	gitsVal[v] = true
	// }

	var scores []int
	for _, v := range peserta {
		// when exist dont apppend
		if pesertaUnqiue[v] {
			continue
		}
		pesertaUnqiue[v] =true
		scores = append(scores, v)
	}

// 2, 1, 3
// 1, 2, 3


//

	n := len(scores)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if scores[j] < scores[j+1] {
				scores[j], scores[j+1] = scores[j+1], scores[j]
			}
		}
	}

	scoreRankMap := make(map[int]int, 0)
	rank := 1
	for _, score := range scores {
		if _, exists := scoreRankMap[rank]; exists {
			continue
		}
		scoreRankMap[score] = rank
		rank += 1
	}
	fmt.Println(scores)
	for _, v := range gits {
		gitsRank, exists := scoreRankMap[v]
			fmt.Println(gitsRank)
		}
	}

	// sort.Sort(sort.Reverse(sort.IntSlice(peserta)))
	// fmt.Println(peserta)

	// rank := 1
	// skorAwal := peserta[0]

	// for _, skor := range peserta  {
	// 	if skor != skorAwal {
	// 		rank++
	// 	}
	// 	skorAwal = skor

	// 	for _, gitsSkor := range gits{
	// 		if skor == gitsSkor{
	// 			fmt.Printf("%d: rank %d, ", gitsSkor, rank)
	// 		}
	// 	}
	// }
}
