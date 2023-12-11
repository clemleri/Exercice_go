package lignes

import (
	"bufio"
	"os"
)

/*
La fonction lignes doit compter le nombre de lignes non vides dans un
fichier dont le nom est indiqué en paramètre.

# Entrée
- fName : une chaîne de caractères correspondant à un nom de fichier

# Sortie
- nLignes : un entier indiquant le nombre de lignes non vides dans le
fichier de nom fName ou -1 si le fichier n'existe pas
*/

func lignes(fName string) (nLignes int) {
	file, err := os.Open(fName)
	if err != nil {
		return -1
	}
	var result int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			result++
		}
	}
	return result
}
