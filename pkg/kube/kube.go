package kube

import (
	"log"

	kubernetes "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient struct {
	Client *kubernetes.Clientset
	config *restclient.Config
}

func NewInClusterClient() (*KubernetesClient, error) {
	// creates the in-cluster coonfig
	log.Printf("[INFO] Creating InClusterConfiguration Kubernetes client\n")
	config, err := restclient.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client := &KubernetesClient{
		Client: clientset,
		config: config,
	}

	return client, nil
}

func NewOutOfClusterClient(kubeconf string) (*KubernetesClient, error) {
	log.Printf("[INFO] Creating OutOfClusterConfiguration Kubernetes client\n")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconf)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client := &KubernetesClient{
		Client: clientset,
		config: config,
	}

	return client, nil
}
