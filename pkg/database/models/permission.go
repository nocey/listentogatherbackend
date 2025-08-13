package models

type Permissions struct {
	Model
	Name  string  `json:"name"`
	Users []Users `gorm:"many2many:user_permissions;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:PermissionID"`
}
