package model

type ReadingStatus string

const (
	NotStarted ReadingStatus = "Not Started"
	InProgres  ReadingStatus = "In Progress"
	Completed  ReadingStatus = "Completed"
)

type Book struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	Author        string        `json:"author"`
	ReadingStatus ReadingStatus `json:"reading_status"`
}
