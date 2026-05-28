package domain

type Product struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Reader interface {
	GetAll() ([]Product, error)
}

type Getter interface {
	GetByID(id int64) (*Product, error)
}

type Creator interface {
	Create(product *Product) error
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
