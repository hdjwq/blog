package modal

type Account struct {
	AccountNumber string `binding:"required" form:"accountNumber" db:"account_number"`
	Password string `binding:"required" form:"password" db:"password"`
}
type AccountResult struct {
	Id int64 `db:"id"`
	Account
}
type UserInfo struct {
	UpdateUserInfo
	Token string `json:"token"`
}
type UpdateUserInfo struct {
	UserId int64 `db:"id" binding:"required" form:"userId" json:"userId"`
	Username string `db:"username" binding:"required" form:"userName" json:"userName"`
	Sex string `db:"sex" binding:"required,number" form:"sex" json:"sex"`
	Introduce string `db:"introduce" form:"introduce" json:"introduce"`
}