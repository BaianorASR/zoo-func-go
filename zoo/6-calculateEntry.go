package zoo

type Entrants struct {
	Name string
	Age int
}

type CountEntrantsReturn struct {
	Child int
	Adult int
	Senior int
}

// CountEntrants deverá receber o Struct de visitantes e retornar um objeto com a contagem de acordo com os seguintes critérios de classificação:
func CountEntrants(entrants []Entrants) (OUTPUT CountEntrantsReturn) {
	// Pecorre entrants para separar pela faixa de idade
	for _, entrant := range entrants {
		switch true {
			case entrant.Age > 49:
				OUTPUT.Senior += 1
			case entrant.Age < 18:
				OUTPUT.Child += 1
			default:
				OUTPUT.Adult += 1
		}
	}
	return
}

// CalculateEntry deverá receber um array de visitantes e a partir da quantidade de visitantes e faixa etária de cada um, deverá retornar o valor total a ser cobrado.
func CalculateEntry(entrants []Entrants) (totalValue float64) {
	// Pega o total de visitantes separados pela idade
	totalEntrants := CountEntrants(entrants)

	totalValue += float64(totalEntrants.Child) * zoologic.Prices.Child
	totalValue += float64(totalEntrants.Adult) * zoologic.Prices.Adult
	totalValue += float64(totalEntrants.Senior) * zoologic.Prices.Senior
	return
}
