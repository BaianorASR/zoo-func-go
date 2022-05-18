package zoo

import (
	"strings"
	"zoologic/data"
)

//Esta função, a partir do nome de uma espécie e uma idade mínima, verifica se todos os animais daquela espécie possuem a idade mínima especificada.
func GetAnimalsOlderThan(animal string, age int) bool {
	haveTheMinimumAge := true

	// Declara a variavel vazia para receber posteriormente o valor
	var SPECIE data.Species

	// Procura dentro do Slice o Struct da specie pelo nome do animal "animal func prop"
	for _, specie := range zoologic.Species {
		if specie.Name == strings.ToLower(animal) {
			SPECIE = specie
			break
		}
	}

	// Verifica se todos os animais daquela especie tem a idade minima "age func prop"
	for _, eachAnimal := range SPECIE.Residents {
		if eachAnimal.Age < age {
			haveTheMinimumAge = false
			break
		}
	}

	return haveTheMinimumAge
}
