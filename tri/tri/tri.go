package tri

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.

# Info
2021-2022, test2, exercice 6
*/

func tri(tinit []int) (tfin []int) {
	for i:=0; i<len(tinit); i++{
		tfin = append(tfin,tinit[i])
	}
	if len(tfin)>1{
		var n int = len(tfin)
		var min int 
		for i:=0;i<n-1; i++ {
			min = i
			for j:=i+1; j<n; j++ {
				if tfin[min]>tfin[j] {
					min = j
				}
			}
			if min != i{
				tfin[min],tfin[i] = tfin[i], tfin[min]
			}
		}
	}
	return tfin
}
