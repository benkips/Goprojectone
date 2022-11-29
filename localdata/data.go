package localdata

type Data interface {
	GetData() []UserModel
}

type data struct {
	userList []UserModel
}

func (d data) GetData() []UserModel {
	return d.userList
}

func InitData() Data {
	return &data{
		userList: []UserModel{
			UserModel{
				ID:     1,
				Name:   "A",
				Age:    18,
				Gender: "Male",
			},
			UserModel{
				ID:     2,
				Name:   "Chi",
				Age:    17,
				Gender: "Female",
			},
			UserModel{
				ID:     3,
				Name:   "Shintaro",
				Age:    18,
				Gender: "Male",
			},
			UserModel{
				ID:     4,
				Name:   "Ayano",
				Age:    18,
				Gender: "Female",
			},
		},
	}
}
