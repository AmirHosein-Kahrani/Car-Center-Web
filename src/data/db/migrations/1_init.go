package migrations

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/constants"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var cfg = config.GetConfig()
var logger = logging.NewLogger(cfg)

func Up_1() {
	database := db.GetDb()
	createTables(database)
	createDefaultInformation(database)
	createCountry(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}
	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	database.Debug().Migrator().CreateTable(tables...)

	logger.Info(logging.Postgres, logging.Migration, "Tables Created", nil)
}

func addNewTable(database *gorm.DB, models interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(models) {
		tables = append(tables, models)
	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExist(database, &defaultRole)

	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u := models.User{UserName: constants.DefaultUserName, FirstName: "test",
		LastName: "test", MobileNumber: "09131232312",
		Password: string(hashedPassword), Email: "admin@admin.com"}

	createAdminUserIfNotExist(database, &u, adminRole.Id)

}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	exist := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exist)
	if exist == 0 {
		database.Create(&r)
	}
}

func createAdminUserIfNotExist(database *gorm.DB, u *models.User, roleId int) {
	exist := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.UserName).
		First(&exist)
	if exist == 0 {
		database.Create(&u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}
func createCountry(database *gorm.DB) {
	const countStarExp = "count(*)"
	count := 0
	database.
		Model(&models.Country{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {

		// iran
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}})
		// usa
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}})

		// germany
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}})

		// china
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}})
		// Italy
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}})
		// France
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}})

		// Japan
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Kyoto"},
		}})

		// Italy
	}
}

func Down_1() {

}
