package zoo

import (
	"log"
)

// isManager será responsável por verificar se uma pessoa colaboradora é gerente ou não. O retorno dessa função deve ser um booleano: true ou false
func isManager(id string) (isManagerBool bool) {
	// Pecorre o Employees para ver se o colaborado tem o ID igual ao que foi passado para a função
	for _, employee := range zoologic.Employees {
		for _, managerID := range employee.Managers {
			if managerID == id {
				isManagerBool = true
				break
			}
		}
	}
	return
}

// GetRelatedEmployees se for uma pessoa colaboradora gerente, deve retornar um array contendo os nomes das pessoas colaboradoras que ela é responsável. se não for uma pessoa colaboradora gerente, deverá ser lançado um erro
func GetRelatedEmployees(managerId string) (relatedEmployees []string) {
	// Lança error se o managerId não for de um colaborador
	if bool := isManager(managerId); !bool {
		log.Fatal("O id inserido não é de uma pessoa colaboradora gerente!")
	}

	// Pecorre Employess para pega o fullName de todos os colaboradores que são gerenciado pelo genrente com ID passado para a função
	for _, employee := range zoologic.Employees {
		for _, ID := range employee.Managers {
			if ID == managerId {
				fullName := employee.FirstName + " " + employee.LastName
				relatedEmployees = append(relatedEmployees, fullName)
			}
		}
	}
	return
}
