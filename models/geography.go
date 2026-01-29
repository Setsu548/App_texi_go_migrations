package models

import "time"

// Country representa la tabla countries
type Country struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	UUID      string     `gorm:"size:500;uniqueIndex:idx_countries_uuid;not null" json:"uuid"`
	Name      string     `gorm:"size:100;uniqueIndex:idx_countries_name;not null" json:"name"`
	IsoCode   *string    `gorm:"size:5" json:"iso_code"`
	PhoneCode *string    `gorm:"size:10" json:"phone_code"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy *int64     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName especifica el nombre de la tabla
func (Country) TableName() string {
	return "countries"
}

// Department representa la tabla departments
type Department struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	UUID      string     `gorm:"size:500;uniqueIndex:idx_departments_uuid;not null" json:"uuid"`
	CountryID int64      `gorm:"not null" json:"country_id"`
	Country   *Country   `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"country,omitempty"`
	Name      string     `gorm:"size:100;not null" json:"name"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy *int64     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName especifica el nombre de la tabla
func (Department) TableName() string {
	return "departments"
}

// Locality representa la tabla localities
type Locality struct {
	ID           int64       `gorm:"primaryKey" json:"id"`
	UUID         string      `gorm:"size:500;uniqueIndex:idx_localities_uuid;not null" json:"uuid"`
	DepartmentID int64       `gorm:"not null" json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"department,omitempty"`
	Name         string      `gorm:"size:100;not null" json:"name"`
	ZipCode      *string     `gorm:"size:20" json:"zip_code"`
	CreatedAt    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy    *int64      `json:"created_by"`
	UpdatedAt    *time.Time  `json:"updated_at"`
	DeletedAt    *time.Time  `json:"deleted_at"`
}

// TableName especifica el nombre de la tabla
func (Locality) TableName() string {
	return "localities"
}
