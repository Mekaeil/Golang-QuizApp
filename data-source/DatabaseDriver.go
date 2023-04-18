package data_source

type Enum string

const (
	MYSQL    Enum = "mysql"
	POSTGRES Enum = "postgres"
)

var drivers = []Enum{
	MYSQL,
	POSTGRES,
}

func GetDrivers() []Enum {
	return drivers
}
