/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsServiceApplicationProfile object
type AlbDnsServiceApplicationProfile struct {
	// Respond to AAAA queries with empty response when there are only IPV4 records. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	AaaaEmptyResponse bool `json:"aaaa_empty_response,omitempty"`
	// Email address of the administrator responsible for this zone. This field is used in SOA records (rname) pertaining to all domain names specified as authoritative domain names. If not configured, the default value 'hostmaster' is used in SOA responses. Default value when not specified in API or module is interpreted by ALB Controller as hostmaster. 
	AdminEmail string `json:"admin_email,omitempty"`
	// Enable DNS query/response over TCP. This enables analytics for pass-through queries as well. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	DnsOverTcpEnabled bool `json:"dns_over_tcp_enabled,omitempty"`
	// DNS zones hosted on this Virtual Service. Maximum of 100 items allowed. 
	DnsZones []AlbDnsZone `json:"dns_zones,omitempty"`
	// Subdomain names serviced by this Virtual Service. These are configured as Ends-With semantics. Maximum of 100 items allowed. 
	DomainNames []string `json:"domain_names,omitempty"`
	// Enable stripping of EDNS client subnet (ecs) option towards client if DNS service inserts ecs option in the DNS query towards upstream servers. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	EcsStrippingEnabled bool `json:"ecs_stripping_enabled,omitempty"`
	// Enable DNS service to be aware of EDNS (Extension mechanism for DNS). EDNS extensions are parsed and shown in logs. For GSLB services, the EDNS client subnet option can be used to influence Load Balancing. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	Edns bool `json:"edns,omitempty"`
	// Specifies the IP address prefix length to use in the EDNS client subnet (ECS) option. When the incoming request does not have any ECS option and the prefix length is specified, an ECS option is inserted in the request passed to upstream server. If the incoming request already has an ECS option, the prefix length (and correspondingly the address) in the ECS option is updated, with the minimum of the prefix length present in the incoming and the configured prefix length, before passing the request to upstream server. Allowed values are 1-32. 
	EdnsClientSubnetPrefixLen int64 `json:"edns_client_subnet_prefix_len,omitempty"`
	// Drop or respond to client when the DNS service encounters an error processing a client query. By default, such a request is dropped without any response, or passed through to a passthrough pool, if configured. When set to respond, an appropriate response is sent to client, e.g. NXDOMAIN response for non-existent records, empty NOERROR response for unsupported queries, etc. Enum options - DNS_ERROR_RESPONSE_ERROR, DNS_ERROR_RESPONSE_NONE. Default value when not specified in API or module is interpreted by ALB Controller as DNS_ERROR_RESPONSE_NONE. 
	ErrorResponse string `json:"error_response,omitempty"`
	// The <domain-name>  of the name server that was the original or primary source of data for this zone. This field is used in SOA records (mname) pertaining to all domain names specified as authoritative domain names. If not configured, domain name is used as name server in SOA response. 
	NameServer string `json:"name_server,omitempty"`
	// Specifies the TTL value (in seconds) for SOA (Start of Authority) (corresponding to a authoritative domain owned by this DNS Virtual Service) record's minimum TTL served by the DNS Virtual Service. Allowed values are 0-86400. Unit is SEC. Default value when not specified in API or module is interpreted by ALB Controller as 30. 
	NegativeCachingTtl int64 `json:"negative_caching_ttl,omitempty"`
	// Specifies the number of IP addresses returned by the DNS Service. Enter 0 to return all IP addresses. Allowed values are 1-20. Special values are 0- 'Return all IP addresses'. Default value when not specified in API or module is interpreted by ALB Controller as 1. 
	NumDnsIp int64 `json:"num_dns_ip,omitempty"`
	// Specifies the TTL value (in seconds) for records served by DNS Service. Allowed values are 0-86400. Unit is SEC. Default value when not specified in API or module is interpreted by ALB Controller as 30. 
	Ttl int64 `json:"ttl,omitempty"`
}
