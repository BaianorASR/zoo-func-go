package zoo

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type InfoDay struct {
  OfficeHour string
  Exhibition []string
}

type Monday struct {
    OfficeHour string
    Exhibition string
}

type Days struct {
  Tuesday, Wednesday, Thursday, Friday, Saturday InfoDay
  Monday 
}

func getOnlyMonday() map[string]Monday{
  return map[string]Monday{ "Monday": { OfficeHour: "CLOSED", Exhibition: "The zoo will be closed!" }}
}

func ScheduleByName(animalName string) (Availability []string) {
  for _, schedule := range zoologic.Species {
    if schedule.Name == animalName {
      Availability = schedule.Availability
      break
    }
  }
  fmt.Println(Availability)
  return
}

func ScheduleDays(dayOfWeek string) map[string]InfoDay {  
  var (
    Day = make(map[string]InfoDay)
    Exhibition = make([]string, 0)

    OfficeHour = fmt.Sprintf("Open from %dam until %dpm",
      zoologic.Hours[strings.ToLower(dayOfWeek)].Open,
      zoologic.Hours[strings.ToLower(dayOfWeek)].Close)
  )
    
    for _, species := range zoologic.Species {
      if slices.Contains(species.Availability, dayOfWeek) {
        Exhibition = append(Exhibition, species.Name)
      }
    }
    
    Day[dayOfWeek] = InfoDay{OfficeHour : OfficeHour, Exhibition: Exhibition}
    return Day
}

func ScheduleAllEntries() (OUTPUT []map[string]interface{})  {
  var (
    allDays = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    // OUTPUT = make([]map[string]interface{}, 0)
  )
  
  for _, day := range allDays {
    var (
      Exhibition = make([]string, 0)
      OfficeHour = fmt.Sprintf("Open from %dam until %dpm",
        zoologic.Hours[strings.ToLower(day)].Open,
        zoologic.Hours[strings.ToLower(day)].Close)
    )

    for _, species := range zoologic.Species {
      if slices.Contains(species.Availability, day) {
        Exhibition = append(Exhibition, species.Name)
      }
    }

    if day == "Monday" {
      OUTPUT = append(OUTPUT, map[string]interface{}{
        "Monday": struct{
            OfficeHour string
            Exhibition string
            }{ OfficeHour: "CLOSED", Exhibition: "The zoo will be closed!" },
          })
    } else {
      OUTPUT = append(OUTPUT, map[string]interface{}{
        day: InfoDay{OfficeHour: OfficeHour, Exhibition: Exhibition},
      })
    }
  }
  return
}

func GetSchedule(scheduleTarget string) any {
  var (
    allDays  = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    allAnimals []string
  )

  for _, species := range zoologic.Species {
    allAnimals = append(allAnimals, species.Name)
  }
  
  if slices.Contains(allDays, scheduleTarget) {
    if scheduleTarget == "Monday" {
      return getOnlyMonday()
    }
    return ScheduleDays(scheduleTarget)
  }

  if slices.Contains(allAnimals, scheduleTarget) {
    return ScheduleDays(scheduleTarget)
  }
  
  return ScheduleAllEntries()
}
