package tri2

import "math"

/*
La fonction triabs doit trier un tableau d'entiers de la plus grande valeure
absolue à la plus petite valeure absolue. Cette fonction ne doit pas modifier
le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         grand (en valeure absolue) au plus petit (en valeure absolue).

# Info
2021-2022, test2, exercice 7
*/

func triabs(tinit []int) (tfin []int) {
	for i:=0; i<len(tinit); i++{
		tfin = append(tfin,tinit[i])
	}
	if len(tfin)>1{
		var n int = len(tfin)
		var max int 
		for i:=0;i<n-1; i++ {
			max = i
			for j:=i+1; j<n; j++ {
				if math.Abs(float64(tfin[max]))< math.Abs(float64(tfin[j])){
					max = j
				}
			}
			if max != i{
				tfin[max],tfin[i] = tfin[i], tfin[max]
			}
		}
	}
	return tfin
}
