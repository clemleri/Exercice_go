package ensemble

/*
Un ensemble d'entier est un paquet de plusieurs entiers, sans doublons.
La fonction estEnsemble doit indiquer si en tableau d'entiers correspond à un
ensemble ou non.

# Entrée
- E : un tableau d'entiers

# Sortie
- b : un booléen indiquant si E est bien un ensemble ou non (nil n'est pas un ensemble)

# Exemples
estEnsemble([]int{1, 2, 3}) = true
estEnsemble([]int{1, 2, 2}) = false
*/
func estEnsemble(E []int) (b bool) {
	var tab_verif []int = make([]int,0,len(E))
	b = true
	if len(E)>0{
		for i:=0;i<len(E);i++ {
			tab_verif = E[i+1:]
			for y:=0; y<len(tab_verif);y++ {
				if E[i] == tab_verif[y]{
					b = false
				}
			}
		}
	}else if E == nil{
		b = false
	}
	return b
}
