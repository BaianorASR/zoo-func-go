package zoo

import (
	"strings"
	"zoologic/data"
)

// GetEmployeeByName é responsável pela busca das pessoas colaboradoras através do primeiro ou do último nome delas
func GetEmployeeByName(employeeName string) data.Employees {
	var output data.Employees

	// Pecorre o slice de employees para encontrar o funcionario que tenho o "firstName" ou "lastName" igual o employeeName, parametro da função
	for _, employee := range zoologic.Employees {
		if strings.EqualFold(employee.FirstName, employeeName) || strings.EqualFold(employee.LastName, employeeName) {
			output = employee
			break
		}
	}

	return output
}
