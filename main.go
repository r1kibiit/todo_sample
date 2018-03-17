package main

import (
	tPresenter "todo/presenter/task"
	tReceiver "todo/reciever/task"
	tRepo "todo/repository/task"
)

func main() {
	var p tPresenter.TaskPresenter
	tasks := tRepo.FindAll()
	for _, t := range tasks {
		switch {
		case t.IsGoogleCalendarTask():
			r := tReceiver.NewGoogleCalendarTaskReceiver(&t)
			p = tPresenter.NewGoogleCalendarTaskPresenter(r)
		default:
			r := tReceiver.NewNormalTaskReceiver(&t)
			p = tPresenter.NewNormalTaskPresenter(r)
		}
		tPresenter.Show(p)
	}
}
