package cli

import (
	"github.com/jochenczemmel/gobenkyoo/app"
)

type Learner struct {
	application *app.App
}

func NewLearner(application *app.App) *Learner {
	return &Learner{application: application}
}

func (l *Learner) Run() error {
	return nil
}
