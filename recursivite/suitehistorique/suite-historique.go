package suitehist

/*
On considère la suite définie par
- U(0) = 0
- U(n) est le double du nombre de termes venant avant U(n) (donc U(n-1), U(n-2), etc jusqu'à U(0)) qui sont plus petits que n

La fonction terme doit permettre d'obtenir le n-ième terme de cette suite.

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite

# Exemples
terme(0) = 0
terme(1) = 2
terme(2) = 2
terme(3) = 6
terme(4) = 6
terme(5) = 6
terme(6) = 6
terme(7) = 14
terme(15) = 30
# Info
2022-2023, test 2, exercice 7
*/

func terme(n int) (un int) {
	if n==0 {
		return 0
	}else {
		if terme(n-1) < n{
			return 2*n
		}else {
			return terme(n-1)
		}	 
	}
}
