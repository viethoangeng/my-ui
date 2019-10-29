package app

const (
	ClientData          = "data"
	ClientMetadata      = "metadata"
	ClientLogging       = "logging"
	ClientCommand       = "command"
	ClientExport        = "export"
	ClientNotifications = "notifications"
	ClientScheduler     = "scheduler"

	APIv1Prefix    = "/api/v1"
	Colon          = ":"
	HttpScheme     = "http://"
	HttpProto      = "HTTP"
	StatusResponse = "pong"
)

/* ---------------- URL PARAM NAMES -----------------------*/
const (
	ID               = "id"
	NAME             = "name"
	DEVICEIDURLPARAM = "{deviceId}"
	OPSTATE          = "opstate"
	URLADMINSTATE    = "adminstate"
	ADMINSTATE       = "adminState"
	YAML             = "yaml"
	COMMANDLIST      = "commandlist"
	COMMAND          = "command"
	COMMANDID        = "commandid"
	COMMANDNAME      = "commandname"
	DEVICE           = "device"
	DEVICENAME       = "devicename"
	KEY              = "key"
	VALUE            = "value"
	PINGRESPONSE     = "pong"
	CONTENTTYPE      = "Content-Type"
	APPLICATIONJSON  = "application/json"
	TEXTPLAIN        = "text/plain"
	UNLOCKED         = "UNLOCKED"
	ENABLED          = "ENABLED"

	PING             = "ping"
	METRIC           = "metric"
	CONFIG           = "config"
	OBJECT           = "object"
	OBJECTNAME       = "objectname"
	LABEL            = "label"
	TYPE             = "type"
	TYPENAME         = "typename"
	OPERATION        = "operation"
	SERVICENAME      = "servicename"
	UPDATENAME       = "updatename"
	UPDATEPROFILE    = "updateprofile"
	UPDATEADMINSTATE = "updateadminstate"
	UPDATEPROTOCOL   = "updateprotocol"
	PROFILE          = "profile"
	PROFILENAME      = "profilename"
	UPLOAD           = "upload"
	READING          = "reading"
	READINGNAME      = "readingname"
	START            = "start"
	END              = "end"
	LIMIT            = "limit"
	VALUEDESCRIPTOR  = "valuedescriptor"
)

/* import "github.com/edgexfoundry/go-mod-core-contracts/clients" */
// Constants related to defined routes in the service APIs
const (
	AppApiGateway = "/gateway"

	// AppApiServicesManager          = "/services-manager"
	// AppApiServicesManagerOperation = AppApiServicesManager + "/" + OPERATION
	// AppApiServicesManagerMetric    = AppApiServicesManager + "/" + METRIC + "/{" + SERVICENAME + "}"
	// AppApiServicesManagerConfig    = AppApiServicesManager + "/" + CONFIG + "/{" + SERVICENAME + "}"

	// AppApiObjectsManager                 = "/objects-manager"
	// AppApiObjectsManagerType             = AppApiObjectsManager + "/" + TYPE + "/{" + TYPENAME + "}"
	// AppApiObjectsManagerObject           = AppApiObjectsManagerType + "/" + OBJECT + "/{" + OBJECTNAME + "}"
	// AppApiObjectsManagerUpdateName       = AppApiObjectsManagerObject + "/" + UPDATENAME
	// AppApiObjectsManagerUpdateProfile    = AppApiObjectsManagerObject + "/" + UPDATEPROFILE
	// AppApiObjectsManagerUpdateAdminstate = AppApiObjectsManagerObject + "/" + UPDATEADMINSTATE
	// AppApiObjectsManagerUpdateProtocol   = AppApiObjectsManagerObject + "/" + UPDATEPROTOCOL
	// AppApiObjectsManagerProfileUpload    = AppApiObjectsManager + "/" + PROFILE + "/" + UPLOAD
	// AppApiObjectsManagerProfileName      = AppApiObjectsManager + "/" + PROFILE + "/" + NAME + "/{" + PROFILENAME + "}"

	// AppApiObjectsCommand              = "/objects-command"
	// AppApiObjectsCommandObject        = AppApiObjectsCommand + "/" + TYPE + "/{" + TYPENAME + "}" + "/" + OBJECT + "/{" + OBJECTNAME + "}"
	// AppApiObjectsCommandObjectCommand = AppApiObjectsCommandObject + "/" + COMMAND + "/{" + COMMANDNAME + "}"
	// AppApiObjectsCommandObjectList    = AppApiObjectsCommandObject + "/" + COMMANDLIST

	AppApiMonitor              = "/monitor"
	AppApiMonitorByReadingName = AppApiMonitor + "/" + READING + "/" + READINGNAME + "/{" + READINGNAME + "}/{" + START + "}/{" + END + "}/{" + LIMIT + "}"
	AppApiMonitorByDeviceName  = AppApiMonitor + "/" + READING + "/" + DEVICENAME + "/{" + DEVICENAME + "}/{" + START + "}/{" + END + "}/{" + LIMIT + "}"
	// AppApiMonitorByDevice        = AppApiMonitor + "/" + READING + "/" + DEVICE + "/{" + DEVICENAME + "}/{" + LIMIT + "}"
	AppApiMonitorValueDescriptorByName       = AppApiMonitor + "/" + VALUEDESCRIPTOR + "/" + NAME + "/{" + NAME + "}"
	AppApiMonitorValueDescriptorByDeviceName = AppApiMonitor + "/" + VALUEDESCRIPTOR + "/" + DEVICENAME + "/{" + DEVICENAME + "}"

	//---------------------------------
	// AppApiSchedule = "/schedule"

	ApiBase                    = "/api/v1"
	ApiAddressableRoute        = "/api/v1/addressable"
	ApiCallbackRoute           = "/api/v1/callback"
	ApiCommandRoute            = "/api/v1/command"
	ApiConfigRoute             = "/api/v1/config"
	ApiDeviceRoute             = "/api/v1/device"
	ApiDeviceProfileRoute      = "/api/v1/deviceprofile"
	ApiDeviceServiceRoute      = "/api/v1/deviceservice"
	ApiEventRoute              = "/api/v1/event"
	ApiLoggingRoute            = "/api/v1/logs"
	ApiMetricsRoute            = "/api/v1/metrics"
	ApiNotificationRoute       = "/api/v1/notification"
	ApiNotifyRegistrationRoute = "/api/v1/notify/registrations"
	ApiPingRoute               = "/api/v1/ping"
	ApiProvisionWatcherRoute   = "/api/v1/provisionwatcher"
	ApiReadingRoute            = "/api/v1/reading"
	ApiRegistrationRoute       = "/api/v1/registration"
	ApiRegistrationByNameRoute = ApiRegistrationRoute + "/name"
	ApiSubscriptionRoute       = "/api/v1/subscription"
	ApiTransmissionRoute       = "/api/v1/transmission"
	ApiValueDescriptorRoute    = "/api/v1/valuedescriptor"
	ApiIntervalRoute           = "/api/v1/interval"
	ApiIntervalActionRoute     = "/api/v1/intervalaction"
)

var (
	Configuration *Config
	MapObjectName = make(map[string]string)
)
