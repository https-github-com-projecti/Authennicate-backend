package pojo

type CheckPass struct {
	ID       string `bson:"id,omitempty"`
	Password string `bson:"password,omitempty"`
}
