// Package config provides data structures for network configurations.
package config

// RobinNetworkEntry represents a single network entry in the robin.io/networks annotation.
// Users can specify an IP pool name and additional SR-IOV or MacVLAN configuration options.
// Example annotation value: '[{"ippool": "ippool_name", "trust": "on", "spoofchk": "off"}]'
type RobinNetworkEntry struct {
	IPPool   string `json:"ippool"`
	Trust    string `json:"trust,omitempty"`
	SpoofChk string `json:"spoofchk,omitempty"`
}