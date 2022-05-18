package zoo

import (
	"zoologic/data"
)

// GetSpeciesByIds é responsável pela busca das espécies de animais por id. Ela retorna um array contendo as espécies referentes aos ids passados como parâmetro, podendo receber um ou mais ids.
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
