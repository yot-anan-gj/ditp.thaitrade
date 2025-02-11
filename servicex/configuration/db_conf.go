package configuration

import "fmt"

const (
	//AWS RDS for PostgreSQL
	POSTGRES_AWS = "pg_aws"
	//GCP Cloud SQL for PostgreSQL
	POSTGRES_GCP = "pg_gcp"
	//On-Premise (or locally hosted)
	POSTGRES_ON_PREMISE = "pg_local"
	MYSQL_ON_PREMISE    = "mysql"
)

type DBConfig struct {
	ContextName             string
	Provider                string
	URL                     string
	User                    string
	Password                string
	DatabaseName            string
	CreateConnectionTimeout int
	InitialScripts          []string
}

func (db DBConfig) String() string {
	return fmt.Sprintf("ContextName: %s, "+
		"Provider: %s, "+
		"URL: %s, "+
		"User: %s, "+
		"Password: %s, "+
		"DatabaseName: %s, "+
		"CreateConnectionTimeout: %d, "+
		"InitialScripts: %v",
		db.ContextName, db.Provider,
		db.URL, db.User, db.Password, db.DatabaseName, db.CreateConnectionTimeout, db.InitialScripts)

}
