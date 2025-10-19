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

	database.Migrator().CreateTable(tables...)

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

func Down_1() {

}
