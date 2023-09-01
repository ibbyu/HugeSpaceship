package migration

import (
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//go:embed migrations
var migrations embed.FS

func ListMigrations() []string {
	migrations, err := migrations.ReadDir("migrations")
	if err != nil {
		return nil
	}
	migrationNames := make([]string, len(migrations))
	for i, migration := range migrations {
		if !strings.HasSuffix(migration.Name(), ".sql") {
			continue
		}
		migrationNames[i] = migration.Name()
	}
	return migrationNames
}

func GetMigrationByNumber(number int) (string, bool, error) {
	for _, file := range ListMigrations() {
		if strings.HasPrefix(file, fmt.Sprintf("%03d", number)) {
			return GetMigration(file)
		}
	}
	return "", false, errors.New("no migrations found for id " + strconv.Itoa(number))
}

func GetMigration(name string) (string, bool, error) {
	migration, err := migrations.ReadFile("migrations/" + name)
	if err != nil {
		return "", true, err
	}

	return string(migration), true, nil
}
