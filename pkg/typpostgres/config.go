package typpostgres

var (
	// DefaultConfigName is default lookup key for postgres configuration
	DefaultConfigName = "PG"

	// DefaultDBName is default value for dbName
	DefaultDBName = "sample"

	// DefaultUser is default value for user
	DefaultUser = "postgres"

	// DefaultPassword is default value for password
	DefaultPassword = "pgpass"

	// DefaultHost is default value for host
	DefaultHost = "localhost"

	// DefaultPort is default value for port
	DefaultPort = 5432

	// DefaultConfig for postgres
	DefaultConfig = &Config{
		DBName:   DefaultDBName,
		User:     DefaultUser,
		Password: DefaultPassword,
		Host:     DefaultHost,
		Port:     DefaultPort,
	}
)

// Config is postgres configuration
type Config struct {
	DBName   string `required:"true"`
	User     string `required:"true" default:"postgres"`
	Password string `required:"true" default:"pgpass"`
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
}
