package go_learning_common

type SysUser struct {
	UserId  		int64		`json:"user_id" bson:"userId"`
	Username		string		`json:"username" bson:"username"`
	Password		string		`json:"password" bson:"password"`
	Salt			string		`json:"salt" bson:"salt"`
	Email			string		`json:"email" bson:"email"`
	Mobile			string		`json:"mobile" bson:"mobile"`
	Status			int			`json:"status" bson:"status"`
	DeptId			int64		`json:"dept_id" bson:"deptId"`
	CreateTime		string		`json:"create_time" bson:"createTime"`
	Avatar			string		`json:"avatar"  bson:"avatar"`
}
