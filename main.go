package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	metallbConfig "github.com/kradalby/metallb-neighbour-helper/metallb-config"

	"github.com/gobuffalo/envy"
	utilnode "k8s.io/kubernetes/pkg/util/node"
)

// var PRODUCTION = "production"

var DEVELOPMENT = "development"
var ENV = envy.Get("GO_ENV", DEVELOPMENT)

func main() {
	var (
		metallbNamespace           = flag.String("namespace", "metallb-system", "Namespace where MetalLB runs")
		metallbConfigMapName       = flag.String("metallb-config", "config", "Name of MetalLB configmap")
		metallbHelperConfigMapName = flag.String("metallb-helper-config", "config-helper", "Name of MetalLB Helper configmap")
	)
	flag.Parse()

	kubeClient, err := getKubernetesClient()
	if err != nil {
		log.Fatalf("[FATAL] Failed to create Kubernetes client with error: \n %s", err)
	}

	var namespace string
	namespaceFile, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Printf("[INFO] Could not detect namespace, using from cli or default")
		namespace = *metallbNamespace
	} else {
		namespace = string(namespaceFile)
	}

	log.Printf(
		"[INFO] Reading MetalLB configuration from configmap '%s' in namespace '%s'",
		*metallbConfigMapName,
		namespace,
	)
	metallbConfigMap, err := kubeClient.Client.CoreV1().ConfigMaps(namespace).Get(
		*metallbConfigMapName,
		metav1.GetOptions{},
	)
	if err != nil {
		log.Fatalf("[FATAL] Failed to read MetalLB configmap with error: \n %s", err)
	}

	log.Printf("[TRACE] Parsing MetalLB configuration \n")
	mlbConfig, err := metallbConfig.Parse([]byte(metallbConfigMap.Data["config"]))
	if err != nil {
		log.Fatalf("[FATAL] Failed to parse MetalLB config with error: \n %s", err)
	}

	log.Printf(
		"[INFO] Reading MetalLB Helper configuration from configmap '%s' in namespace '%s'",
		*metallbHelperConfigMapName,
		namespace,
	)
	metallbHelperConfigMap, err := kubeClient.Client.CoreV1().ConfigMaps(namespace).Get(
		*metallbHelperConfigMapName,
		metav1.GetOptions{},
	)
	if err != nil {
		log.Fatalf("[FATAL] Failed to read MetalLB Helper configmap with error: \n %s", err)
	}

	log.Printf("[TRACE] Parsing MetalLB Helper configuration \n")
	providers, err := Parse([]byte(metallbHelperConfigMap.Data["config"]))
	if err != nil {
		log.Fatalf("[FATAL] Failed to parse MetalLB config with error: \n %s", err)
	}

	asNumberMap := pairProvidersAndASNumbers(providers, mlbConfig.Peers)

	log.Printf("[INFO] Getting list of Nodes from Kubernetes cluster")
	nodes, err := kubeClient.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("[FATAL] Failed to get list of Kubernetes Nodes with error: \n %s", err)
	}

	for _, node := range nodes.Items {
		err := addNode(node, asNumberMap, providers)
		if err != nil {
			log.Println(err)
		}
	}

	log.Printf("[INFO] Watching Kubernetes nodes for change")
	w, err := kubeClient.Client.CoreV1().Nodes().Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("[FATAL] Failed to watch Kubernetes cluster with error: \n %s", err)
	}

	for event := range w.ResultChan() {
		switch event.Type {
		case watch.Added:
			node, ok := event.Object.(*corev1.Node)
			if !ok {
				log.Printf("[INFO] Could not cast event object to Node type: %#v", event.Object)
			} else {
				err := addNode(*node, asNumberMap, providers)
				if err != nil {
					log.Println(err)
				}

			}
		// We are not dealing with Modified for now.
		// case watch.Modified:
		// 	node, ok := event.Object.(*corev1.Node)
		// 	if !ok {
		// 		log.Fatalf(errors.New("unexpected object type, not Node"))
		// 	}
		// 	log.Printf("Modified: %#v \n", node)
		case watch.Deleted:
			node, ok := event.Object.(*corev1.Node)
			if !ok {
				log.Printf("[INFO] Could not cast event object to Node type: %#v", event.Object)
			} else {
				err := deleteNode(node, asNumberMap, providers)
				if err != nil {
					log.Println(err)
				}

			}
		}
	}

}

func getKubernetesClient() (*KubernetesClient, error) {
	if ENV == DEVELOPMENT {
		client, err := NewOutOfClusterClient(envy.Get("KUBECONFIG", "~/.kube/config"))
		if err != nil {
			return nil, err
		}

		return client, err
	}

	client, err := NewInClusterClient()
	if err != nil {
		return nil, err
	}

	return client, err
}

func addNode(node corev1.Node, asNumberMap map[BgpProvider][]uint32, providers []BgpProvider) error {
	ip, err := utilnode.GetNodeHostIP(&node)
	if err != nil {
		return fmt.Errorf("[ERROR] Could not get IP of node %s, error: %s", node.Name, err)

	}
	for _, provider := range providers {
		for _, asNumber := range asNumberMap[provider] {
			log.Printf(
				"[INFO] Adding node %s with ip %s to BGP provider %s with AS %d",
				node.Name,
				ip.String(),
				provider.Name(),
				asNumber,
			)
			err := provider.Add(ip, asNumber)
			if err != nil {
				return fmt.Errorf(
					"[ERROR] Could not add ip %s of node %s to provider %s, error: %s",
					ip.String(),
					node.Name,
					provider.Name(),
					err,
				)
			}
		}
	}
	return nil
}

func deleteNode(node *corev1.Node, asNumberMap map[BgpProvider][]uint32, providers []BgpProvider) error {
	ip, err := utilnode.GetNodeHostIP(node)
	if err != nil {
		return fmt.Errorf("[ERROR] Could not get IP of node %s, error: %s", node.Name, err)

	}
	for _, provider := range providers {
		for _, asNumber := range asNumberMap[provider] {
			log.Printf("[INFO] Deleting node %s with ip %s to BGP provider %s", node.Name, ip.String(), provider.Name())
			err := provider.Delete(ip, asNumber)
			if err != nil {
				return fmt.Errorf(
					"[ERROR] Could not delete ip %s of node %s to provider %s, error: %s",
					ip.String(),
					node.Name,
					provider.Name(),
					err,
				)
			}
		}
	}
	return nil

}

func pairProvidersAndASNumbers(providers []BgpProvider, peers []*metallbConfig.Peer) map[BgpProvider][]uint32 {
	pairs := make(map[BgpProvider][]uint32)

	for _, provider := range providers {
		log.Printf("[TRACE] Finding AS numbers associated with %s", provider.Name())

		asNumbers := []uint32{}
		for _, peer := range peers {
			if provider.PeerIP().Equal(peer.Addr) {
				log.Printf("[TRACE] Adding MetalLB AS (%d) to provider: %s", peer.MyASN, provider.Name())
				asNumbers = append(asNumbers, peer.MyASN)
			}
		}

		pairs[provider] = asNumbers
	}
	return pairs
}
