package deuxpetits

/*
La fonction lesDeuxPlusPetits retourne les deux plus petits entiers présents dans un tableau

# Entrée
- t : un tableau d'entier qui en contient au moins 2

# Sorties
- v1 : le plus petit entier dans t
- v2 : le deuxième plus petit entier dans t

# Info
2022-2023, test 4, exercice 3
*/

func lesDeuxPlusPetits(t []int) (v1, v2 int) {
	if len(t)==2 {
		if t[0]>=t[1] {
			return t[1], t[0]
		}else {
			return t[0], t[1]
		}
	}else {
		v1 = t[0]
		for i:=0; i<len(t); i++ {
			if v1>=t[i]{
				v1, v2 = t[i], v1
			}
		}
	}
	return v1, v2
}
