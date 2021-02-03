# Kubefilter

Kubefilter is a small utility to filter the output of `kubectl get xxx yyy -o yaml`. It is intended to be used as a filter of kubectl:

```
kubectl get pods nginx -o yaml | kubefilter
```

## How it works

It reads the YAML output from stdin, and outputs the filter YAML in stdout. By default filters:

* All null values
* All empty objects
* Some specified entries (like `metadata.managedFields` (no more annoying `f:` items xD) or `metadata.selfLink` among others)

> kubefilter do not have any knowledge about the meaning of the YAML, it just parses the YAML and filters some specified keys from it.

## Additional parameters

kubefilter allows following additional parameters:

* `-remove-null=true|false`: To remove or leave null values (default true)
* `-remove-empty`: To remove or leave empty valyues (`{}`) (default true)
* `-remove-owner-refs`: To remove or leave `metadata.ownerReferences` entry if found. (default false)
* `-remove-keys`: Additional comma-separated keys to remove from YAML. Must use the full qualified name (like `metadata.name`)
* -`log-level=0|8`: Log Level. 0 (default) only outputs filtered YAML. 8 outputs debug information

## Example

```
PS> kubectl get pod -n monitoring prometheus-k8s-1 -o yaml | .\out\kubefilter.exe -remove-owner-refs=true -remove-keys spec
apiVersion: v1
kind: Pod
metadata:
  generateName: prometheus-k8s-
  labels:
    app: prometheus
    app.kubernetes.io/component: prometheus
    app.kubernetes.io/name: prometheus
    app.kubernetes.io/part-of: kube-prometheus
    app.kubernetes.io/version: 2.24.0
    controller-revision-hash: prometheus-k8s-7547946ffc
    operator.prometheus.io/name: k8s
    operator.prometheus.io/shard: "0"
    prometheus: k8s
    statefulset.kubernetes.io/pod-name: prometheus-k8s-1
  name: prometheus-k8s-1
  namespace: monitoring
```

## How to build

Just download the repo and use `go build`

Go version used: 1.15

## How to contribute

Do you want to contribute? That's amazing!! :D

You can open an issue or submit a PR. Before submitting a PR please, check issues to ensure that no current work is being done to address the same issue :)

## Pending

A lot of things :) Just check issues to see future improvements and ideas... and feel free to propose yours!
