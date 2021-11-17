package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func Migrate1636995600add_table(db *gorm.DB) *gormigrate.Gormigrate {
  return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
       {
            ID: "1636995600add_table",
            Migrate: func(tx *gorm.DB) error {
		   type Option string

		   const (
					Avatar	Option = "AVATAR"
					News 	Option = "NEWS"
				)

            type Image struct {
            	Id     		int			`gorm:"primaryKey" json:"id"`
            	Name   		string 		`json:"name"`
            	UserRefer 	int			`json:"user_refer"`
            	Option      Option		`json:"option" orm:"type:enum('AVATAR', 'NEWS'); default:'NEWS'"`
            	CreatedTime time.Time   `json:"created_time"`
            	UpdatedTime time.Time   `json:"updated_time"`
            }

		   type Ip struct {
			   Id     int          `gorm:"primary_key" json:"id"`
			   Ip     string       `json:"ip"`
			   Method string       `json:"method"`
			   RequestAt time.Time `json:"created_at"`
	       }

	       type Users struct {
	       		Id 		 		int 		`gorm:"primaryKey" json:"id"`
	       		Username 		string 		`json:"username"`
	       		Password 		string 		`json:"password"`
	       		Email    		string 		`json:"email"`
	       		Name     		string 		`json:"name"`
	       		Phone    		string 		`json:"phone"`
	       		Address  		string 		`json:"address"`
	       		Sex      		string 		`json:"sex"`
	       		Checked  		bool   		`json:"checked"`
	       		CreateTime 		time.Time	`json:"create_time"`
	       		UpdateTime 		time.Time	`json:"update_time" gorm:"default:null"`
	       		Image           Image		`gorm:"foreignKey:UserRefer"`
	       }

                  return tx.AutoMigrate(&Users{}, &Ip{}, &Image{})
          },
            Rollback: func(tx *gorm.DB) error {

                  return tx.Migrator().DropTable("users","ips","images")
			},
		  },
  })

}

