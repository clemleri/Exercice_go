package decroissant

/*
La fonction decroissant doit trier un tableau d'entiers dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 7
*/

func decroissant(tab []int) {
	var max int
	for i:=0; i<len(tab)-1; i++{
		max = i
		for j:=i; j<len(tab); j++ {
			if estPlusGrandque(tab[j], tab[max]) {
				max = j
			}
		}
		if i!=max {
			tab[i], tab[max] = tab[max], tab[i]
		}
		
	}
}

func estPlusGrandque(v1, v2 int) (estplusgrand bool) {
	return v1 > v2
}
