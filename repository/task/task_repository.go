package task

const (
	normal         = "normal"
	googleCalendar = "google_calendar"
)

func FindAll() []Task {
	var tasks []Task
	tasks = append(tasks, Task{
		Type:   "normal",
		Status: "未着手",
		Title:  "◯◯についての調査",
		Url:    "https:google.com",
	}, Task{
		Type:   "google_calendar",
		Status: "完了",
		Title:  "C仕様打ち合わせ",
		Place:  "会議室",
		Url:    "https:google.com",
	})
	return tasks
}
