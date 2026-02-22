package k8s

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExtractSecondaryNetworkIp_IPv4(t *testing.T) {
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
			Annotations: map[string]string{
				"k8s.v1.cni.cncf.io/networks-status": `[
					{
						"name": "kube-system/calico",
						"ips": ["192.168.1.10"],
						"default": true,
						"dns": {}
					},{
						"name": "netperf/test-net1",
						"interface": "eth1",
						"ips": ["10.0.0.5"],
						"mac": "3a:e8:08:89:9a:93",
						"dns": {}
					}
				]`,
			},
		},
	}

	ip, err := ExtractSecondaryNetworkIp(pod)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if ip != "10.0.0.5" {
		t.Errorf("Expected IP 10.0.0.5, got: %s", ip)
	}
}

func TestExtractSecondaryNetworkIp_IPv6Fallback(t *testing.T) {
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
			Annotations: map[string]string{
				"k8s.v1.cni.cncf.io/networks-status": `[
					{
						"name": "kube-system/calico",
						"ips": ["2a00:fbc:1170:14aa:b414:4002:0:49b8"],
						"default": true,
						"dns": {}
					},{
						"name": "namespace_name/podname-net1",
						"interface": "eth1",
						"ips": ["2a00:fbc:1270:14c0:10c:2b06:2:1"],
						"mac": "3a:e8:08:89:9a:93",
						"dns": {}
					}
				]`,
			},
		},
	}

	ip, err := ExtractSecondaryNetworkIp(pod)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if ip != "2a00:fbc:1270:14c0:10c:2b06:2:1" {
		t.Errorf("Expected IPv6 fallback IP 2a00:fbc:1270:14c0:10c:2b06:2:1, got: %s", ip)
	}
}

func TestExtractSecondaryNetworkIp_NoEth1(t *testing.T) {
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
			Annotations: map[string]string{
				"k8s.v1.cni.cncf.io/networks-status": `[
					{
						"name": "kube-system/calico",
						"ips": ["192.168.1.10"],
						"default": true,
						"dns": {}
					}
				]`,
			},
		},
	}

	_, err := ExtractSecondaryNetworkIp(pod)
	if err == nil {
		t.Fatal("Expected error when eth1 interface is not found, but got nil")
	}
}

func TestExtractSecondaryNetworkIp_MissingAnnotation(t *testing.T) {
	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "test-pod",
			Annotations: map[string]string{},
		},
	}

	_, err := ExtractSecondaryNetworkIp(pod)
	if err == nil {
		t.Fatal("Expected error when annotation is missing, but got nil")
	}
}

func TestBuildMultusNetworkAnnotations_Set(t *testing.T) {
	multusJSON := `[{"ippool": "pool1", "trust": "on", "spoofchk": "off"}]`
	annotations := buildMultusNetworkAnnotations(multusJSON)

	val, ok := annotations["robin.io/networks"]
	if !ok {
		t.Fatal("Expected robin.io/networks annotation to be set")
	}
	if val != multusJSON {
		t.Errorf("Expected annotation value %q, got %q", multusJSON, val)
	}
}

func TestBuildMultusNetworkAnnotations_Empty(t *testing.T) {
	annotations := buildMultusNetworkAnnotations("")
	if len(annotations) != 0 {
		t.Errorf("Expected empty annotations map for empty input, got %v", annotations)
	}
}
