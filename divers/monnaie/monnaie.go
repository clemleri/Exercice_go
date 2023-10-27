package monnaie

import (
	"errors"
)

var errPasAssezPaye error = errors.New("le montant payé est inférieur au montant de l'achat, impossible de rendre la monnaie")

/*
Étant donnés un montant d'achat en euros et un montant payé en euros,
la fonction rendreMonnaie retourne la somme qu'il faut rendre au client.

# Entrées
- eurosAchat : la partie entière du montant d'achat en euros
- centimesAchat : la partie décimale du montant d'achat (donc les centimes)
- eurosPayes : la partie entière du montant payé en euros
- centimesPayes : la partie décimale du montant payé (donc les centimes)

# Sorties
- eurosRendus : la partie entière du montant à rendre en euros
- centimesRendus : la partie décimale du montant à rendre (donc les centimes)
- err : une erreur qui doit être errPasAssezPaye si le montant payé est inférieur au montant de l'achat et nil sinon

# Exemple
rendreMonnaie(12, 20, 15, 0) = 2, 80
endreMonnaie(5, 50, 6, 70) = 1, 20
rendreMonnaie(5, 50, 6, 20) = 0, 70
rendreMonnaie(5, 50, 6, 30) = 0, 80
rendreMonnaie(5, 50, 6, 40) = 0, 90
*/
func rendreMonnaie(eurosAchat, centimesAchat, eurosPayes, centimesPayes int) (eurosRendus, centimesRendus int, err error) {
	total_achat := eurosAchat*100 + centimesAchat
	total_payes := eurosPayes*100 + centimesPayes

	if total_payes<total_achat {
		err = errPasAssezPaye
	}else {
		total_rendus := total_payes-total_achat
		
		eurosRendus = total_rendus/100
		centimesRendus = total_rendus%100
		err = nil
	}
	return eurosRendus, centimesRendus, err
}
