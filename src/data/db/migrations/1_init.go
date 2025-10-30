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

const countStarExp = "count(*)"

func Up_1() {
	database := db.GetDb()

	createTables(database)
	createDefaultInformation(database)
	createCountry(database)

	createPropertyCategory(database)

	createCarType(database)
	createGearbox(database)

}

func createTables(database *gorm.DB) {

	// tables
	tables := []interface{}{}

	tables = addNewTable(database, models.Country{}, tables)
	tables = addNewTable(database, models.City{}, tables)
	tables = addNewTable(database, models.File{}, tables)
	tables = addNewTable(database, models.PersianYear{}, tables)
	// Property
	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	// Car
	tables = addNewTable(database, models.Company{}, tables)
	tables = addNewTable(database, models.Gearbox{}, tables)
	tables = addNewTable(database, models.Color{}, tables)
	tables = addNewTable(database, models.CarType{}, tables)

	tables = addNewTable(database, models.CarModel{}, tables)
	tables = addNewTable(database, models.CarModelColor{}, tables)
	tables = addNewTable(database, models.CarModelYear{}, tables)
	tables = addNewTable(database, models.CarModelImage{}, tables)
	tables = addNewTable(database, models.CarModelPriceHistory{}, tables)
	tables = addNewTable(database, models.CarModelProperty{}, tables)
	// ! switch debug mode in production
	err := database.Debug().Migrator().CreateTable(tables...)

	if err != nil {
		panic(err)
	}

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
		Where("user_name = ?", u.UserName).
		First(&exist)
	if exist == 0 {
		database.Create(&u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createCountry(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Country{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}, Companies: []models.Company{
			{Name: "Saipa"},
			{Name: "Iran khodro"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}, Companies: []models.Company{
			{Name: "Tesla"},
			{Name: "Jeep"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}, Companies: []models.Company{
			{Name: "Opel"},
			{Name: "Benz"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}, Companies: []models.Company{
			{Name: "Chery"},
			{Name: "Geely"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}, Companies: []models.Company{
			{Name: "Ferrari"},
			{Name: "Fiat"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}, Companies: []models.Company{
			{Name: "Renault"},
			{Name: "Bugatti"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}, Companies: []models.Company{
			{Name: "Toyota"},
			{Name: "Honda"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}, Companies: []models.Company{
			{Name: "Kia"},
			{Name: "Hyundai"},
		}})
	}
}

func Down_1() {

}

func createPropertyCategory(database *gorm.DB) {
	count := 0

	database.
		Model(&models.PropertyCategory{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})
		database.Create(&models.PropertyCategory{Name: "Engine"})
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})
		database.Create(&models.PropertyCategory{Name: "Suspension"})
		database.Create(&models.PropertyCategory{Name: "Equipment"})
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})
		database.Create(&models.PropertyCategory{Name: "Lights"})
		database.Create(&models.PropertyCategory{Name: "Multimedia"})
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"})
		database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})
	}
	createProperty(database, "Body")
	createProperty(database, "Engine")
	createProperty(database, "Drivetrain")
	createProperty(database, "Suspension")
	createProperty(database, "Comfort")
	createProperty(database, "Driver support systems")
	createProperty(database, "Lights")
	createProperty(database, "Multimedia")
	createProperty(database, "Safety equipment")
	createProperty(database, "Seats and steering wheel")
	createProperty(database, "Windows and mirrors")

}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.
		Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select(countStarExp).
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}
	var props *[]models.Property
	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)

	case "Engine":
		props = getEngineProperties(catModel.Id)

	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)

	case "Suspension":
		props = getSuspensionProperties(catModel.Id)

	case "Comfort":
		props = getComfortProperties(catModel.Id)

	case "Driver support systems":
		props = getDriverSupportSystemProperties(catModel.Id)

	case "Lights":
		props = getLightsProperties(catModel.Id)

	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)

	case "Safety equipment":
		props = getSafetyEquipmentProperties(catModel.Id)

	case "Seats and steering wheel":
		props = getSeatsProperties(catModel.Id)

	case "Windows and mirrors":
		props = getWindowsProperties(catModel.Id)

	default:
		props = &([]models.Property{})
	}

	for _, prop := range *props {
		database.Create(&prop)
	}
}

func createCarType(database *gorm.DB) {
	count := 0
	database.
		Model(&models.CarType{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.CarType{Name: "Crossover"})
		database.Create(&models.CarType{Name: "Sedan"})
		database.Create(&models.CarType{Name: "Sports"})
		database.Create(&models.CarType{Name: "Coupe"})
		database.Create(&models.CarType{Name: "Hatchback"})
	}
}

func createGearbox(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Gearbox{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Gearbox{Name: "Manual"})
		database.Create(&models.Gearbox{Name: "Automatic"})
	}
}
