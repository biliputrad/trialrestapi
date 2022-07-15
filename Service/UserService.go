package Service

import (
	"RestAPI-GETNPOST/Entity"
	"RestAPI-GETNPOST/Helper"
	"RestAPI-GETNPOST/Repository"
	"RestAPI-GETNPOST/dtos"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserService interface {
	CreateUser(request Entity.UserRequest) (dtos.Response, error)
	FindAllUser() (dtos.Response, error, Entity.Users)
	FindUserByID(id int) (dtos.Response, error, Entity.User)
	FindUsernameandPassword(username string, password string) (dtos.Response, error, Entity.User)
	UpdateUser(id int, userRequest Entity.UserRequest) (dtos.Response, error)
	DeleteUser(id int) (dtos.Response, error)
	ForgetPassword(id int, forgetPass Entity.ForgetPass) (dtos.Response, error)
	FindUsernameandEmail(username string, email string) (dtos.Response, error)
	FindUsernameandPhone(username string, phone uint64) (dtos.Response, error)
	PaginationUser(context *gin.Context, pagination *dtos.Pagination) (dtos.Response, error)
	GenerateUser(totaluser int) (dtos.Response, [][]string)
}
type serviceuser struct {
	userrepository Repository.UserRepository
}

func (s *serviceuser) FindUsernameandPassword(username string, password string) (dtos.Response, error, Entity.User) {
	user, err := s.userrepository.FindUsername(username)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err, user
	}
	err = Helper.CheckPassword(password, user.Password)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err, user
	}
	return dtos.Response{Success: true, Message: "OK"}, err, user

}

func (s *serviceuser) FindUserByID(id int) (dtos.Response, error, Entity.User) {
	user, err := s.userrepository.FindUserByID(id)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request", Data: user}, err, user
	}
	return dtos.Response{Success: true, Message: "OK", Data: user}, err, user
}

func NewServiceUser(userrepository Repository.UserRepository) *serviceuser {
	return &serviceuser{userrepository}
}
func (s *serviceuser) CreateUser(userRequest Entity.UserRequest) (dtos.Response, error) {
	err, hashedpass := Helper.HashPassword(userRequest.Password)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
	user := Entity.User{
		UserName:  userRequest.UserName,
		Email:     userRequest.Email,
		Password:  hashedpass,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Phone:     userRequest.Phone,
	}
	newUser, err := s.userrepository.CreateUser(user)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
	return dtos.Response{Success: true, Message: "OK", Data: newUser}, err
}

func (s *serviceuser) FindAllUser() (dtos.Response, error, Entity.Users) {
	users, err := s.userrepository.FindAllUser()
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request", Data: users}, err, users
	}
	return dtos.Response{Success: true, Message: "OK", Data: users}, err, users
}

func (s *serviceuser) UpdateUser(id int, userRequest Entity.UserRequest) (dtos.Response, error) {
	phone := userRequest.Phone
	user, err := s.userrepository.FindUserByID(id)
	err, hashedpass := Helper.HashPassword(userRequest.Password)
	user.UserName = userRequest.UserName
	user.Email = userRequest.Email
	user.Password = hashedpass
	user.Email = userRequest.Email
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Phone = uint64(phone)
	newUser, err := s.userrepository.UpdateUser(user)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request", Data: newUser}, err
	}
	return dtos.Response{Success: true, Message: "OK", Data: newUser}, err
}

func (s *serviceuser) DeleteUser(id int) (dtos.Response, error) {
	user, err := s.userrepository.FindUserByID(id)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
	deletedUser, err := s.userrepository.DeleteUser(user)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request", Data: deletedUser}, err
	}
	return dtos.Response{Success: true, Message: "OK", Data: deletedUser}, err
}

func (s *serviceuser) ForgetPassword(id int, forgetPass Entity.ForgetPass) (dtos.Response, error) {
	user, err := s.userrepository.FindUserByID(id)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request", Data: user}, err
	}
	err, hashpass := Helper.HashPassword(forgetPass.Password)
	user.Password = hashpass
	_, err2 := s.userrepository.UpdateUser(user)
	if err2 != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err2
	} else {
		return dtos.Response{Success: true, Message: "OK"}, err2
	}
}
func (s *serviceuser) FindUsernameandEmail(username string, email string) (dtos.Response, error) {
	user, err := s.userrepository.FindUsername(username)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
	if email == user.Email {
		return dtos.Response{Success: true, Message: "OK"}, err
	} else {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
}
func (s *serviceuser) FindUsernameandPhone(username string, phone uint64) (dtos.Response, error) {
	user, err := s.userrepository.FindUsername(username)
	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
	if phone == user.Phone {
		return dtos.Response{Success: true, Message: "OK"}, err
	} else {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}
}
func (s *serviceuser) PaginationUser(c *gin.Context, pagination *dtos.Pagination) (dtos.Response, error) {

	err, datapaginations, totalPages := s.userrepository.PaginationUser(pagination)

	if err != nil {
		return dtos.Response{Success: false, Message: "Bad Request"}, err
	}

	urlPath := c.Request.URL.Path

	datapaginations.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort)
	datapaginations.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort)

	if datapaginations.Page > 0 {

		datapaginations.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, datapaginations.Page-1, pagination.Sort)

	}
	if datapaginations.Page < totalPages {
		datapaginations.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, datapaginations.Page+1, pagination.Sort)
	}
	if datapaginations.Page > totalPages {
		datapaginations.PreviousPage = " "
	}
	return dtos.Response{Success: true, Message: "OK", Data: datapaginations}, err
}
func (s *serviceuser) GenerateUser(totaluser int) (dtos.Response, [][]string) {
	rows := [][]string{}
	for i := 1; i <= totaluser; i++ {
		response, err, user := s.FindUserByID(i)
		if err != nil {
			return response, rows
		}
		if user.ID != 0 {
			id := strconv.Itoa(user.ID)
			username := user.UserName
			firstname := user.FirstName
			lastname := user.LastName
			rows = append(rows, []string{id, username, firstname, lastname})
		}
	}
	return dtos.Response{Success: true, Message: "OK"}, rows
}
