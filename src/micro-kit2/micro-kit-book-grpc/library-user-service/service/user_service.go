package service

import (
	"context"
	"errors"
	"log"
	"micro/kit/library-user-service/dao"
	"micro/kit/library-user-service/dto"
	"micro/kit/library-user-service/models"

	pbbook "micro/kit/protos/book"

	"github.com/jinzhu/gorm"
)

var (
	ErrNotFound    = errors.New(" User is not exist! ")
	ErrUserExisted = errors.New(" User is existed! ")
	ErrRegistering = errors.New(" User can't regist! ")
)

type UserService interface {
	Register(ctx context.Context, vo *dto.RegisterUser) (*dto.UserInfo, error)
	FindByID(ctx context.Context, id uint64) (*dto.UserInfo, error)
	FindByEmail(ctx context.Context, email string) (*dto.UserInfo, error)
	FindBooksByUserID(ctx context.Context, id uint64) (interface{}, error)
}

type UserServiceImpl struct {
	userDao    dao.UserDao
	bookClient pbbook.BookClient
}

func NewUserServiceImpl(userDao dao.UserDao, bookClient pbbook.BookClient) UserService {
	return &UserServiceImpl{
		userDao:    userDao,
		bookClient: bookClient,
	}
}

func (u *UserServiceImpl) Register(ctx context.Context,
	vo *dto.RegisterUser) (*dto.UserInfo, error) {
	user, err := u.userDao.SelectByEmail(vo.Email)
	if user != nil {
		log.Println("User is already exist!")
		return &dto.UserInfo{}, ErrUserExisted
	}
	if err == gorm.ErrRecordNotFound || err == nil {
		newUser := &models.User{
			Username: vo.Username,
			Password: vo.Password,
			Email:    vo.Email,
		}
		err = u.userDao.Save(newUser)
		if err != nil {
			return nil, ErrRegistering
		}
		return &dto.UserInfo{
			ID:       newUser.ID,
			Username: newUser.Username,
			Email:    newUser.Email,
		}, nil
	}
	return nil, err
}

func (u *UserServiceImpl) FindByID(ctx context.Context, id uint64) (*dto.UserInfo, error) {
	user, err := u.userDao.SelectByID(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return &dto.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (u *UserServiceImpl) FindByEmail(ctx context.Context, email string) (*dto.UserInfo, error) {
	user, err := u.userDao.SelectByEmail(email)
	if err != nil {
		return nil, ErrNotFound
	}
	return &dto.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// FindBooksByUserID grpc的调用
func (u *UserServiceImpl) FindBooksByUserID(ctx context.Context, id uint64) (interface{}, error) {
	req := &pbbook.BooksByUserIDRequest{UserID: id}
	//调用library-book-grpc-service  FindBooksByUserID
	res, err := u.bookClient.FindBooksByUserID(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
