package application

type EnvironmentMode string

const (
	Unknown     EnvironmentMode = "unknown"
	Development EnvironmentMode = "dev"
	Production  EnvironmentMode = "prod"
)
