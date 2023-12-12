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
	var t_trie []int = make([]int, len(tinit), len(tinit))
	copy(t_trie, tinit)
	for i := 0; i < len(t_trie)-1; i++ {
		min := i
		for j := i; j < len(t_trie); j++ {
			if t_trie[j] < t_trie[min] {
				min = j
			}
		}
		if i != min {
			t_trie[i], t_trie[min] = t_trie[min], t_trie[i]
		}
	}
	return t_trie
}
