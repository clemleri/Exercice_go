package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres pairs du tableau, 
dans l'ordre croissant, puis tous les nombres impairs, dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 2, exercice 4
*/

func tri(t []int) {
	var tab_pairs []int = make([]int, 0, len(t))
	var tab_impairs []int = make([]int, 0, len(t))
	for i:=0; i<len(t); i++ {
		if t[i]%2==0{
			tab_pairs = append(tab_pairs, t[i])
		}else {
			tab_impairs = append(tab_impairs,t[i])
		}
	}
	if len(tab_pairs) > 1{
		var n int = len(tab_pairs)
		var min int 
		for i:=0; i<n-1; i++ {
			min = i
			for j:=i+1; j<n; j++ {
				if tab_pairs[j]<tab_pairs[min] {
					min = j
				}
			}
			if min != i{
				tab_pairs[i], tab_pairs[min] = tab_pairs[min], tab_pairs[i]
			}
		}
	}
	if len(tab_impairs) > 1{
		n = len(tab_impairs)
		var max int 
		for i:=0; i<n-1; i++ {
			max = i
			for j:=i+1; j<n; j++ {
				if tab_impairs[j] >tab_impairs[max] {
					max = j
				}
			}
			if max != i{
				tab_impairs[i], tab_impairs[max] = tab_impairs[max], tab_impairs[i]
			}
		}
	}
	for y:=0;y<len(tab_pairs)+len(tab_impairs);y++ {
		if y<len(tab_pairs) {
			t[y] = tab_pairs[y]
		}else {
			t[y] = tab_impairs[y-len(tab_pairs)]
		}
	}
}

