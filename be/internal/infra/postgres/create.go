package postgres

func createDBIfNotExists(config Config) error {
	if config.DefaultDBName == "" {
		config.DBName = "postgres"
	} else {
		config.DBName = config.DefaultDBName
	}

	db, err := connect(config)
	if err != nil {
		return err
	}

	var exists bool
	err = db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", config.DBName)
	if err != nil {
		return err
	}

	if !exists {
		_, err = db.Exec("CREATE DATABASE $1", config.DBName)
		if err != nil {
			return err
		}
	}

	return db.Close()
}
