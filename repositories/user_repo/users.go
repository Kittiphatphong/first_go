package user_repo

import (
	"clickcash_backend/logs"
	"gorm.io/gorm"
)


type UserRepo interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(ID uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type userRepo struct {
	db *gorm.DB
}

func (u userRepo) GetUserByEmail(email string) (*User, error) {
	user := User{}
	u.db.Where("email",email).Find(&user)
	return &user,nil
}

func (u userRepo) GetUserByID(ID uint) (*User, error) {
	user := User{}
	u.db.Where("id",ID).Find(&user)
	return &user,nil
}

func (u userRepo) CreateUser(user *User) (*User, error) {
	//var users User
	//user.ID= uuid.New()
	err := u.db.Create(&user).Error
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return user, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	//db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	return &userRepo{db: db}
}
