package dto

type Grade struct {
	ID           string  `json:"id"`
	AssignmentID string  `json:"assignment_id"`
	TeacherID    string  `json:"teacher_id"`
	Score        float64 `json:"score"`
	Feedback     string  `json:"feedback"`
}

type ResponseGrade struct {
	ID           string  `json:"id"`
	AssignmentID string  `json:"assignment_id"`
	TeacherID    string  `json:"teacher_id"`
	Score        float64 `json:"score"`
	Feedback     string  `json:"feedback"`
	Assignment   struct {
		ID        string `json:"id"`
		Subject   string `json:"subject"`
		Tittle    string `json:"tittle"`
		StudentID string `json:"student_id"`
		Content   string `json:"content"`
	} `json:"assignment"`
}
