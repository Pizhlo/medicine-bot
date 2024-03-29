package user

type UserSrv struct {
	userEditor userEditor
}

type userEditor interface{}

func New(userEditor userEditor) *UserSrv {
	return &UserSrv{userEditor: userEditor}
}
