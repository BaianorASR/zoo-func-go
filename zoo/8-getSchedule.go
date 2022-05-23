package zoo

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Exhibition interface {
	string | []string
}

type InfoDay[T Exhibition] struct {
	OfficeHour string
	Exhibition T
}

// Retorna o map falado que o zoologic não está operando se o dia for "Monday"
func getOnlyMonday() map[string]InfoDay[string] {
	return map[string]InfoDay[string]{"Monday": {OfficeHour: "CLOSED", Exhibition: "The zoo will be closed!"}}
}

// Retorna todos os dias em que o anaimal esta disponivel no zoologico
func scheduleByName(animalName string) (Availability []string) {
	// Pecorre as "Species" para encontrar o animal correto comparando o nome
	for _, schedule := range zoologic.Species {
		if schedule.Name == animalName {
			Availability = schedule.Availability
			break
		}
	}
	return
}

func scheduleDays(dayOfWeek string) (OneDay map[string]interface{}) {
	// Salva todos os dados em AllEntries
	AllEntries := scheduleAllEntries()

	// Pega em "AllEntries" a chave correspodente ao dia
	for _, day := range AllEntries {
		if err := day[dayOfWeek]; err != nil {
			OneDay = day
		}
	}
	return
}

func scheduleAllEntries() (AllEntries []map[string]interface{}) {
	// Peccorre "daysOfTheWeek" para tereee o dia como base para comparações
	for _, day := range daysOfTheWeek {
		var (
			Exhibition = make([]string, 0)
			OfficeHour = fmt.Sprintf("Open from %dam until %dpm",
				zoologic.Hours[strings.ToLower(day)].Open,
				zoologic.Hours[strings.ToLower(day)].Close)
		)

		// Pecorre as "Species" para encontrar o animal correto comparando o nome
		for _, species := range zoologic.Species {
			if slices.Contains(species.Availability, day) {
				Exhibition = append(Exhibition, species.Name)
			}
		}

		// Se o dia for "Monday" deve informar que o zoologico está fechado, se não, mostra o horario de funcionamento e os animais diponiveis do dia
		if day == "Monday" {
			AllEntries = append(AllEntries, map[string]interface{}{
				"Monday": struct {
					OfficeHour string
					Exhibition string
				}{OfficeHour: "CLOSED", Exhibition: "The zoo will be closed!"},
			})
		} else {
			AllEntries = append(AllEntries, map[string]interface{}{
				day: InfoDay[[]string]{OfficeHour: OfficeHour, Exhibition: Exhibition},
			})
		}
	}
	return
}

// GetSchedule é responsável por disponibilizar as informações de horário dos animais em uma consulta para o usuário, que pode querer ter acesso ao cronograma da semana, de um dia ou de um animal em específico.
func GetSchedule(scheduleTarget string) any {
	// Se privido o nome da specie ira retornar os diasa em que a especie esta disponíveis
	for _, species := range zoologic.Species {
		if species.Name == scheduleTarget {
			return scheduleByName(scheduleTarget)
		}
	}

	// Se privido um dia da semana ira retornar o animais disponíveis do dia
	if slices.Contains(daysOfTheWeek, scheduleTarget) {
		if scheduleTarget == "Monday" {
			return getOnlyMonday()
		}
		return scheduleDays(scheduleTarget)
	}

	// Caso "scheduleTarget" não seja o nome de uma specie ou dia ira retornar todos os dias e os animais daquele dia
	return scheduleAllEntries()
}
