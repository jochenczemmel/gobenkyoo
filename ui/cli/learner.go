package cli

import (
	"fmt"

	"github.com/jochenczemmel/gobenkyoo/app"
)

type Learner struct {
	application *app.App
}

func NewLearner(application *app.App) *Learner {
	return &Learner{application: application}
}

func (l *Learner) Run() error {
	err := l.application.LoadBoxes()
	if err != nil {
		return fmt.Errorf("load application: %w", err)
	}

	fmt.Printf("Q: world\nA: ")
	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "世界" {
		fmt.Println("ok")
	} else {
		fmt.Println("wrong")
	}
	fmt.Print("continue (y/n): ")
	fmt.Scanf("%s", &answer)
	if answer == "y" {
		return fmt.Errorf("not implemented")
	}

	fmt.Print("save answer (y/n): ")
	fmt.Scanf("%s", &answer)
	if answer == "y" || answer == "Y" {
		// TODO: store results
		fmt.Println("saved")
	}

	return nil
}
