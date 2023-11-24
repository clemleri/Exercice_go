package doublons7


/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement les 
nombres de k à k + n - 1 (pas forcément dans l'ordre) pour un certain k non connu à l'avance. On voudrait vérifier que 
c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de k à k + len(tab) - 1 et qui vaut false sinon

# Info
2023-2024, test 2, exercice 7
*/

func doublons(tab []int) (ok bool) {
	var min int 
	for i:=0; i<len(tab); i++ {
		min = i
		for j:=i; j<len(tab); j++ {
			if tab[min] > tab[j] {
				min = j
			}
		} 
		if min != i {
			tab[min], tab[i] = tab[i], tab[min]
		}
	}

	for y:=0; y<len(tab)-1; y++ {
		if tab[y]+1 != tab[y+1] {
			return false
		}
	}
	return true
}
