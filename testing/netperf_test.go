package netperf

import (
	"testing"

	"github.com/cloud-bulldozer/k8s-netperf/pkg/config"
)

// TestParseConf Test for success. Ensure we successfully parse a good config file
func TestParseConf(t *testing.T) {
	file := "test-config.yml"
	_, err := config.ParseConf(file)
	if err != nil {
		t.Fatal("Parsing config file failed")
	}
}

// TestParseConf Test for success. Ensure we successfully parse a good config file
func TestParseV2Conf(t *testing.T) {
	file := "test-v2config.yml"
	_, err := config.ParseV2Conf(file)
	if err != nil {
		t.Fatal("Parsing config file failed")
	}
}

// TestParseConf Test for success. Ensure we successfully parse the default config
func TestShippingConf(t *testing.T) {
	file := "../netperf.yml"
	_, err := config.ParseV2Conf(file)
	if err != nil {
		t.Fatal("Parsing config file failed")
	}
}

// TestMissingParseConf Testing for failure. Test profile regex
func TestMissingParseConf(t *testing.T) {
	file := "test-bad-missing-config.yml"
	_, err := config.ParseConf(file)
	if err == nil {
		t.Fatal("Parsing config file should have failed but succeeded")
	}
}

// TestMissingParseV2Conf Testing for failure. Test profile regex
func TestMissingParseV2Conf(t *testing.T) {
	file := "test-bad-missing-v2config.yml"
	_, err := config.ParseV2Conf(file)
	if err == nil {
		t.Fatal("Parsing config file should have failed but succeeded")
	}
}

// TestBadParseConf Test for failure. User leaves out a config field
func TestBadParseConf(t *testing.T) {
	file := "test-bad-profile-config.yml"
	_, err := config.ParseConf(file)
	if err == nil {
		t.Fatal("Parsing config file should have failed but succeeded")
	}
}

// TestMultiFlowConf verifies successful parsing of a multi-flow config with randsrcport.
func TestMultiFlowConf(t *testing.T) {
	file := "test-multiflow-config.yml"
	cfgs, err := config.ParseV2Conf(file)
	if err != nil {
		t.Fatalf("Parsing multi-flow config file failed: %v", err)
	}
	if len(cfgs) == 0 {
		t.Fatal("Expected at least one config, got none")
	}
	for _, c := range cfgs {
		if c.Parallelism < 100 {
			t.Errorf("Expected parallelism >= 100, got %d", c.Parallelism)
		}
		if !c.RandSrcPort {
			t.Errorf("Expected randsrcport to be true")
		}
	}
}
