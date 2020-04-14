package go_modemmanager

import (
	"github.com/godbus/dbus/v5"
	"time"
)

// The SMS interface Defines operations and properties of a single SMS message.

const (
	SmsInterface = ModemManagerInterface + ".Sms"

	SmssObjectPath = modemManagerMainObjectPath + "SMSs"

	/* Methods */
	SmsSend  = SmsInterface + ".Send"
	SmsStore = SmsInterface + ".Store"

	/* Property */
	SmsPropertyState   = SmsInterface + ".State"   // readable   u
	SmsPropertyPduType = SmsInterface + ".PduType" //  readable   u
	SmsPropertyNumber  = SmsInterface + ".Number"  // readable   s
	SmsPropertyText    = SmsInterface + ".Text"    // readable   s
	SmsPropertyData    = SmsInterface + ".Data"    // readable   ay

	SmsPropertySMSC            = SmsInterface + ".SMSC"            // readable   s
	SmsPropertyValidity        = SmsInterface + ".Validity"        // readable   (uv)
	SmsPropertyClass           = SmsInterface + ".Class"           // readable   i
	SmsPropertyTeleserviceId   = SmsInterface + ".TeleserviceId"   // readable   u
	SmsPropertyServiceCategory = SmsInterface + ".ServiceCategory" // readable   u

	SmsPropertyDeliveryReportRequest = SmsInterface + ".DeliveryReportRequest" // readable   b
	SmsPropertyMessageReference      = SmsInterface + ".MessageReference"      // readable   u
	SmsPropertyTimestamp             = SmsInterface + ".Timestamp"             // readable   s
	SmsPropertyDischargeTimestamp    = SmsInterface + ".DischargeTimestamp"    //  readable   s
	SmsPropertyDeliveryState         = SmsInterface + ".DeliveryState"         // readable   u

	SmsPropertyStorage = SmsInterface + ".Storage" // readable   u

)

type Sms interface {
	/* METHODS */

	// Returns object path
	GetObjectPath() dbus.ObjectPath

	// If the message has not yet been sent, queue it for delivery.
	Send() error

	// This method requires a MMSmsStorage value, describing the storage where this message is to be kept; or
	// MM_SMS_STORAGE_UNKNOWN if the default storage should be used.
	Store(MMSmsStorage) error

	/* PROPERTIES */

	// A MMSmsState value, describing the state of the message.
	GetState() (MMSmsState, error)

	// A MMSmsPduType value, describing the type of PDUs used in the SMS message.
	GetPduType() (MMSmsPduType, error)

	// Number to which the message is addressed.
	GetNumber() (string, error)

	// Message text, in UTF-8.
	// When sending, if the text is larger than the limit of the technology or modem, the message will be broken
	// into multiple parts or messages.
	// Note that Text and Data are never given at the same time.
	GetText() (string, error)

	// Message data.
	// When sending, if the data is larger than the limit of the technology or modem, the message will be broken
	// into multiple parts or messages.
	// Note that Text and Data are never given at the same time.
	GetData() ([]byte, error)

	// Indicates the SMS service center number.
	// Always empty for 3GPP2/CDMA.
	GetSMSC() (string, error)

	// Indicates when the SMS expires in the SMSC.
	// This value is composed of a MMSmsValidityType key, with an associated data which contains type-specific validity information:
	// 		MM_SMS_VALIDITY_TYPE_RELATIVE: The value is the length of the validity period in minutes, given as an unsigned integer (D-Bus signature 'u').
	GetValidity() (map[MMSmsValidityType]uint32, error)

	// 3GPP message class (-1..3). -1 means class is not available or is not used for this message, otherwise the 3GPP SMS message class.
	// Always -1 for 3GPP2/CDMA.
	GetClass() (int64, error)

	// A MMSmsCdmaTeleserviceId value.
	// Always MM_SMS_CDMA_TELESERVICE_ID_UNKNOWN for 3GPP.
	GetTeleserviceId() (MMSmsCdmaTeleserviceId, error)

	// A MMSmsCdmaServiceCategory value.
	// Always MM_SMS_CDMA_SERVICE_CATEGORY_UNKNOWN for 3GPP.
	GetServiceCategory() (MMSmsCdmaServiceCategory, error)

	// #TRUE if delivery report request is required, #FALSE otherwise.
	GetDeliveryReportRequest() (bool, error)

	// Message Reference of the last PDU sent/received within this SMS.
	// If the PDU type is MM_SMS_PDU_TYPE_STATUS_REPORT, this field identifies the Message Reference of the
	// PDU associated to the status report.
	GetMessageReference() (MMSmsPduType, error)

	// Time when the first PDU of the SMS message arrived the SMSC, in ISO8601 format. This field is only applicable if
	// the PDU type is MM_SMS_PDU_TYPE_DELIVER. or MM_SMS_PDU_TYPE_STATUS_REPORT.
	GetTimestamp() (time.Time, error)

	// Time when the first PDU of the SMS message left the SMSC, in ISO8601 format.
	// This field is only applicable if the PDU type is MM_SMS_PDU_TYPE_STATUS_REPORT.
	GetDischargeTimestamp() (time.Time, error)

	// A MMSmsDeliveryState value, describing the state of the delivery reported in the Status Report message.
	// This field is only applicable if the PDU type is MM_SMS_PDU_TYPE_STATUS_REPORT.
	GetDeliveryState() (MMSmsDeliveryState, error)

	// A MMSmsStorage value, describing the storage where this message is kept.
	GetStorage() (MMSmsStorage, error)

	MarshalJSON() ([]byte, error)
}

func NewSms(objectPath dbus.ObjectPath) (Sms, error) {
	var ss sms
	return &ss, ss.init(SimInterface, objectPath)
}

type sms struct {
	dbusBase
}

func (ss sms) GetObjectPath() dbus.ObjectPath {
	return ss.obj.Path()
}

func (ss sms) Send() error {
	panic("implement me")
}

func (ss sms) Store(MMSmsStorage) error {
	panic("implement me")
}

func (ss sms) GetState() (MMSmsState, error) {
	panic("implement me")
}

func (ss sms) GetPduType() (MMSmsPduType, error) {
	panic("implement me")
}

func (ss sms) GetNumber() (string, error) {
	panic("implement me")
}

func (ss sms) GetText() (string, error) {
	panic("implement me")
}

func (ss sms) GetData() ([]byte, error) {
	panic("implement me")
}

func (ss sms) GetSMSC() (string, error) {
	panic("implement me")
}

func (ss sms) GetValidity() (map[MMSmsValidityType]uint32, error) {
	panic("implement me")
}

func (ss sms) GetClass() (int64, error) {
	panic("implement me")
}

func (ss sms) GetTeleserviceId() (MMSmsCdmaTeleserviceId, error) {
	panic("implement me")
}

func (ss sms) GetServiceCategory() (MMSmsCdmaServiceCategory, error) {
	panic("implement me")
}

func (ss sms) GetDeliveryReportRequest() (bool, error) {
	panic("implement me")
}

func (ss sms) GetMessageReference() (MMSmsPduType, error) {
	panic("implement me")
}

func (ss sms) GetTimestamp() (time.Time, error) {
	panic("implement me")
}

func (ss sms) GetDischargeTimestamp() (time.Time, error) {
	panic("implement me")
}

func (ss sms) GetDeliveryState() (MMSmsDeliveryState, error) {
	panic("implement me")
}

func (ss sms) GetStorage() (MMSmsStorage, error) {
	panic("implement me")
}

func (ss sms) MarshalJSON() ([]byte, error) {
	panic("implement me")
}
