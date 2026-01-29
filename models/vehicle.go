package models

import "time"

// ServiceType representa la tabla service_types
type ServiceType struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:100;uniqueIndex:uk_service_name;not null" json:"name"`
	Description string     `gorm:"size:255" json:"description"`
	Status      string     `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt   CustomTime `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy   *int64     `json:"created_by"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	DeletedBy   *int64     `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (ServiceType) TableName() string {
	return "service_types"
}

// Vehicle representa la tabla vehicles
type Vehicle struct {
	ID              int64        `gorm:"primaryKey" json:"id"`
	UUID            string       `gorm:"size:36;uniqueIndex:idx_vehicles_uuid;not null" json:"uuid"`
	UserID          int64        `gorm:"not null;foreignKey:ID" json:"user_id"`
	ServiceTypeID   int64        `gorm:"not null;foreignKey:ID" json:"service_type_id"`
	Brand           string       `gorm:"size:100;not null" json:"brand"`
	Year            int64        `gorm:"not null" json:"year"`
	Model           string       `gorm:"size:100;not null" json:"model"`
	LicensePlate    string       `gorm:"size:20;uniqueIndex:idx_vehicles_license_plate;not null" json:"license_plate"`
	Color           string       `gorm:"size:50;not null" json:"color"`
	TittleDeed      string       `gorm:"size:500;not null" json:"tittle_deed"`
	UsedBy          *int64       `json:"used_by" example:"1" description:"ID del usuario que está usando el vehículo actualmente"`
	VIN             string       `gorm:"size:50;not null" json:"vin"`
	InsurancePolicy string       `gorm:"size:100;not null" json:"insurance_policy"`
	ImageCars       ImageCars    `gorm:"foreignKey:VehicleID;references:ID" json:"image_cars,omitempty"`
	User            *User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	ServiceType     *ServiceType `gorm:"foreignKey:ServiceTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"service_type,omitempty"`
	CreatedAt       time.Time    `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy       int64        `gorm:"not null" json:"created_by"`
	DeletedAt       *time.Time   `json:"deleted_at,omitempty"`
	DeletedBy       *int64       `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Vehicle) TableName() string {
	return "vehicles"
}

type ImageCar struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	VehicleID int64      `gorm:"not null;foreignKey:ID" json:"vehicle_id"`
	ImageName string     `gorm:"size:50;not null" json:"image_name"`
	Image     string     `gorm:"size:500" json:"image"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy int64      `gorm:"not null" json:"created_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *int64     `json:"deleted_by,omitempty"`
	Vehicle   *Vehicle   `gorm:"foreignKey:VehicleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"vehicle,omitempty"`
}
type ImageCars []ImageCar

func (ImageCar) TableName() string {
	return "image_cars"
}

type VehicleList struct {
	ID              int64      `gorm:"primaryKey" json:"id"`
	UUID            string     `gorm:"size:36;uniqueIndex:idx_vehicles_uuid;not null" json:"uuid"`
	Owner           string     `gorm:"not null;foreignKey:ID" json:"owner"` // user_id renamed to owner
	ServiceType     string     `gorm:"not null;foreignKey:ID" json:"service_type"`
	Brand           string     `gorm:"size:100;not null" json:"brand"`
	Year            int64      `gorm:"not null" json:"year"`
	Model           string     `gorm:"size:100;not null" json:"model"`
	LicensePlate    string     `gorm:"size:20;uniqueIndex:idx_vehicles_license_plate;not null" json:"license_plate"`
	Color           string     `gorm:"size:50;not null" json:"color"`
	TittleDeed      string     `gorm:"size:500;not null" json:"tittle_deed"`
	DriverName      string     `gorm:"size:150;not null" json:"driver_name"`
	UsedBy          *int64     `json:"used_by" example:"1" description:"ID del usuario que está usando el vehículo actualmente"`
	VIN             string     `gorm:"size:50;not null" json:"vin"`
	InsurancePolicy string     `gorm:"size:100;not null" json:"insurance_policy"`
	ImageCars       ImageCars  `gorm:"foreignKey:VehicleID;references:ID" json:"image_cars,omitempty"`
	CreatedAt       CustomTime `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy       int64      `gorm:"not null" json:"created_by"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	DeletedBy       *int64     `json:"deleted_by,omitempty"`
}

type VehiclesList []VehicleList

func (VehicleList) TableName() string {
	return "vehicles"
}
