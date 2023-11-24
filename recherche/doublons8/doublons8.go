package doublons8


/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement p fois 
chaque nombre compris entre k et k + (n/p) - 1 (pas forcément dans l'ordre) pour un certain k et un certain p non connus à l'avance (p différent de 0).
On voudrait vérifier que c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement p fois chaque entier de k à k + len(tab)/p - 1 et qui vaut false sinon

# Info
2023-2024, test 2, exercice 8
*/

func doublons(tab []int) bool {
	if len(tab) == 0 {
		return false
	}

	valeur := tab[0]
	p := 0
	for k := 0; k < len(tab); k++ {
		if tab[k] == valeur {
			p++
		}
	}

	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}

	var lst_valeurs []int 
	for y := 0; y < len(tab); y = y + p {
		lst_valeurs = append(lst_valeurs, tab[y])
		if p > 1 && y+p-1 <= len(tab)-1 {
			for x := y; x < y+p-1; x++ {
				if tab[x] != tab[x+1] {
					return false
				}
			}
		}else {
			for v := 0; v < len(tab)-1; v++ {
				if tab[v]+1 != tab[v+1] {
					return false
				}
			}
		}
	} 
	if len(tab)!=len(lst_valeurs)*p {
		return false
	}
	return true
}