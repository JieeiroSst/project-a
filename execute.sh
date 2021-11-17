create_migrate() {
timer=$(date -d "${orig}" +"%s")
name=$1

touch ./db/migration/$timer$name.go

cat << EOF > ./db/migration/$timer$name.go
package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func migrate$timer$name(db *gorm.DB) *gormigrate.Gormigrate {
  return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
       {
            ID: "$timer$name",
            Migrate: func(tx *gorm.DB) error {

                  return tx.AutoMigrate()
          },
            Rollback: func(tx *gorm.DB) error {

                  return tx.Migrator()
          },
		  },
  })

}

EOF
}

"$@"