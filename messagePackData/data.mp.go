package messagePackData

type ResponseTop struct {
	_msgpack struct{} `msgpack:",asArray"`
	Ts       int
	Pid      int
	Rev      int
	Login    Login
}

type Login struct {
	_msgpack     struct{} `msgpack:",asArray"`
	UserStatus   UserStatus
	UserCardList []UserCard
}

type UserStatus struct {
	_msgpack struct{} `msgpack:",asArray"`
	UserId   int
	UserName string
	Exp      int
}

type UserCard struct {
	_msgpack struct{} `msgpack:",asArray"`
	CardId   int
	Level    int
}
