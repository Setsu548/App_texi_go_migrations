package models

import "time"

// DocumentType representa la tabla document_type
type DocumentType struct {
	ID               int64      `gorm:"primaryKey" json:"id"`
	DocumentTypeName string     `gorm:"size:100;not null" json:"document_type_name"`
	Description      string     `gorm:"size:100;not null" json:"description"`
	CreatedAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy        int64      `gorm:"not null" json:"created_by"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
	DeletedBy        *int64     `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (DocumentType) TableName() string {
	return "document_type"
}

// DocumentInfo representa la tabla document_info para almacenar documentos de usuarios
type DocumentInfo struct {
	ID                 int64         `gorm:"primaryKey" json:"id"`
	UserID             int64         `gorm:"not null" json:"user_id"`
	User               *User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	DocumentTypeID     int64         `gorm:"not null" json:"document_type_id"`
	DocumentType       *DocumentType `gorm:"foreignKey:DocumentTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"document_type,omitempty"`
	DocumentNumber     string        `gorm:"size:255;not null" json:"document_number"`
	FrontDocumentImage string        `gorm:"type:text;not null" json:"front_document_image"`
	BackDocumentImage  string        `gorm:"type:text;not null" json:"back_document_image"`
	FaceImage          *string       `gorm:"type:text" json:"face_image"`
	ExpireDate         time.Time     `gorm:"not null" json:"expire_date"`
	IsActive           bool          `gorm:"not null" json:"is_active"`
	CreatedAt          time.Time     `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	CreatedBy          int64         `gorm:"not null" json:"created_by"`
	DeletedAt          *time.Time    `json:"deleted_at,omitempty"`
	DeletedBy          *int64        `json:"deleted_by,omitempty"`
}

// TableName especifica el nombre de la tabla
func (DocumentInfo) TableName() string {
	return "document_info"
}

// DocumentType IDs constants
const (
	DocumentTypeDNI           = 1
	DocumentTypeDriverLicense = 6
	DocumentTypeMotoLicense   = 7
	DocumentTypeTruckLicense  = 8
)
