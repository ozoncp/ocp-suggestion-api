package models

import "fmt"

//Suggestion хранит информацию о предложении курса пользователю
type Suggestion struct {
	Id       uint64
	UserId   uint64
	CourseId uint64
}

//NewSuggestion создает новую структуру типа Suggestion
func NewSuggestion(id, userid, courseid uint64) Suggestion {
	return Suggestion{Id: id, UserId: userid, CourseId: courseid}
}

//String имплементация метода для Suggestion
func (s Suggestion) String() string {
	return fmt.Sprintf(
		"Suggestion{Id: %v, UserId: %v, CourseId: %v}",
		s.Id,
		s.UserId,
		s.CourseId,
	)
}
