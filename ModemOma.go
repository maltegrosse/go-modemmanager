package go_modemmanager

import "github.com/godbus/dbus/v5"

// This interface allows clients to handle device management operations as specified by the Open Mobile Alliance (OMA).
// Device management sessions are either on-demand (client-initiated), or automatically initiated by either the device
// itself or the network.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.

const (
	ModemOmaInterface = ModemInterface + ".Oma"

	/* Methods */
	ModemOmaSetup = ModemOmaInterface + ".Setup"
	ModemOmaStartClientInitiatedSession = ModemOmaInterface + ".StartClientInitiatedSession"
	ModemOmaAcceptNetworkInitiatedSession = ModemOmaInterface + ".AcceptNetworkInitiatedSession"
	ModemOmaCancelSession = ModemOmaInterface + ".CancelSession"
	/* Property */
	ModemOmaPropertyFeatures =  ModemOmaInterface + ".Features" // readable   u
	ModemOmaPropertyPendingNetworkInitiatedSessions =  ModemOmaInterface + ".PendingNetworkInitiatedSessions" // readable   a(uu)
	ModemOmaPropertySessionType =  ModemOmaInterface + ".SessionType" // readable   u
	ModemOmaPropertySessionState =  ModemOmaInterface + ".SessionState" // readable   i

)

type ModemOma interface {
	/* METHODS */
	// Returns object path
	GetObjectPath() dbus.ObjectPath
	MarshalJSON() ([]byte, error)

	// Configures which OMA device management features should be enabled.
	// Bitmask of MMModemOmaFeature flags, specifying which device management
	// features should get enabled or disabled. MM_OMA_FEATURE_NONE will disable all features.
	Setup(features []MMOmaFeature) error

	// Starts a client-initiated device management session.
	// Type of client-initiated device management session,given as a MMModemOmaSessionType
	StartClientInitiatedSession (sessionType MMOmaSessionType) error

	// Accepts or rejects a network-initiated device management session.
	// 		IN u session_id: Unique ID of the network-initiated device management session.
	// 		IN b accept: Boolean specifying whether the session is accepted or rejected.
	AcceptNetworkInitiatedSession(sessionId uint32, accept bool) error

	// Cancels the current on-going device management session.
	CancelSession () error

	// The session state changed.
	//		i old_session_state: Previous session state, given as a MMOmaSessionState.
	//		i new_session_state: Current session state, given as a MMOmaSessionState.
	//		u session_state_failed_reason: Reason of failure, given as a MMOmaSessionStateFailedReason, if session_state is MM_OMA_SESSION_STATE_FAILED.
	Subscribe() <-chan *dbus.Signal
	Unsubscribe()

	/* PROPERTIES */

}
// todo add properties & implement
func NewModemOma(objectPath dbus.ObjectPath) (ModemOma, error) {
	var om modemOma
	return &om, om.init(ModemManagerInterface, objectPath)
}

type modemOma struct {
	dbusBase
}

func (om modemOma) Setup(features []MMOmaFeature) error {
	panic("implement me")
}

func (om modemOma) StartClientInitiatedSession(sessionType MMOmaSessionType) error {
	panic("implement me")
}

func (om modemOma) AcceptNetworkInitiatedSession(sessionId uint32, accept bool) error {
	panic("implement me")
}

func (om modemOma) CancelSession() error {
	panic("implement me")
}

func (om modemOma) Subscribe() <-chan *dbus.Signal {
	panic("implement me")
}

func (om modemOma) Unsubscribe() {
	panic("implement me")
}

func (m modemOma) GetObjectPath() dbus.ObjectPath {
	panic("implement me")
}

func (m modemOma) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

