package Helper

import (
	"RestAPI-GETNPOST/Entity"
)

func ConvertData(user Entity.User) Entity.DisplayUser {
	display := Entity.DisplayUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
	}
	return display

}
