package task

import (
	"fmt"
	"todo/reciever/task"
)

func Show(t TaskPresenter) {
	fmt.Printf("### %s\n", t.present())
}

type TaskPresenter interface {
	present() string
}

type NormalTaskPresenter struct {
	receiver *task.NormalTaskReciever
}

type GoogleCalendarTaskPresenter struct {
	receiver *task.GoogleCalendarTaskReciever
}

func NewNormalTaskPresenter(r *task.NormalTaskReciever) *NormalTaskPresenter {
	return &NormalTaskPresenter{
		receiver: r,
	}
}

func (f NormalTaskPresenter) present() string {
	return fmt.Sprintf("[%s]%s\n%s\n",
		f.receiver.Status,
		f.receiver.Title,
		f.receiver.Url,
	)
}

func NewGoogleCalendarTaskPresenter(r *task.GoogleCalendarTaskReciever) *GoogleCalendarTaskPresenter {
	return &GoogleCalendarTaskPresenter{
		receiver: r,
	}
}

func (f GoogleCalendarTaskPresenter) present() string {
	fmt.Printf("%s", f.receiver.Status)
	return fmt.Sprintf("[%s]%s(%s)\n%s\n",
		f.receiver.Status,
		f.receiver.Title,
		f.receiver.Place,
		f.receiver.Url,
	)
}
