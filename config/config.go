package config

type PostgreSQLConfiguration struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBOptions  string
	Locale     string
}
