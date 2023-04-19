package config

type Enum string

const (
	PRODUCTION Enum = "production"
	STAGING    Enum = "staging"
	LOCAL      Enum = "local"
)

var environments = []Enum{
	PRODUCTION,
	STAGING,
	LOCAL,
}

func GetEnvironment() []Enum {
	return environments
}

const (
	DAILY Enum = "daily"
	ALL   Enum = "all"
)

var logChannels = []Enum{
	DAILY,
	ALL,
}

func GetLogChannels() []Enum {
	return logChannels
}
