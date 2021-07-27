package models

import "fmt"

//Suggestion хранит информацию о предложении курса пользователю
type Suggestion struct {
	ID       uint64
	UserID   uint64
	CourseID uint64
}

//String имплементация метода для Suggestion
func (s Suggestion) String() string {
	return fmt.Sprintf(
		"Suggestion{ID: %d, UserID: %d, CourseID: %d}",
		s.ID,
		s.UserID,
		s.CourseID,
	)
}
