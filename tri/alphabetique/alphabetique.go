package alphabetique

/*
La fonction alphabetique trie un tableau de chaînes de caractères dans l'ordre
alphabétique.

# Entrée
- dico : le tableau de chaînes de caractères à trier
*/




func alphabetique(dico []string) {
	ordre := func(m1,m2 string) (m1avantm2 bool){
		if len(m1)==0 && len(m2)==0{
			m1avantm2=false
		}else if len(m1)==0 || len(m2)==0{
			if len(m1)==0{
				m1avantm2 = true
			}else{
				m1avantm2 =false
			}
		}else if m1[0]<m2[0] {
			m1avantm2 = true
		}else if m1[0]>m2[0] {
			m1avantm2 = false
		}else{
			for i:=0; i<len(m1);i++ {
				if m1[i]==m2[i]{
					if i+1==len(m1)&&i+1!=len(m2){
						m1avantm2 = true
						break
					}else if i+1==len(m2)&&i+1!=len(m1){
						m1avantm2 = false
						break
					}
					m1avantm2 = false
					continue
				}else if m1[i]<m2[i] {
					m1avantm2 = true
					break
				}else if m1[i]>m2[i] {
					m1avantm2 = false
					break
				}
			}
		}
		return m1avantm2
	}
		var mot_avant int 
		var n int = len(dico)
		var b_mot_avant bool 
		for i:=0; i<n-1;i++ {
			mot_avant = i
			for j:=i+1; j<n; j++ {
				b_mot_avant = ordre(dico[mot_avant],dico[j])
				if !b_mot_avant{
					mot_avant = j
				}
			}
			if i!=mot_avant{
				dico[i], dico[mot_avant] = dico[mot_avant], dico[i]
			}
		}
	}
	

