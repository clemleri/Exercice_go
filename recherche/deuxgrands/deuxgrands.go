package deuxGrands

/*
La fonction lesDeuxPlusGrands retourne les deux plus grands entiers présents dans un tableau

# Entrée
- t : un tableau d'entiers qui en contient au moins 2

# Sorties
- v1 : le plus grand entier dans t
- v2 : le deuxième plus grand entier dans t

# Info
2022-2023, test 1, exercice 8
*/

func lesDeuxPlusGrands(t []int) (v1, v2 int) {
	if len(t)==2 {
		if t[0]<=t[1] {
			return t[1], t[0]
		}else {
			return t[0], t[1]
		}
	}else {
		indice_max := 0
		for i:=0; i<len(t); i++ {
			if v1<=t[i]{
				v1, v2 = t[i], v1
				indice_max = i
			}
			if v2<t[i] && indice_max!=i{
				v2 = t[i]
			}
		}
	}
	return v1, v2
}

