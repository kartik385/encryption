package ports


type PasswordService interface{
	GetPassword(key string)(string,error)
	SetPassword(key string,password string)(error)
	DeletePassword(key string)(error)
}