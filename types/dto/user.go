package dto

type User struct {
	Username 	string			`json:"username" bson:"username"`
	FullName 	string			`json:"full_name" bson:"full_name"`
	Sex 		string			`json:"sex" bson:"sex"`
	Birthday	int64			`json:"birthday" bson:"birthday"`
	Phone		string			`json:"phone" bson:"phone"`
	Avatar		string			`json:"avatar" bson:"avatar"`
	OS 			string			`json:"os" bson:"os"`
	Email 		string			`json:"email" bson:"email"`
}
