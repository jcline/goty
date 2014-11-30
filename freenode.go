package goty

type Freenode_User_Deaf struct {
	Mode
}

func NewFreenode_User_Deaf() *Freenode_User_Deaf { return &Freenode_User_Deaf{ Mode{repr: "D" }} }
func (m *Freenode_User_Deaf) CanAdd() bool { return false }
func (m *Freenode_User_Deaf) CanRemove() bool { return false }

type Freenode_User_Caller_ID struct {
	Mode
}

func NewFreenode_User_Caller_ID() *Freenode_User_Caller_ID { return &Freenode_User_Caller_ID{ Mode{repr: "g" }} }
func (m *Freenode_User_Caller_ID) CanAdd() bool { return false }
func (m *Freenode_User_Caller_ID) CanRemove() bool { return false }

type Freenode_User_No_Forwarding struct {
	Mode
}

func NewFreenode_User_No_Forwarding() *Freenode_User_No_Forwarding { return &Freenode_User_No_Forwarding{ Mode{repr: "Q" }} }
func (m *Freenode_User_No_Forwarding) CanAdd() bool { return false }
func (m *Freenode_User_No_Forwarding) CanRemove() bool { return false }

type Freenode_User_Block_Unidentified struct {
	Mode
}

func NewFreenode_User_Block_Unidentified() *Freenode_User_Block_Unidentified { return &Freenode_User_Block_Unidentified{ Mode{repr: "R" }} }
func (m *Freenode_User_Block_Unidentified) CanAdd() bool { return false }
func (m *Freenode_User_Block_Unidentified) CanRemove() bool { return false }

type Freenode_User_Connected_SSL struct {
	Mode
}

func NewFreenode_User_Connected_SSL() *Freenode_User_Connected_SSL { return &Freenode_User_Connected_SSL{ Mode{repr: "Z" }} }
func (m *Freenode_User_Connected_SSL) CanAdd() bool { return false }
func (m *Freenode_User_Connected_SSL) CanRemove() bool { return false }

// ------

type Freenode_Channel_Color_Filter struct {
	Mode
}

func NewFreenode_Channel_Color_Filter() *Freenode_Channel_Color_Filter { return &Freenode_Channel_Color_Filter{ Mode{repr: "c" }} }
func (m *Freenode_Channel_Color_Filter) CanAdd() bool { return false }
func (m *Freenode_Channel_Color_Filter) CanRemove() bool { return false }

type Freenode_Channel_Block_CTCP struct {
	Mode
}

func NewFreenode_Channel_Block_CTCP() *Freenode_Channel_Block_CTCP { return &Freenode_Channel_Block_CTCP{ Mode{repr: "C" }} }
func (m *Freenode_Channel_Block_CTCP) CanAdd() bool { return false }
func (m *Freenode_Channel_Block_CTCP) CanRemove() bool { return false }

type Freenode_Channel_Forward_Uninvited struct {
	Mode
}

func NewFreenode_Channel_Forward_Uninvited() *Freenode_Channel_Forward_Uninvited { return &Freenode_Channel_Forward_Uninvited{ Mode{repr: "f" }} }
func (m *Freenode_Channel_Forward_Uninvited) CanAdd() bool { return false }
func (m *Freenode_Channel_Forward_Uninvited) CanRemove() bool { return false }

type Freenode_Channel_Permit_Any_Inviter struct {
	Mode
}

func NewFreenode_Channel_Permit_Any_Inviter() *Freenode_Channel_Permit_Any_Inviter { return &Freenode_Channel_Permit_Any_Inviter{ Mode{repr: "g" }} }
func (m *Freenode_Channel_Permit_Any_Inviter) CanAdd() bool { return false }
func (m *Freenode_Channel_Permit_Any_Inviter) CanRemove() bool { return false }

type Freenode_Channel_Join_Throttle struct {
	Mode
}

func NewFreenode_Channel_Join_Throttle() *Freenode_Channel_Join_Throttle { return &Freenode_Channel_Join_Throttle{ Mode{repr: "j" }} }
func (m *Freenode_Channel_Join_Throttle) CanAdd() bool { return false }
func (m *Freenode_Channel_Join_Throttle) CanRemove() bool { return false }

type Freenode_Channel_Join_Limit struct {
	Mode
}

func NewFreenode_Channel_Join_Limit() *Freenode_Channel_Join_Limit { return &Freenode_Channel_Join_Limit{ Mode{repr: "l" }} }
func (m *Freenode_Channel_Join_Limit) CanAdd() bool { return false }
func (m *Freenode_Channel_Join_Limit) CanRemove() bool { return false }

type Freenode_Channel_Large_Ban_List struct {
	Mode
}

func NewFreenode_Channel_Large_Ban_List() *Freenode_Channel_Large_Ban_List { return &Freenode_Channel_Large_Ban_List{ Mode{repr: "L" }} }
func (m *Freenode_Channel_Large_Ban_List) CanAdd() bool { return false }
func (m *Freenode_Channel_Large_Ban_List) CanRemove() bool { return false }

type Freenode_Channel_Paranoid struct {
	Mode
}

// ??? RFC specifies "p" to be private, is this the same?
func NewFreenode_Channel_Paranoid() *Freenode_Channel_Paranoid { return &Freenode_Channel_Paranoid{ Mode{repr: "p" }} }
func (m *Freenode_Channel_Paranoid) CanAdd() bool { return false }
func (m *Freenode_Channel_Paranoid) CanRemove() bool { return false }

type Freenode_Channel_Permanent struct {
	Mode
}

func NewFreenode_Channel_Permanent() *Freenode_Channel_Permanent { return &Freenode_Channel_Permanent{ Mode{repr: "P" }} }
func (m *Freenode_Channel_Permanent) CanAdd() bool { return false }
func (m *Freenode_Channel_Permanent) CanRemove() bool { return false }

type Freenode_Channel_Block_Forwarded_Users struct {
	Mode
}

func NewFreenode_Channel_Block_Forwarded_Users() *Freenode_Channel_Block_Forwarded_Users { return &Freenode_Channel_Block_Forwarded_Users{ Mode{repr: "Q" }} }
func (m *Freenode_Channel_Block_Forwarded_Users) CanAdd() bool { return false }
func (m *Freenode_Channel_Block_Forwarded_Users) CanRemove() bool { return false }

type Freenode_Channel_Away struct {
	Mode
}

func NewFreenode_Channel_Away() *Freenode_Channel_Away { return &Freenode_Channel_Away{ Mode{repr: "c" }} }
func (m *Freenode_Channel_Away) CanAdd() bool { return false }
func (m *Freenode_Channel_Away) CanRemove() bool { return false }

type Freenode_Channel_Away struct {
	Mode
}

func NewFreenode_Channel_Away() *Freenode_Channel_Away { return &Freenode_Channel_Away{ Mode{repr: "c" }} }
func (m *Freenode_Channel_Away) CanAdd() bool { return false }
func (m *Freenode_Channel_Away) CanRemove() bool { return false }
