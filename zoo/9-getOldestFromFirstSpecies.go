package zoo

import (
	"zoologic/data"
)

// GetOldestFromFirstSpecies busca por informações do animal mais velho da primeira espécie gerenciada pela pessoa colaboradora do parâmetro.
func GetOldestFromFirstSpecies(id string) []interface{} {
  // Declara variaveis auxiliares
  var firstAnimalId string
  var mostOlder data.Residents
  
  // Pecorre os Employees para filtrar pelo id. E captura o id do primeiro animal pelo qual é responsavel 
  for _, employee := range zoologic.Employees {
      if employee.ID == id {
        firstAnimalId = employee.ResponsibleFor[0]
        break
      }
  }

  // Pecorre as Species para encontrar o a specie pelo ID e então filtrar pelo residente mais velho
  for _, specie := range zoologic.Species {
    if specie.ID == firstAnimalId {
      mostOlder =specie.Residents[0]
      for _, resident := range specie.Residents {
        if resident.Age > mostOlder.Age {
          mostOlder = resident
        }
      }
      break
    }
  }

  return []interface{}{mostOlder.Name, mostOlder.Sex, mostOlder.Age}
}
