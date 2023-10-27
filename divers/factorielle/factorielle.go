package factorielle



/*
La fonction factorielle calcule la factorielle d'un nombre entier positif ou
nul. Pour rappel : 0! = 1 et pour n > 0, n! = n*(n-1)*(n-2)*...*1.

# Entrée
- n : l'entier dont on veut calculer la factorielle

# Sortie
- fact : la factorielle de n

# Exemple
factorielle(5) = 120
*/
func factorielle(n uint) (fact uint) {
	if n == 0 {
        return 1 
    }
    fact = uint(1)
    for i := uint(1); i <= n; i++ {
        fact *= i
    }
    return fact
}

