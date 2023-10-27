package doublons4

import "fmt"

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement k fois chaque nombre compris entre 1 et n/k. On voudrait vérifier
cela, mais on ne connaît pas k. La fonction doublons va calculer k, si c'est
possible.

# Entrée
- tab : un tableau d'entiers

# Sortie
- k : un entier (s'il existe) tel que tab contient exactement k fois chaque
entiers de 1 à len(tab)/k
- ok : un booléen qui indique si k existe ou non
*/

func doublons(tab []int) (k int, ok bool) {
	if len(tab) == 0{
		return k, true
	}
	for y:=0; y<len(tab);y++ {
		if tab[0]==tab[y] {
			k++
		}
	}
	fmt.Println(k)
	for i:=0;i<len(tab)/k;i++ {
		cpt:=0
		for j:=0; j<len(tab); j++{
			if i==tab[j] {
				fmt.Println("i est bien égale tab[j]")
				cpt++
			}
		}
		if cpt!=k{
			fmt.Println("cpt: ",cpt)
			return k, false
		}
	}
	fmt.Println("k return: ",k )
	return k, true
}
