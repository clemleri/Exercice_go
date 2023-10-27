package matrice

import (
	"errors"
)

var errPasMat error = errors.New("mat n'est pas vraiment une matrice")

/*
La fonction compte indiquera combien de fois un nombre n apparaît dans une matrice.

# Entrées
- n : le nombre à chercher
- mat : une matrice d'entiers (un tableau de tableaux qui font tous la même taille)

# Sorties
- num : le nombre de fois que n apparaît dans mat
- err : errPasMat si mat n'est pas une vraie matrice (toutes les lignes n'ont pas
la même longueur ou mat vaut nil ou une ligne vaut nil)

# Exemple
compte(0, [][]int{[]int{0, 1}, []int{0, 0}}) = 3
*/

func compte(n int, mat [][]int) (num int, err error) {
	if mat == nil {
		return num, errPasMat
	}else if len(mat) ==0{
		return num, err
	}
	longueur := len(mat[0])
	for i:=0; i<len(mat); i++ {
		if len(mat[i])!=longueur || mat[i]==nil{
			err = errPasMat
		}
		for j:=0; j<len(mat[i]); j++{
			if mat[i][j] == n {
				num ++
			}
		} 
	}
	return num, err 
}