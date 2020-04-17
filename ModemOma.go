package modemmanager

import (
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
)

// Paths of methods and properties
const (
	ModemOmaInterface = ModemInterface + ".Oma"

	/* Methods */
	ModemOmaSetup                         = ModemOmaInterface + ".Setup"
	ModemOmaStartClientInitiatedSession   = ModemOmaInterface + ".StartClientInitiatedSession"
	ModemOmaAcceptNetworkInitiatedSession = ModemOmaInterface + ".AcceptNetworkInitiatedSession"
	ModemOmaCancelSession                 = ModemOmaInterface + ".CancelSession"
	/* Property */
	ModemOmaPropertyFeatures                        = ModemOmaInterface + ".Features"                        // readable   u
	ModemOmaPropertyPendingNetworkInitiatedSessions = ModemOmaInterface + ".PendingNetworkInitiatedSessions" // readable   a(uu)
	ModemOmaPropertySessionType                     = ModemOmaInterface + ".SessionType"                     // readable   u
	ModemOmaPropertySessionState                    = ModemOmaInterface + ".SessionState"                    // readable   i

)

// This interface allows clients to handle device management operations as specified by the Open Mobile Alliance (OMA).
// Device management sessions are either on-demand (client-initiated), or automatically initiated by either the device
// itself or the network.
// This interface will only be available once the modem is ready to be registered in the cellular network.
// 3GPP devices will require a valid unlocked SIM card before any of the features in the interface can be used.
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
	StartClientInitiatedSession(sessionType MMOmaSessionType) error

	// Accepts or rejects a network-initiated device management session.
	// 		IN u session_id: Unique ID of the network-initiated device management session.
	// 		IN b accept: Boolean specifying whether the session is accepted or rejected.
	AcceptNetworkInitiatedSession(sessionId uint32, accept bool) error

	// Cancels the current on-going device management session.
	CancelSession() error

	// The session state changed.
	//		i old_session_state: Previous session state, given as a MMOmaSessionState.
	//		i new_session_state: Current session state, given as a MMOmaSessionState.
	//		u session_state_failed_reason: Reason of failure, given as a MMOmaSessionStateFailedReason, if session_state is MM_OMA_SESSION_STATE_FAILED.
	Subscribe() <-chan *dbus.Signal
	Unsubscribe()

	/* PROPERTIES */

	// Bitmask of MMModemOmaFeature flags, specifying which device management features are enabled or disabled.
	GetFeatures() ([]MMOmaFeature, error)

	// List of network-initiated sessions which are waiting to be accepted or rejected, given as an array of unsigned integer pairs, where:
	// 		The first integer is a MMOmaSessionType.
	// 		The second integer is the unique session ID.
	GetPendingNetworkInitiatedSessions() ([]modemOmaInitiatedSession, error)

	// Type of the current on-going device management session, given as a MMOmaSessionType.
	GetSessionType() (MMOmaSessionType, error)

	// State of the current on-going device management session, given as a MMOmaSessionState.
	GetSessionState() (MMOmaSessionState, error)
}

func NewModemOma(objectPath dbus.ObjectPath) (ModemOma, error) {
	var om modemOma
	return &om, om.init(ModemManagerInterface, objectPath)
}

type modemOma struct {
	dbusBase
	sigChan chan *dbus.Signal
}
type modemOmaInitiatedSession struct {
	SessionType MMOmaSessionType `json:"session-type"` // network-initiated session type
	SessionId   uint32           `json:"session-id"`   // network-initiated session id
}

func (mois modemOmaInitiatedSession) String() string {
	return "SessionType: " + fmt.Sprint(mois.SessionType) +
		", SessionId: " + fmt.Sprint(mois.SessionId)
}
func (om modemOma) GetObjectPath() dbus.ObjectPath {
	return om.obj.Path()
}
func (om modemOma) Setup(features []MMOmaFeature) error {
	// todo: untested
	var tmp MMOmaFeature
	return om.call(ModemOmaSetup, tmp.SliceToBitmask(features))
}

func (om modemOma) StartClientInitiatedSession(sessionType MMOmaSessionType) error {
	// todo: untested
	return om.call(ModemOmaStartClientInitiatedSession, sessionType)
}

func (om modemOma) AcceptNetworkInitiatedSession(sessionId uint32, accept bool) error {
	// todo: untested
	return om.call(ModemOmaAcceptNetworkInitiatedSession, sessionId, accept)
}

func (om modemOma) CancelSession() error {
	// todo: untested
	return om.call(ModemOmaCancelSession)
}

func (om modemOma) GetFeatures() ([]MMOmaFeature, error) {
	var tmp MMOmaFeature
	res, err := om.getUint32Property(ModemOmaPropertyFeatures)
	if err != nil {
		return nil, err
	}
	return tmp.BitmaskToSlice(res), nil

}

func (om modemOma) GetPendingNetworkInitiatedSessions() (result []modemOmaInitiatedSession, err error) {
	res, err := om.getSliceSlicePairProperty(ModemOmaPropertyPendingNetworkInitiatedSessions)
	if err != nil {
		return nil, err
	}
	for _, e := range res {
		var tmp modemOmaInitiatedSession
		sType, ok := e.GetLeft().(uint32)
		if !ok {
			return nil, errors.New("wrong type")
		}
		tmp.SessionType = MMOmaSessionType(sType)
		sId, ok := e.GetLeft().(uint32)
		if !ok {
			return nil, errors.New("wrong type")
		}
		tmp.SessionId = sId
		result = append(result, tmp)
	}
	return
}

func (om modemOma) GetSessionType() (MMOmaSessionType, error) {
	res, err := om.getUint32Property(ModemOmaPropertySessionType)
	if err != nil {
		return MmOmaSessionTypeUnknown, err
	}
	return MMOmaSessionType(res), nil
}

func (om modemOma) GetSessionState() (MMOmaSessionState, error) {
	res, err := om.getUint32Property(ModemOmaPropertySessionState)
	if err != nil {
		return MmOmaSessionStateUnknown, err
	}
	return MMOmaSessionState(res), nil
}

func (om modemOma) Subscribe() <-chan *dbus.Signal {
	if om.sigChan != nil {
		return om.sigChan
	}

	om.subscribeNamespace(ModemManagerObjectPath)
	om.sigChan = make(chan *dbus.Signal, 10)
	om.conn.Signal(om.sigChan)

	return om.sigChan
}

func (om modemOma) Unsubscribe() {
	om.conn.RemoveSignal(om.sigChan)
	om.sigChan = nil
}

func (om modemOma) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
