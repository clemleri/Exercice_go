package valabs

import "math"

/*
La fonction valabs doit trier un tableau d'entiers de la plus petite valeur absolue
à la plus grande valeur absolue. En cas
d'égalité de valeur absolue, les nombres
négatifs doivent être placés avant les
nombres positifs.

# Entrée
- tab : un tableau d'entiers
*/

func valabs(tab []int) {
	if len(tab)>1{
		var n int = len(tab)
		var min int 
		for i:=0;i<n-1; i++ {
			min = i
			for j:=i+1; j<n; j++ {
				if math.Abs(float64(tab[min]))>math.Abs(float64(tab[j])){
					min = j
				}else if math.Abs(float64(tab[min]))==math.Abs(float64(tab[j])) {
					if tab[min]==tab[j] {
						continue
					}else if tab[min]>tab[j]{
						min = j
					}
				}
			}
			if min != i{
				tab[min],tab[i] = tab[i], tab[min]
			}
		}
	}
}
