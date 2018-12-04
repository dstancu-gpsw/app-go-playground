package main

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gopro/ext-k8s"
	corev1 "github.com/gopro/ext-k8s/apis/core/v1"
	"io/ioutil"
	"log"
	"os"
)

func loadClient(kubeconfigPath string) (*k8s.Client, error) {
	data, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("read kubeconfig: %v", err)
	}

	// Unmarshal YAML into a Kubernetes config object.
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
	}
	//config.AuthInfo[0]
	return k8s.NewClient(&config)
}


func main() {
	client, err := loadClient("/Users/dstancu/.kube/eks-2-cluster.config")
	//client, err := loadClient("/Users/dstancu/.kube/config")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var nodes corev1.NodeList
	if err := client.List(context.Background(), "", &nodes); err != nil {
		log.Fatal(err)
	}
	for _, node := range nodes.Items {
		fmt.Printf("name=%q schedulable=%t\n", *node.Metadata.Name, !*node.Spec.Unschedulable)
	}
}

