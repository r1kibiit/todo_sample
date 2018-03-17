package task

import "todo/repository/task"

type NormalTaskReciever struct {
	Status  string
	Title   string
	Url     string
	StartAt string
	EndAt   string
}

func NewNormalTaskReceiver(t *task.Task) *NormalTaskReciever {
	return &NormalTaskReciever{
		t.Status,
		t.Title,
		t.Url,
		t.StartAt,
		t.EndAt,
	}
}

type GoogleCalendarTaskReciever struct {
	Status  string
	Title   string
	Url     string
	Place   string
	StartAt string
	EndAt   string
}

func NewGoogleCalendarTaskReceiver(t *task.Task) *GoogleCalendarTaskReciever {
	return &GoogleCalendarTaskReciever{
		t.Status,
		t.Title,
		t.Url,
		t.Place,
		t.StartAt,
		t.EndAt,
	}
}
