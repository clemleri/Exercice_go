package moitie

/*
La fonction moitieDePairs indique si au 
moins de la moitié des nombres contenus dans 
un tableau sont pairs

# Entrée
- t : un tableau d'entiers

# Sortie
- reponse : un booléen qui vaut true si au moins la moitié des nombres de t sont pairs et false sinon

# Info
2022-2023, test 4, exercice 4
*/

func moitieDePairs(t []int) (reponse bool) {
	cpt_pos := 0
	for i:=0; i<len(t); i++ {
		if t[i]%2==0{
			cpt_pos++
		}
	}
	if len(t)%2==0 {
		return cpt_pos>=len(t)/2
	}else {
		return cpt_pos>len(t)/2
	}
}
