package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime es un tipo personalizado para time.Time que se serializa sin microsegundos
type CustomTime struct {
	time.Time
}

// MarshalJSON implementa la interfaz json.Marshaler para serializar sin microsegundos
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, ct.Time.Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON implementa la interfaz json.Unmarshaler
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}
	// Remover comillas
	str = str[1 : len(str)-1]
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// Value implementa la interfaz driver.Valuer para GORM
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	return ct.Time.Format("2006-01-02 15:04:05"), nil
}

// Scan implementa la interfaz sql.Scanner para GORM
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	// Manejar diferentes tipos de valores que pueden venir de la BD
	switch v := value.(type) {
	case time.Time:
		ct.Time = v
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			// Intentar con el formato de MySQL que incluye microsegundos
			t, err = time.Parse("2006-01-02 15:04:05.000000", v)
			if err != nil {
				return err
			}
		}
		ct.Time = t
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			// Intentar con el formato de MySQL que incluye microsegundos
			t, err = time.Parse("2006-01-02 15:04:05.000000", string(v))
			if err != nil {
				return err
			}
		}
		ct.Time = t
	default:
		return fmt.Errorf("cannot scan %T into CustomTime", value)
	}
	return nil
}

// User representa la tabla usuarios del sistema
type User struct {
	ID           int64      `gorm:"primaryKey" json:"id"`
	UUID         string     `gorm:"size:500;uniqueIndex;not null" json:"uuid"`
	TypeUserID   int64      `gorm:"not null" json:"type_user_id"`
	RoleID       *int64     `json:"role_id,omitempty"`
	UserName     string     `gorm:"size:500;uniqueIndex;not null" json:"user_name"` // Email del usuario
	PasswordHash *string    `gorm:"size:500" json:"-"`
	RegisterStep int        `gorm:"default:0;not null" json:"register_step"`
	IsVerified   bool       `gorm:"default:false;not null" json:"is_verified"`
	Status       string     `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy    *int64     `json:"created_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UpdatedBy    *int64     `json:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	DeletedBy    *int64     `json:"deleted_by,omitempty"`

	TypeUser      *TypeUser      `gorm:"foreignKey:TypeUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"type_user,omitempty"`
	Role          *Role          `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role,omitempty"`
	Staff         *Staff         `gorm:"foreignKey:ID;references:ID"  json:"staff,omitempty"`
	Employee      *Employee      `gorm:"foreignKey:ID;references:ID" json:"employee,omitempty"`
	DocumentsInfo []DocumentInfo `gorm:"foreignKey:UserID;references:ID" json:"documents_info,omitempty"`
}

// `gorm:"foreignKey:ID;references:UserID" json:"user,omitempty"`

// TableName especifica el nombre de la tabla
func (User) TableName() string {
	return "users"
}

// TypeUser representa la tabla type_user
type TypeUser struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	UUID      string     `gorm:"size:500;uniqueIndex;not null" json:"uuid"`
	TypeName  string     `gorm:"size:500;not null" json:"type_name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy int64      `gorm:"not null" json:"created_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy *int64     `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (TypeUser) TableName() string {
	return "type_user"
}

// Constants para tipos de usuario
const (
	UserTypeAdmin     = "admin"
	UserTypePassenger = "passanger"
	UserTypeDriver    = "driver"
	UserTypeSupport   = "support"
	UserTypeManager   = "manager"
	UserTypeOperator  = "operator"
)

// Status constants
const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusPending  = "pending"
	UserStatusBlocked  = "blocked"
)
