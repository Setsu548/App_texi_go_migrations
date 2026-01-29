package models

import "time"

// Fare representa la tabla fares (tarifas)
type Fare struct {
	ID              int64        `gorm:"primaryKey" json:"id"`
	BaseFare        float64      `gorm:"type:decimal(10,2);not null" json:"base_fare"`
	ServiceTypeID   int64        `gorm:"not null" json:"service_type_id"`
	ServiceType     *ServiceType `gorm:"foreignKey:ServiceTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"service_type,omitempty"`
	City            string       `gorm:"size:100;not null" json:"city"`
	Country         string       `gorm:"size:100;not null" json:"country"`
	PerKilometer    float64      `gorm:"type:decimal(10,2);not null" json:"per_kilometer"`
	MinimumDistance *float64     `gorm:"type:decimal(10,2)" json:"minimum_distance"`
	AirportFare     *float64     `gorm:"type:decimal(10,2)" json:"airport_fare"`
	WaitingCost     *float64     `gorm:"type:decimal(10,2)" json:"waiting_cost"`
	Status          string       `gorm:"type:varchar(20);default:'active';not null" json:"status"`
	CreatedAt       time.Time    `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy       *int64       `json:"created_by"`
	UpdatedAt       *time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy       *int64       `json:"updated_by"`
	DeletedAt       *time.Time   `json:"deleted_at"`
	DeletedBy       *int64       `json:"deleted_by"`
}

// TableName especifica el nombre de la tabla
func (Fare) TableName() string {
	return "fares"
}

// Trip representa la tabla trips (viajes)
type Trip struct {
	ID                   int64      `gorm:"primaryKey" json:"id"`
	PassengerID          int64      `gorm:"not null" json:"passenger_id"`
	Passenger            *User      `gorm:"foreignKey:PassengerID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"passenger,omitempty"`
	DriverID             *int64     `json:"driver_id"`
	Driver               *User      `gorm:"foreignKey:DriverID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"driver,omitempty"`
	VehicleID            *int64     `json:"vehicle_id"`
	Vehicle              *Vehicle   `gorm:"foreignKey:VehicleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"vehicle,omitempty"`
	OriginAddress        string     `gorm:"size:255;not null" json:"origin_address"`
	OriginLatitude       float64    `gorm:"type:decimal(10,8);not null" json:"origin_latitude"`
	OriginLongitude      float64    `gorm:"type:decimal(11,8);not null" json:"origin_longitude"`
	DestinationAddress   string     `gorm:"size:255;not null" json:"destination_address"`
	DestinationLatitude  float64    `gorm:"type:decimal(10,8);not null" json:"destination_latitude"`
	DestinationLongitude float64    `gorm:"type:decimal(11,8);not null" json:"destination_longitude"`
	Status               string     `gorm:"type:varchar(20);default:'pending';not null" json:"status"`
	RequestedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"requested_at"`
	StartedAt            *time.Time `json:"started_at"`
	EndedAt              *time.Time `json:"ended_at"`
	UsedBy               *int64     `json:"used_by"`
	EstimatedFare        *float64   `gorm:"type:decimal(10,2)" json:"estimated_fare"`
	FinalFare            *float64   `gorm:"type:decimal(10,2)" json:"final_fare"`
}

// TableName especifica el nombre de la tabla
func (Trip) TableName() string {
	return "trips"
}

// Trip status constants
const (
	TripStatusPending    = "pending"
	TripStatusInProgress = "in_progress"
	TripStatusCanceled   = "canceled"
	TripStatusCompleted  = "completed"
)
