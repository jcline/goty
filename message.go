package goty

type IRCMessage struct {
}

type Pass struct {
	password	string
	IRCMessage
}

type Nick struct {
	nickname string
	IRCMessage
}

type User struct {
	user, mode, unused, realname string
	IRCMessage
}

type Oper struct {
	name, password string
	IRCMessage
}

type User_Mode struct {
	nickname string
	IRCMessage
	added, removed []UserModes
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}

type struct {
	string
	IRCMessage
}
