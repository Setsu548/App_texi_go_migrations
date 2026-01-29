package models

import "time"

// Role representa la tabla roles para control de acceso
type Role struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:500;uniqueIndex:idx_roles_name;not null" json:"name"`
	Description string     `gorm:"size:500;not null" json:"description"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy   int64      `gorm:"not null" json:"created_by"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	DeletedBy   *int64     `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Role) TableName() string {
	return "roles"
}

// Permission representa la tabla permissions
type Permission struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:500;uniqueIndex:idx_permissions_name;not null" json:"name"`
	Description string     `gorm:"size:500;not null" json:"description"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy   int64      `gorm:"not null" json:"created_by"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	DeletedBy   *int64     `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Permission) TableName() string {
	return "permissions"
}

// UserRole representa la tabla user_roles
type UserRole struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	UserID     int64     `gorm:"not null;uniqueIndex:uk_user_role" json:"user_id"`
	User       *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	RoleID     int64     `gorm:"not null;uniqueIndex:uk_user_role" json:"role_id"`
	Role       *Role     `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role,omitempty"`
	AssignedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"assigned_at"`
	AssignedBy int64     `gorm:"not null" json:"assigned_by"`
}

// TableName especifica el nombre de la tabla
func (UserRole) TableName() string {
	return "user_roles"
}

// RolePermission representa la tabla role_permissions
type RolePermission struct {
	ID           int64       `gorm:"primaryKey" json:"id"`
	RoleID       int64       `gorm:"not null;uniqueIndex:uk_role_permission" json:"role_id"`
	Role         *Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role,omitempty"`
	PermissionID int64       `gorm:"not null;uniqueIndex:uk_role_permission" json:"permission_id"`
	Permission   *Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"permission,omitempty"`
	AssignedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP;not null" json:"assigned_at"`
	AssignedBy   int64       `gorm:"not null" json:"assigned_by"`
}

// TableName especifica el nombre de la tabla
func (RolePermission) TableName() string {
	return "role_permissions"
}
