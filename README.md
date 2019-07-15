# metallb-neighbour-helper

This project aims to "help" a [Kubernetes](https://kubernetes.io) cluster which uses [MetalLB](https://metallb.universe.tf/) to register
its nodes to the BGP router MetalLB talks to.

The development is best effort and mostly for fun and to solve my own problems.

Currently, this projects implements OPNsense with the FRR package. I hope to add
more providers in the future, and I have VMware vCloud in mind.

The implementation is pretty naive, it takes a look at the MetalLB configmap and
uses the AS numbers defined there.

When the service is added to the Kubernetes cluster it will check if all the
Nodes are registered with the BGP host and if not, add them. There is currently
not any smart logic to select which nodes should be registered.

After the initial registration, a watcher of the Node object in the cluster is
started and if nodes are added or deleted, the BGP host will be updated
accordingly.

There is currently not any smart logic if any of the configmaps (MetalLB or the
helper) is updated and the pod should be restarted.

Contributions and suggestions are welcome 😀

## Installation

First install MetalLB according to their [installation documentation](https://metallb.universe.tf/installation/) and [configuration documentation](https://metallb.universe.tf/configuration/).

If you install by applying the `yaml` file, the configuration in the `example/`
directory.

The `example/` configuration should be fine in most cases, but a few things
might need to be configured:

In the `metallb-helper.yaml` the following parts might need configuration:
(See comments in the file)

-   `namespace`: for the `Deployment` must be the same as MetalLB
-   `serviceAccountName`: must be the same as the `metallb-speaker` (probably `speaker` or `metallb-speaker`)
-   `args`: must be the name of the `ConfigMap` for MetalLB and the MetalLB Helper

In the `metallb-helper-configmap.yaml` the following parts might need configuration:
(See comments in the file)

-   `namespace`: for the `ConfigMap` must be the same as MetalLB
-   `peer-address`: must be the same IP/Address as the `peer-address` in the associated peer in MetalLB
