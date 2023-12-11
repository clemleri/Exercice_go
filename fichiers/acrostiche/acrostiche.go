package acrostiche

import (
	"bufio"
	"os"
)

/*
La fonction acrostiche doit reconstituer le mot formé par le premier caractère
de chaque ligne d'un fichier, en ignorant les lignes vides.

# Entrée
- fName : le nom d'un fichier

# Sortie
- mot : la chaîne de caractère obtenue en mettant l'une après l'autre, dans
        l'ordre des lignes du fichier, les premières lettres de chacunes des
        lignes du fichier dont le nom est fName (on considère que ce fichier
        existe toujours)
*/

func acrostiche(fName string) (mot string) {
	file, err := os.Open(fName)
	if err != nil {
		return
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			result = result + string(line[0])
		}
	}
	return result
}
