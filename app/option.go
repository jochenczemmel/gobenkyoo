package app

// AppOption provides options for the App following
// the 'functional options' pattern.
type AppOption func(*App)

// WithLoader sets the loader to use.
func WithLoader(loader Loader) AppOption {
	return func(a *App) {
		a.loader = loader
	}
}

/*
// WithStorer sets the storer to use.
func WithStorer(storer Storer) AppOption {
	return func(a *App) {
		a.storer = storer
	}
}

// WithImporter sets the importer to use.
func WithImporter(importer Loader) AppOption {
	return func(a *App) {
		a.importer = importer
	}
}

// WithRunner sets the application type to execute.
func WithRunner(runner Runner) AppOption {
	return func(a *App) {
		a.runner = runner
	}
}
*/
