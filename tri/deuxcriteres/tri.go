package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres strictement négatifs du tableau
, dans l'ordre décroissant, puis tous les nombres positifs ou nuls, dans l'ordre croissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 8
*/

func tri(t []int) {
	var tab_neg []int = make([]int, 0, len(t))
	var tab_pos []int = make([]int, 0, len(t))
	for i:=0; i<len(t); i++ {
		if t[i]>0{
			tab_pos = append(tab_pos, t[i])
		}else {
			tab_neg = append(tab_neg,t[i])
		}
	}
	if len(tab_pos) > 1{
		var n int = len(tab_pos)
		var min int 
		for i:=0; i<n-1; i++ {
			min = i
			for j:=i+1; j<n; j++ {
				if tab_pos[j]<tab_pos[min] {
					min = j
				}
			}
			if min != i{
				tab_pos[i], tab_pos[min] = tab_pos[min], tab_pos[i]
			}
		}
	}
	if len(tab_neg) > 1{
		var n int = len(tab_neg)
		var max int 
		for i:=0; i<n-1; i++ {
			max = i
			for j:=i+1; j<n; j++ {
				if tab_neg[j] >tab_neg[max] {
					max = j
				}
			}
			if max != i{
				tab_neg[i], tab_neg[max] = tab_neg[max], tab_neg[i]
			}
		}
	}
	for y:=0;y<len(tab_pos)+len(tab_neg);y++ {
		if y<len(tab_neg) {
			t[y] = tab_neg[y]
		}else {
			t[y] = tab_pos[y-len(tab_neg)]
		}
	}
}
