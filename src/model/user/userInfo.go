package user


//账号信息
type AccountInfo struct {
	Id int64 `json:"id" db:"id"`
	UserName string `json:"userName" db:"username"`
	PassWord string `json:"password" db:"password"`
}
//登录或注册
type LoginOrRegister struct {
	 AccountNumber string `form:"accountNumber" binding:"required"`
	 PassWord string `form:"password" binding:"required"`
}
//用户信息
type UpdateInfo struct {
	Id int64 `db:"id" form:"id" binding:"required"`
	Sex int8 `db:"sex" form:"sex" binding:"min=0,max=3"`
	UserName string `db:"username" form:"userName"`
	Introduce string `db:"introduce" form:"introduce"`
}
//返回用户信息
type UserInfoResult struct {
	UserId int64 `json:"userId" db:"id"`
	Sex int8 `json:"sex"`
	UserName string `json:"userName"`
	Introduce string `json:"introduce"`
	Token string `json:"token"`
}
//登录验证
type LoginResult struct {
	 Id int64 `db:"id"`
	 PassWord string `db:"password"`
}
