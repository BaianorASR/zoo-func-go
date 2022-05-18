package zoo

import (
	"zoologic/data"
)

// GetSpeciesByIds retorna um array filtrado com as Species que possuem o ID correspodente
func GetSpeciesByIds(ids ...string) []data.Species {
	array := make([]data.Species, 0, len(ids))

	// Pecorre o slice Species para comparar se o ID da specie é igual ao id passado para a função
	for _, specie := range zoologic.Species {
		for _, id := range ids {
			if specie.ID == id {
				array = append(array, specie)
			}
		}
	}

	return array
}
