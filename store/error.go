package store

type ConfigurationError string

func (c ConfigurationError) Error() string {
	return string(c)
}
