package models

type Users struct {
	Model
	Name        string        `json:"name"`
	Password    string        `json:"password"`
	Permissions []Permissions `gorm:"many2many:user_permissions;foreignKey:ID;joinForeignKey:PermissionID;References:ID;joinReferences:UserID" json:"permissions"`
}

func (user *Users) HasPermission(perm *Permissions) bool {
	userHasPerm := false
	if user != nil {
		for _, userPerm := range user.Permissions {
			if userPerm.Name == perm.Name {
				userHasPerm = true
				break
			}
		}
	}
	return userHasPerm
}
