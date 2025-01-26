package dto

type Grade struct {
	ID           string  `json:"id"`
	AssignmentID string  `json:"assignment_id"`
	TeacherID    string  `json:"teacher_id"`
	Score        float64 `json:"score"`
	Feedback     string  `json:"feedback"`
}

type ResponseGrade struct {
	ID                  string  `json:"id"`
	AssignmentID        string  `json:"assignment_id"`
	TeacherID           string  `json:"teacher_id"`
	Score               float64 `json:"score"`
	Feedback            string  `json:"feedback"`
	AssignmentSubject   string  `json:"assignment_subject"`
	AssignmentTittle    string  `json:"assignment_tittle"`
	AssignmentStudentID string  `json:"assignment_student_id"`
	AssignmentContent   string  `json:"assignment_content"`
	AssignmentStatus    string  `json:"assignment_status"`
}
