package decroissant

/*
La fonction decroissant doit trier un tableau d'entiers dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 7
*/

func decroissant(tab []int) {
	if len(tab) > 1{
		var n int = len(tab)
		var max int 
		for i:=0; i<n-1; i++ {
			max = i
			for j:=i+1; j<n; j++ {
				if tab[j] >tab[max] {
					max = j
				}
			}
			if max != i{
				tab[i], tab[max] = tab[max], tab[i]
			}
		}
	}
}
