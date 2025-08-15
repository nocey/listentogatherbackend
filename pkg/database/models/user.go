package models

import logger "github.com/listentogether/log"

type Users struct {
	Model
	Name        string        `json:"name"`
	Password    string        `json:"password"`
	Permissions []Permissions `gorm:"many2many:user_permissions;foreignKey:ID;joinForeignKey:user_id;References:ID;joinReferences:permission_id" json:"permissions"`
	UsersPosts  []Posts       `gorm:"foreignKey:UserID" json:"posts"`
}

func (user *Users) HasPermission(perm *Permissions) bool {
	userHasPerm := false
	if user != nil {
		logger.Debug("Checking permissions for user:", user.Permissions)
		for _, userPerm := range user.Permissions {
			if userPerm.Name == perm.Name {
				userHasPerm = true
				break
			}
		}
	}
	return userHasPerm
}
