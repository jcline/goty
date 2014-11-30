package goty

type Mode struct {
    repr string
}

func (m *Mode) Repr() string {
    return m.repr
}

type User_Away struct {
	Mode
}

func NewUser_Away() *User_Away { return &User_Away{ Mode{repr: "a" }} }
func (m *User_Away) CanAdd() bool { return false }
func (m *User_Away) CanRemove() bool { return false }

type User_Invisible struct {
	Mode
}

func NewUser_Invisible() *User_Invisible { return &User_Invisible{ Mode{repr: "i"} } }
func (m *User_Invisible) CanAdd() bool { return true }
func (m *User_Invisible) CanRemove() bool { return true }

type User_Wallops struct {
	Mode
}

func NewUser_Wallops() *User_Wallops { return &User_Wallops{ Mode{repr: "w"} } }
func (m *User_Wallops) CanAdd() bool { return true }
func (m *User_Wallops) CanRemove() bool { return true }

type User_Restricted_Connection struct {
	Mode
}

func NewUser_Restricted_Connection() *User_Restricted_Connection { return &User_Restricted_Connection{ Mode{repr: "r"} } }
func (m *User_Restricted_Connection) CanAdd() bool { return true }
func (m *User_Restricted_Connection) CanRemove() bool { return false }

type User_Operator struct {
	Mode
}

func NewUser_Operator() *User_Operator { return &User_Operator{ Mode{repr: "o"} } }
func (m *User_Operator) CanAdd() bool { return false }
func (m *User_Operator) CanRemove() bool { return true }

type User_Local_Operator struct {
	Mode
}

func NewUser_Local_Operator() *User_Local_Operator { return &User_Local_Operator{ Mode{repr: "O"} } }
func (m *User_Local_Operator) CanAdd() bool { return false }
func (m *User_Local_Operator) CanRemove() bool { return true }

type User_Server_Notices struct {
	Mode
}

func NewUser_Server_Notices() *User_Server_Notices { return &User_Server_Notices{ Mode{repr: "s"} } }
func (m *User_Server_Notices) CanAdd() bool { return true }
func (m *User_Server_Notices) CanRemove() bool { return true }

type Modes struct {
	Modes []Mode
}

type UserModes struct {
	Modes
	validModes []Mode
}

// -------------------------
type Channel_Creator struct {
	Mode
}

func NewChannel_Creator() *Channel_Creator { return &Channel_Creator{ Mode{repr: "O" }} }
func (m *Channel_Creator) CanAdd() bool { return false }
func (m *Channel_Creator) CanRemove() bool { return false }

type Channel_Operator struct {
	Mode
}

func NewChannel_Operator() *Channel_Operator { return &Channel_Operator{ Mode{repr: "o" }} }
func (m *Channel_Operator) CanAdd() bool { return false }
func (m *Channel_Operator) CanRemove() bool { return false }

type Channel_Voice struct {
	Mode
}

func NewChannel_Voice() *Channel_Voice { return &Channel_Voice{ Mode{repr: "v" }} }
func (m *Channel_Voice) CanAdd() bool { return false }
func (m *Channel_Voice) CanRemove() bool { return false }

type Channel_Anonymous struct {
	Mode
}

func NewChannel_Anonymous() *Channel_Anonymous { return &Channel_Anonymous{ Mode{repr: "a" }} }
func (m *Channel_Anonymous) CanAdd() bool { return false }
func (m *Channel_Anonymous) CanRemove() bool { return false }

type Channel_Invite_Only struct {
	Mode
}

func NewChannel_Invite_Only() *Channel_Invite_Only { return &Channel_Invite_Only{ Mode{repr: "i" }} }
func (m *Channel_Invite_Only) CanAdd() bool { return false }
func (m *Channel_Invite_Only) CanRemove() bool { return false }

type Channel_Moderated struct {
	Mode
}

func NewChannel_Moderated() *Channel_Moderated { return &Channel_Moderated{ Mode{repr: "m" }} }
func (m *Channel_Moderated) CanAdd() bool { return false }
func (m *Channel_Moderated) CanRemove() bool { return false }

type Channel_No_External_Messages struct {
	Mode
}

func NewChannel_No_External_Messages() *Channel_No_External_Messages { return &Channel_No_External_Messages{ Mode{repr: "n" }} }
func (m *Channel_No_External_Messages) CanAdd() bool { return false }
func (m *Channel_No_External_Messages) CanRemove() bool { return false }

type Channel_Quiet struct {
	Mode
}

func NewChannel_Quiet() *Channel_Quiet { return &Channel_Quiet{ Mode{repr: "q" }} }
func (m *Channel_Quiet) CanAdd() bool { return false }
func (m *Channel_Quiet) CanRemove() bool { return false }

type Channel_Private struct {
	Mode
}

func NewChannel_Private() *Channel_Private { return &Channel_Private{ Mode{repr: "p" }} }
func (m *Channel_Private) CanAdd() bool { return false }
func (m *Channel_Private) CanRemove() bool { return false }

type Channel_Secret struct {
	Mode
}

func NewChannel_Secret() *Channel_Secret { return &Channel_Secret{ Mode{repr: "s" }} }
func (m *Channel_Secret) CanAdd() bool { return false }
func (m *Channel_Secret) CanRemove() bool { return false }

type Channel_Server_Reop struct {
	Mode
}

func NewChannel_Server_Reop() *Channel_Server_Reop { return &Channel_Server_Reop{ Mode{repr: "r" }} }
func (m *Channel_Server_Reop) CanAdd() bool { return false }
func (m *Channel_Server_Reop) CanRemove() bool { return false }

type Channel_Topic_Set_Operator_Only struct {
	Mode
}

func NewChannel_Topic_Set_Operator_Only() *Channel_Topic_Set_Operator_Only { return &Channel_Topic_Set_Operator_Only{ Mode{repr: "t" }} }
func (m *Channel_Topic_Set_Operator_Only) CanAdd() bool { return false }
func (m *Channel_Topic_Set_Operator_Only) CanRemove() bool { return false }

type Channel_Key_Set struct {
	Mode
}

func NewChannel_Key_Set() *Channel_Key_Set { return &Channel_Key_Set{ Mode{repr: "k" }} }
func (m *Channel_Key_Set) CanAdd() bool { return false }
func (m *Channel_Key_Set) CanRemove() bool { return false }

type Channel_User_Limit struct {
	Mode
}

func NewChannel_User_Limit() *Channel_User_Limit { return &Channel_User_Limit{ Mode{repr: "l" }} }
func (m *Channel_User_Limit) CanAdd() bool { return false }
func (m *Channel_User_Limit) CanRemove() bool { return false }

type Channel_Ban_Mask struct {
	Mode
}

func NewChannel_Ban_Mask() *Channel_Ban_Mask { return &Channel_Ban_Mask{ Mode{repr: "b" }} }
func (m *Channel_Ban_Mask) CanAdd() bool { return false }
func (m *Channel_Ban_Mask) CanRemove() bool { return false }

type Channel_Exception_Mask struct {
	Mode
}

func NewChannel_Exception_Mask() *Channel_Exception_Mask { return &Channel_Exception_Mask{ Mode{repr: "e" }} }
func (m *Channel_Exception_Mask) CanAdd() bool { return false }
func (m *Channel_Exception_Mask) CanRemove() bool { return false }

type Channel_Invitation_Mask struct {
	Mode
}

func NewChannel_Invitation_Mask() *Channel_Invitation_Mask { return &Channel_Invitation_Mask{ Mode{repr: "I" }} }
func (m *Channel_Invitation_Mask) CanAdd() bool { return false }
func (m *Channel_Invitation_Mask) CanRemove() bool { return false }
