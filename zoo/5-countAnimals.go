package zoo

import (
	"zoologic/data"
)

type Animal struct {
	Specie string
	Sex string
}

// Declara o valor globalbalmente para todas funções terem acesso
var COUNTER int

// getAllAnimalsLenght retorna um map com todos os animais e sua quantidade de residentes
func getAllAnimalsLenght() map[string]int {
	// Faz um make de um map vazio para depois adicionar os animais formatado
	allAnimalsLength := make(map[string]int)
	
	// Pecorre o Slice de Species para captura o nome do animal e a sua quantidade e adicionar no map
	for _, specie := range zoologic.Species {
		allAnimalsLength[specie.Name] = len(specie.Residents)
	}
	
	return allAnimalsLength
}

// filterBySex ira pecorre os residentes da specie alvo para aumentar o COUNTER somente pela quantidade de residentes do sexo alvo
func filterBySex(specie data.Species, sex string) {
	for _, residents := range specie.Residents {
		if residents.Sex == sex {
			COUNTER += 1
		}
	} 
}

func CountAnimals(animal *Animal) interface{} {
	// Verifica primeiramente se não foi passado algum parametro, depois se o nome da specie está vazio para retornar um map com o nome: quantidade de todos animais do zoologico
	if  animal == nil || len(animal.Specie) == 0 {
		allAnimalsLength := getAllAnimalsLenght()
		return allAnimalsLength
	}

	// Pecorre o Slice de Species para pegarr o tamanho do slice de residentes, se o parametro Sex foi mandado irá chamar a função "filterBySex" para cuidar do resultado
	for _, specie := range zoologic.Species {
		if specie.Name == animal.Specie {
			if len(animal.Sex) > 0 {
				filterBySex(specie, animal.Sex)
				break
			}
			COUNTER = len(specie.Residents)
			break
		}
	}

	return COUNTER
}
