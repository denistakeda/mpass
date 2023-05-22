package ports

import "github.com/rs/zerolog"

type (
	LogService interface {
		ComponentLogger(component string) zerolog.Logger
	}
)
