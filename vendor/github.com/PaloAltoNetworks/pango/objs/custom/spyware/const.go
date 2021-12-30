package spyware

// Valid values for Entry.DefaultAction.
const (
	DefaultActionAllow       = "allow"
	DefaultActionAlert       = "alert"
	DefaultActionDrop        = "drop"
	DefaultActionResetClient = "reset-client"
	DefaultActionResetServer = "reset-server"
	DefaultActionResetBoth   = "reset-both"
	DefaultActionBlockIp     = "block-ip"
)

const (
	singular = "custom spyware object"
	plural   = "custom spyware objects"
)
