package dns

type DNS interface {
	CreateARecord(hostName string, rootDomain string, ipAddress string) error
}
