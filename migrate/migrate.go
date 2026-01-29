package migrate

import (
	"fmt"

	"github.com/Setsu548/App_texi_go_migrations/config"
	"github.com/Setsu548/App_texi_go_migrations/models"
	"github.com/Setsu548/App_texi_go_migrations/pkg/logger"
	"gorm.io/gorm"
)

type Migration struct {
	DB     *gorm.DB
	Config config.Config
}

func NewMigration(db *gorm.DB, cfg config.Config) *Migration {
	mg := &Migration{
		DB:     db,
		Config: cfg,
	}

	return mg

}

func (m *Migration) Migrate() error {
	logger.Info("Starting migration...")

	// Add your models here to be migrated
	modelsToMigrate := []interface{}{
		// User related tables
		&models.User{},
		&models.Staff{},
		&models.Employee{},
		&models.DocumentInfo{},

		// Core tables
		&models.TypeUser{},
		&models.DocumentType{},
		&models.Role{},
		&models.Permission{},

		// Geography tables (must be in order due to foreign keys)
		&models.Country{},
		&models.Department{},
		&models.Locality{},

		// Service and vehicle tables
		&models.ServiceType{},
		&models.Vehicle{},
		&models.ImageCar{},
		&models.Fare{},

		// Security and session tables
		&models.Session{},
		&models.PasswordReset{},
		&models.PhoneVerification{},

		// Access control tables
		&models.UserRole{},
		&models.RolePermission{},

		// Trip tables
		&models.Trip{},
	}

	for _, model := range modelsToMigrate {
		// Use HasTable to check if table already exists
		tableName := m.DB.NamingStrategy.TableName(fmt.Sprintf("%T", model))

		// AutoMigrate will handle table creation and column updates
		err := m.DB.AutoMigrate(model)
		if err != nil {
			logger.Error(err, "Migration failed for model", map[string]interface{}{"model": model, "table": tableName})
			return fmt.Errorf("migration error for %s: %v", tableName, err)
		}
		logger.Info("Migrated model successfully", map[string]interface{}{"model": model, "table": tableName})
	}

	// Add manual foreign key constraints for Staff and Employee after all tables exist
	logger.Info("Adding foreign key constraints...")

	// Check if staffs table exists before adding constraint
	if m.DB.Migrator().HasTable("staffs") {
		// Add foreign key for staffs.id -> users.id
		if !m.DB.Migrator().HasConstraint(&models.Staff{}, "fk_users_staff") {
			err := m.DB.Exec(`ALTER TABLE staffs ADD CONSTRAINT fk_users_staff FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE`).Error
			if err != nil {
				logger.Error(err, "Failed to add foreign key constraint for staffs", nil)
				return err
			}
			logger.Info("Added foreign key constraint for staffs -> users")
		}
	} else {
		logger.Info("Staffs table does not exist, skipping foreign key constraint")
	}

	// Check if employees table exists before adding constraint
	if m.DB.Migrator().HasTable("employees") {
		// Add foreign key for employees.id -> users.id
		if !m.DB.Migrator().HasConstraint(&models.Employee{}, "fk_users_employee") {
			err := m.DB.Exec(`ALTER TABLE employees ADD CONSTRAINT fk_users_employee FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE`).Error
			if err != nil {
				logger.Error(err, "Failed to add foreign key constraint for employees", nil)
				return err
			}
			logger.Info("Added foreign key constraint for employees -> users")
		}
	} else {
		logger.Info("Employees table does not exist, skipping foreign key constraint")
	}

	logger.Info("Migration completed successfully.")
	return nil
}
