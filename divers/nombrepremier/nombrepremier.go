package nombrepremier

/*
La fonction estPremier doit indiquer  par un booléen si un nombre est premier
ou pas

# Entrée
- n : un nombre entier

# Sortie
- b : un booléen indiquant si n est premier ou pas

# Exemples
estPremier(5) = true
estPremier(10) = false
*/
func estPremier(n int) (b bool) {
	var lst []int
	if n==0{
		b = false
	}else if n > 1 {
		for i:=1; i<=n; i++ {
			if n%i==0{
				lst = append(lst,i)
			}
		}
		if len(lst)>=2 {
			if lst[0]==1 && lst[1]==n {
			b = true
			}
		}
	}
	return b
}
