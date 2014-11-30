package goty

type Mode struct {
}

func (m *Mode) CanRemove() (bool) {
	switch mType := m.(type) {
	case User_Away:
		return true
	case User_Invisible:
		return false
	case User_Wallops:
		return false
	case User_Restricted_Connection:
		return add
	case User_Operator:
		return true
	case User_Local_Operator:
		return true
	case User_Server_Notices:
		return true
	}
}

func (m *Mode) CanAdd() (bool) {
	switch mType := m.(type) {
	case User_Away:
		return true
	case User_Invisible:
		return false
	case User_Wallops:
		return false
	case User_Restricted_Connection:
		return add
	case User_Operator:
		return true
	case User_Local_Operator:
		return true
	case User_Server_Notices:
		return true
	}
}

type User_Away struct {
	Mode
}

type User_Invisible struct {
	Mode
}

type User_Wallops struct {
	Mode
}

type User_Restricted_Connection struct {
	Mode
}

type User_Operator struct {
	Mode
}

type User_Local_Operator struct {
	Mode
}

type User_Server_Notices struct {
	Mode
}

type Modes struct {
	Modes []Mode
}

type UserModes struct {
	Modes
	validModes []Mode
}

type FreenodeModes struct {
	UserModes
}

