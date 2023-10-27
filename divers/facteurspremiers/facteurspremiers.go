package facteurspremiers


/*
La fonction facteursPremiers doit retourner un tableau contenant la liste de tous
les facteurs premiers d'un entier n, doublons compris

# Entrée
- n : un nombre entier positif

# Sortie
- facteurs : un tableau contenant tous les facteurs premiers de n, si n vaut 0 il
faut retourner un tableau à zéro éléments.

# Exemple
premiers(24) = [2 2 2 3] (l'ordre n'a pas d'importance)
*/
func facteursPremiers(n uint) (facteurs []uint) {
	if n<2 {
		return []uint{}
	}else if n >1 {
		var diviseur uint = 2
		for n>1 {
			if n%diviseur==0 {
				facteurs = append(facteurs,diviseur)
				n = n/diviseur
			}else {
				diviseur = diviseur + 1
			}
		}
	}
	
	return facteurs
}
