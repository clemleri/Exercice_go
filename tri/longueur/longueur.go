package longueur

/*
On souhaite trier un tableau de chaînes de caractères (minuscules sans accents) de la plus longue à la plus courte (et par ordre alphabétique
dans le cas où plusieurs chaînes ont la même longueur). La fonction trier doit réaliser ce tri en place. On rappelle que l'opérateur < sur les chaînes
de caractères définit l'ordre alphabétique.

# Entrée
- tab : le tableau de chaînes de caractères à trier

# Info
2022-2023, test3, exercice 1
*/

func trier(tab []string) {
	var plus_long int 
	var n int = len(tab)
	for i:=0; i<n-1; i++ {
		plus_long = i
		for j:=i+1; j<n; j++ {
			if len(tab[plus_long])<len(tab[j]) {
				plus_long = j
			}else if len(tab[plus_long])==len(tab[j]) {
				for y:=0; y<len(tab[plus_long]);y++ {
					if tab[plus_long]==tab[j]{
						continue
					}else if tab[plus_long][y]<tab[j][y] {
						break
					}else if tab[plus_long][y]>tab[j][y] {
						plus_long = j
						break
					}
				}
			}
		}
		if plus_long!=i {
			tab[i],tab[plus_long] = tab[plus_long], tab[i]
		}
	}
}
