package data

const  (
	lionId string = "0938aa23-f153-4937-9f88-4858b24d6bce"
	ottersId string = "533bebf3-6bbe-41d8-9cdf-46f7d13b62ae"
	elephantsId string = "bb2a76d8-5fe3-4d03-84b7-dba9cfc048b5"
	snakesId string = "78460a91-f4da-4dea-a469-86fd2b8ccc84"
	frogsId string = "89be95b3-47e4-4c5b-b687-1fabf2afa274"
	bearsId string = "baa6e93a-f295-44e7-8f70-2bcdc6f6948d"
	tigersId string = "e8481c1d-42ea-4610-8e11-1752cfc05a46"
	stephanieId string = "9e7d4524-363c-416a-8759-8aa7e50c0992"
	olaId string = "fdb2543b-5662-46a7-badc-93d960fdc0a8"
	burlId string = "0e7b460e-acf4-4e17-bcb3-ee472265db83"
)

type ZoologicStruct struct {
	Species   []Species   `json:"species"`
	Employees []Employees `json:"employees"`
	Hours     Hours       `json:"hours"`
	Prices    Prices      `json:"prices"`
}
type Residents struct {
	Name string `json:"name"`
	Sex  string `json:"Sex"`
	Age  int    `json:"Age"`
}
type Species struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Popularity   int         `json:"popularity"`
	Location     string      `json:"location"`
	Availability []string    `json:"availability"`
	Residents    []Residents `json:"residents"`
}
type Employees struct {
	ID             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

type Day struct {
	Open  int `json:"open"`
	Close int `json:"close"`
}

type Hours map[string]Day

type Prices struct {
	Adult  float64 `json:"adult"`
	Senior float64 `json:"senior"`
	Child  float64 `json:"child"`
}

var Zoologic = ZoologicStruct{
	Species: []Species{
		{
			ID:           lionId,
			Name:         "lions",
			Popularity:   4,
			Location:     "NE",
			Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
			Residents: []Residents{
				{
					Name: "Zena",
					Sex:  "female",
					Age:  12,
				},
				{
					Name: "Maxwell",
					Sex:  "male",
					Age:  15,
				},
				{
					Name: "Faustino",
					Sex:  "male",
					Age:  7,
				},
				{
					Name: "Dee",
					Sex:  "female",
					Age:  14,
				},
			},
		},
		{
			ID:           tigersId,
			Name:         "tigers",
			Popularity:   5,
			Location:     "NW",
			Availability: []string{"Wednesday", "Friday", "Saturday", "Tuesday"},
			Residents: []Residents{
				{
					Name: "Shu",
					Sex:  "female",
					Age:  19,
				},
				{
					Name: "Esther",
					Sex:  "female",
					Age:  17,
				},
			},
		},
		{
			ID:           bearsId,
			Name:         "bears",
			Popularity:   5,
			Location:     "NW",
			Availability: []string{"Tuesday", "Wednesday", "Sunday", "Saturday"},
			Residents: []Residents{
				{
					Name: "Hiram",
					Sex:  "male",
					Age:  4,
				},
				{
					Name: "Edwardo",
					Sex:  "male",
					Age:  4,
				},
				{
					Name: "Milan",
					Sex:  "male",
					Age:  4,
				},
			},
		},
		{
			ID:           "ef3778eb-2844-4c7c-b66c-f432073e1c6b",
			Name:         "penguins",
			Popularity:   4,
			Location:     "SE",
			Availability: []string{"Tuesday", "Wednesday", "Sunday", "Saturday"},
			Residents: []Residents{
				{
					Name: "Joe",
					Sex:  "male",
					Age:  10,
				},
				{
					Name: "Tad",
					Sex:  "male",
					Age:  12,
				},
				{
					Name: "Keri",
					Sex:  "female",
					Age:  2,
				},
				{
					Name: "Nicholas",
					Sex:  "male",
					Age:  2,
				},
			},
		},
		{
			ID:           ottersId,
			Name:         "otters",
			Popularity:   4,
			Location:     "SE",
			Availability: []string{"Friday", "Thursday", "Wednesday", "Saturday"},
			Residents: []Residents{
				{
					Name: "Neville",
					Sex:  "male",
					Age:  9,
				},
				{
					Name: "Lloyd",
					Sex:  "male",
					Age:  8,
				},
				{
					Name: "Mercedes",
					Sex:  "female",
					Age:  9,
				},
				{
					Name: "Margherita",
					Sex:  "female",
					Age:  10,
				},
			},
		},
		{
			ID:           frogsId,
			Name:         "frogs",
			Popularity:   2,
			Location:     "SW",
			Availability: []string{"Saturday", "Friday", "Thursday", "Wednesday"},
			Residents: []Residents{
				{
					Name: "Cathey",
					Sex:  "female",
					Age:  3,
				},
				{
					Name: "Annice",
					Sex:  "female",
					Age:  2,
				},
			},
		},
		{
			ID:           snakesId,
			Name:         "snakes",
			Popularity:   3,
			Location:     "SW",
			Availability: []string{"Sunday", "Saturday", "Friday", "Thursday"},
			Residents: []Residents{
				{
					Name: "Paulette",
					Sex:  "female",
					Age:  5,
				},
				{
					Name: "Bill",
					Sex:  "male",
					Age:  6,
				},
			},
		},
		{
			ID:           elephantsId,
			Name:         "elephants",
			Popularity:   5,
			Location:     "NW",
			Availability: []string{"Friday", "Saturday", "Sunday", "Tuesday"},
			Residents: []Residents{
				{
					Name: "Ilana",
					Sex:  "female",
					Age:  11,
				},
				{
					Name: "Orval",
					Sex:  "male",
					Age:  15,
				},
				{
					Name: "Bea",
					Sex:  "female",
					Age:  12,
				},
				{
					Name: "Jefferson",
					Sex:  "male",
					Age:  4,
				},
			},
		},
		{
			ID:           "01422318-ca2d-46b8-b66c-3e9e188244ed",
			Name:         "giraffes",
			Popularity:   4,
			Location:     "NE",
			Availability: []string{"Wednesday", "Thursday", "Tuesday", "Friday"},
			Residents: []Residents{
				{
					Name: "Gracia",
					Sex:  "female",
					Age:  11,
				},
				{
					Name: "Antone",
					Sex:  "male",
					Age:  9,
				},
				{
					Name: "Vicky",
					Sex:  "female",
					Age:  12,
				},
				{
					Name: "Clay",
					Sex:  "male",
					Age:  4,
				},
				{
					Name: "Arron",
					Sex:  "male",
					Age:  7,
				},
				{
					Name: "Bernard",
					Sex:  "male",
					Age:  6,
				},
			},
		},
	},
	Employees: []Employees{
		{
			ID:             "c5b83cb3-a451-49e2-ac45-ff3f54fbe7e1",
			FirstName:      "Nigel",
			LastName:       "Nelson",
			Managers:       []string{burlId, olaId},
			ResponsibleFor: []string{lionId, tigersId},
		},
		{
			ID:        burlId,
			FirstName: "Burl",
			LastName:  "Bethea",
			Managers:  []string{stephanieId},
			ResponsibleFor: []string{
				lionId,
				tigersId,
				bearsId,
				"ef3778eb-2844-4c7c-b66c-f432073e1c6b",
			},
		},
		{
			ID:        olaId,
			FirstName: "Ola",
			LastName:  "Orloff",
			Managers:  []string{stephanieId},
			ResponsibleFor: []string{
				ottersId,
				frogsId,
				snakesId,
				elephantsId,
			},
		},
		{
			ID:             "56d43ba3-a5a7-40f6-8dd7-cbb05082383f",
			FirstName:      "Wilburn",
			LastName:       "Wishart",
			Managers:       []string{burlId, olaId},
			ResponsibleFor: []string{snakesId, elephantsId},
		},
		{
			ID:        stephanieId,
			FirstName: "Stephanie",
			LastName:  "Strauss",
			Managers:  []string{},
			ResponsibleFor: []string{
				ottersId,
				"01422318-ca2d-46b8-b66c-3e9e188244ed",
			},
		},
		{
			ID:             "4b40a139-d4dc-4f09-822d-ec25e819a5ad",
			FirstName:      "Sharonda",
			LastName:       "Spry",
			Managers:       []string{burlId, olaId},
			ResponsibleFor: []string{ottersId, frogsId},
		},
		{
			ID:        "c1f50212-35a6-4ecd-8223-f835538526c2",
			FirstName: "Ardith",
			LastName:  "Azevado",
			Managers:  []string{"b0dc644a-5335-489b-8a2c-4e086c7819a2"},
			ResponsibleFor: []string{
				tigersId,
				bearsId,
			},
		},
		{
			ID:        "b0dc644a-5335-489b-8a2c-4e086c7819a2",
			FirstName: "Emery",
			LastName:  "Elser",
			Managers:  []string{stephanieId},
			ResponsibleFor: []string{
				lionId,
				bearsId,
				elephantsId,
			},
		},
	},
	Hours: Hours{
		"tuesday":   Day{Open: 8, Close: 6},
		"wednesday": Day{Open: 8, Close: 6},
		"thursday":  Day{Open: 10, Close: 8},
		"friday":    Day{Open: 10, Close: 8},
		"saturday":  Day{Open: 8, Close: 10},
		"sunday":    Day{Open: 8, Close: 8},
		"monday":    Day{Open: 0, Close: 0},
	},
	Prices: Prices{
		Adult:  49.99,
		Senior: 24.99,
		Child:  20.99,
	},
}
