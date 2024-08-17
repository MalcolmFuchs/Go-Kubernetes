package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", "/Users/malcolmfuchs/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %s, Namespace: %s\n", pod.Name, pod.Namespace)
	}

}
