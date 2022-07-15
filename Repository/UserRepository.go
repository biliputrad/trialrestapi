package Repository

import (
	"RestAPI-GETNPOST/Entity"
	"RestAPI-GETNPOST/dtos"
	"gorm.io/gorm"
	"math"
)

type UserRepository interface {
	CreateUser(user Entity.User) (Entity.User, error)
	FindAllUser() ([]Entity.User, error)
	FindUserByID(id int) (Entity.User, error)
	FindUsername(username string) (Entity.User, error)
	UpdateUser(user Entity.User) (Entity.User, error)
	DeleteUser(user Entity.User) (Entity.User, error)
	PaginationUser(pagination *dtos.Pagination) (error, *dtos.Pagination, int)
}
type userrepository struct {
	db *gorm.DB
}

func (r *userrepository) FindUserByID(id int) (Entity.User, error) {
	var user Entity.User

	err := r.db.Find(&user, id).Error

	return user, err
}

func NewRepositoryUser(db *gorm.DB) *userrepository {
	return &userrepository{db}
}
func (r *userrepository) CreateUser(user Entity.User) (Entity.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
func (r *userrepository) FindAllUser() ([]Entity.User, error) {
	var users []Entity.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userrepository) FindUsername(username string) (Entity.User, error) {
	var user Entity.User

	err := r.db.Find(&user, "user_name = ? ", username).Error
	if err != nil {
		panic("Error")
	}
	return user, err
}

func (r *userrepository) UpdateUser(user Entity.User) (Entity.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userrepository) DeleteUser(user Entity.User) (Entity.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
func (r *userrepository) PaginationUser(pagination *dtos.Pagination) (error, *dtos.Pagination, int) {
	var users Entity.Users

	var totalRows int64

	totalPages, fromRow, toRow := 0, 0, 0

	offset := pagination.Page * pagination.Limit

	errFind := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&users).Error
	var usersResponse []Entity.DisplayUser

	for _, u := range users {
		userRespond := Entity.DisplayUser{
			ID:        u.ID,
			UserName:  u.UserName,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Phone:     u.Phone,
		}
		usersResponse = append(usersResponse, userRespond)
	}
	if errFind != nil {
		return errFind, pagination, totalPages
	}
	pagination.Rows = usersResponse

	errCount := r.db.Model(&Entity.User{}).Count(&totalRows).Error

	if errCount != nil {
		return errCount, pagination, totalPages
	}

	pagination.TotalRows = int(totalRows)

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return errCount, pagination, totalPages
}
