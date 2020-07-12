package wxxx

type UserDataStorage interface {
	GetUserData() interface{}
	SetUserData(interface{})
}

type BaseUserData struct {
	userData interface{}
}

func (this *BaseUserData) GetUserData() interface{} {
	return this.userData
}

func (this *BaseUserData) SetUserData(data interface{}) {
	this.userData = data
}
