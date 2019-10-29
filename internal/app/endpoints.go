package app

var endpoints = make(map[string]interface{})

func GetEndpoint(service string) string {
	clientinfo := endpoints[service].(ClientInfo)
	return HttpScheme + clientinfo.Endpoint() + APIv1Prefix + "/"
}

func InitEndpoints(config *Config) {
	Configuration = config
	endpoints[ClientData] = config.Clients["Data"]
	endpoints[ClientMetadata] = config.Clients["Metadata"]
	endpoints[ClientCommand] = config.Clients["Command"]
	endpoints[ClientLogging] = config.Clients["Logging"]
	endpoints[ClientExport] = config.Clients["Export"]
	endpoints[ClientNotifications] = config.Clients["Notifications"]
	endpoints[ClientScheduler] = config.Clients["Scheduler"]
}
