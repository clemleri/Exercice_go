package vraifaux

/*
On souhaite trier un tableau de booléens en mettant ceux dont la valeur est true au début et ceux dont la valeur est false à la fin. La fonction trier doit réaliser ce tri en place.

# Entrée
- tab : le tableau de booléens à trier

# Info
2022-2023, test3, exercice 0
*/

func trier(tab []bool) {
	if len(tab)>1{
		var n int = len(tab)
		var index_f int 
		for i:=0;i<n-1; i++ {
			index_f = i
			for j:=i+1; j<n; j++ {
				if !tab[index_f]{
					index_f = j
				}
			}
			if index_f != i{
				tab[index_f],tab[i] = tab[i], tab[index_f]
			}
		}
	}
}
