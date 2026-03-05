# k8s-netperf examples

## netperf-full.yml
Example containing a full suite of tests that we run in our Continious Performance Testing env.

### Why a single stream w/ Service
Today k8s-netperf only supports a single stream through a service. 

## netperf-multiflow.yml
Example configuration for testing with 100+ flows, each using a different random source port.
This is particularly useful for ECMP (Equal-Cost Multi-Path) load-balancing tests, where you need
distinct flow 5-tuples (src_ip, src_port, dst_ip, dst_port, proto) to distribute traffic across
multiple network paths.

```shell
k8s-netperf --config examples/netperf-multiflow.yml
```
