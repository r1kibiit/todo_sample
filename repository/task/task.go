package task

type Task struct {
	Type    string
	Status  string
	Title   string
	Url     string
	Place   string
	StartAt string
	EndAt   string
}

func (t Task) IsNormalTask() bool {
	return t.Type == normal
}

func (t Task) IsGoogleCalendarTask() bool {
	return t.Type == googleCalendar
}
