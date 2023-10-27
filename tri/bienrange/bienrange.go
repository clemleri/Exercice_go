package bienrange

/*
La fonction bienrange indique si un tableau d'entiers est bien trié par ordre
croissant ou pas.

# Entrée
- tab : le tableau d'entiers à analyser

# Sortie
- estrange : un booléen qui vaut true si les entiers de tab sont bien triés en
ordre croissant et false s'ils ne sont pas bien triés.
*/

func bienrange(tab []int) (estrange bool) {
	estrange = true
	if len(tab) >1{
		var tab_verif []int = make([]int, 0, len(tab))
		for i:=1; i<len(tab);i++ {
			tab_verif = tab[i:]
			for y:=0; y<len(tab_verif); y++ {
				if tab[i-1] > tab_verif[y]{
					estrange = false
				}
			}
		}
	}
	return estrange
}
