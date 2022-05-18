package zoo

import (
	"strings"
	"zoologic/data"
)

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

	// Verifica se todos os animais daquela especie tem a idade minima "age frun prop"
	for _, eachAnimal := range SPECIE.Residents {
		if eachAnimal.Age < age {
			haveTheMinimumAge = false
			break
		}
	}

	return haveTheMinimumAge
}
