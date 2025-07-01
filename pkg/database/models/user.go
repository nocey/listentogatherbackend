package models

type User struct {
	Model
	Name        string        `json:"name"`
	Permissions []Permissions `gorm:"many2many:user_permissions;foreignKey:ID;joinForeignKey:PermissionID;References:ID;joinReferences:UserID"`
}

func (user *User) HasPermission(perm *Permissions) bool {
	userHasPerm := false
	if user != nil {
		for _, userPerm := range user.Permissions {
			if userPerm == *perm {
				userHasPerm = true
				break
			}
		}
	}
	return userHasPerm
}
