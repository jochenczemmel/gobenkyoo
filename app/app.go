// Package app provides entry functions to the application.
package app

import "fmt"

type AppOption func(*App)

/*
* TODO: Options definieren
func WithLoader(storeloader StoreLoader) AppOption {
	return func(a *App) {
		a.Loader = storeloader
		a.Storer = storeloader
	}
}
*/

// App provides access to the application
type App struct{}

// New returns a configured App object.
func New(opts ...AppOption) *App {
	result := &App{}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

// Load loads the application data from the storage
func (a *App) Load() error {
	return nil
}

// Run executes the application.
func (a *App) Run() error {
	// TODO: call real code
	// this version is only for setup system test

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
		fmt.Println("saved")
	}

	return nil
}
