package models

type TeacherDay struct {
	WeekDay     int
	LessonNum   int
	Class       string
	Teacher     string
	Point       string
	Color       string
	Numerator   bool
	Denominator bool
	Error       string
}
