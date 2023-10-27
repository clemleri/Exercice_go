package nombrespremiers

/*
La fonction premiers doit retourner un tableau contenant les nombres
premiers compris entre 0 et un entier n.

# Entrée
- n : un nombre entier

# Sortie
- p : un tableau contenant tous les nombres premiers compris entre 0 et n inclus,
s'il n'existe pas de tels nombres premiers, il faut que p soit un tableau contenant 0 éléments

# Exemple
premiers(10) = [2 3 5 7] (l'ordre n'a pas d'importance)
*/
func premiers(n int) (p []int) {
	var verif_premier bool = true
	if n>0 {
		for i:=0; i<n;i++ {
			for y:=2; y<i; y++ {
				if i%y==0 {
					verif_premier = false
				}
			}
			if verif_premier {
				p = append(p, i)
			}
			verif_premier = false
		}
	}else{
		return p
	}
	return p
}

