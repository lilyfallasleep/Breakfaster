package model

// Employee is employee model type
type Employee struct {
	EmpID         string `gorm:"type:char(7);primaryKey"`
	LineUID       string `gorm:"type:char(33);uniqueIndex"`
	AccessCardNbr string `gorm:"type:char(10);unique"`
	Dept          string `gorm:"type:char(128)"`
	ResignAt      int64
	UpdatedAt     int64 `gorm:"autoUpdateTime"`
	CreatedAt     int64 `gorm:"autoCreateTime"`

	Foods []Food `gorm:"many2many:orders;"`
}
