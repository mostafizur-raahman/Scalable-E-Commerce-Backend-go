package domain

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Reader interface {
	GetAll() ([]User, error)
}

type Getter interface {
	GetByID(id int64) (*User, error)
}

type Creator interface {
	Create(User *User) error
}

type Updater interface {
	Update(id int64) error
}

type Deleter interface {
	Delete(id int64) error
}

type Repository interface {
	Reader
	Getter
	Creator
	Updater
	Deleter
}
