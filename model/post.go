package model

const (
	PostStatusPrivate = 0
	PostStatusPublic  = 1
)

type Post struct {
	Id        int    `json:"id"`
	Content   string `json:"content" gorm:"type:string"`
	Timestamp int64  `json:"timestamp" gorm:"type:int64"`
	Status    int    `json:"status" gorm:"type:int;default:1"`
}

func (p *Post) Insert() error {
	var err error
	err = DB.Create(p).Error
	return err
}

func (p *Post) Update() error {
	var err error
	err = DB.Model(p).Updates(p).Error
	return err
}

func (p *Post) Delete() error {
	var err error
	err = DB.Delete(p).Error
	return err
}
