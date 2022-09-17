package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var kubeconfig *string
	home := homedir.HomeDir()
	fmt.Println("HOME: ", home)
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	fmt.Println("KUBECONFIG: ", *kubeconfig)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	podlist, err := clientset.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(podlist.Items)
	for _, pod := range podlist.Items {
		fmt.Println(pod.Name)
	}
}
