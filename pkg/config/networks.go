// Package config provides data structures for network configurations.
package config

// NetworkConfig represents the configuration for a network.
type NetworkConfig struct {
    Type   string `json:"type"` // Network type (e.g., SR-IOV, MacVLAN)
    Name   string `json:"name"` // Name of the network
    VLANID int    `json:"vlan_id,omitempty"` // VLAN ID for MacVLAN
    IPAM   IPAM   `json:"ipam"` // IP Address Management configuration
}

// IPAM represents the configuration for IP Address Management.
type IPAM struct {
    Gateway string   `json:"gateway"` // Gateway IP
    Subnet  string   `json:"subnet"`  // Subnet in CIDR notation
    AllocatedIPs []string `json:"allocated_ips"` // List of allocated IPs
}

// SRIOVConfig represents the configuration specific to SR-IOV.
type SRIOVConfig struct {
    ResourceName string `json:"resource_name"` // Name of the SR-IOV resource
    PFName       string `json:"pf_name"`       // Physical Function name
}

// MacVLANConfig represents the configuration specific to MacVLAN.
type MacVLANConfig struct {
    Parent string `json:"parent"` // Parent interface name for MacVLAN
}