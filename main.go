package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Label struct {
	Key   string `toml:"key"`
	Value string `toml:"value"`
}

type Container struct {
	Image string `toml:"image"`
	Tag   string `toml:"tag"`
	Name  string `toml:"name"`
	Port  []map[string]interface{}
}

type Deployment struct {
	Replicas  int32 `toml:"replicas"`
	Name      string
	Label     []Label
	Container []Container
}

var dply Deployment

func ParseFile(file_path string) {

	_, err := toml.DecodeFile(file_path, &dply)

	if err != nil {
		fmt.Println("Error while parsing toml data")
		log.Fatal(err)
	}

	fmt.Printf("Replicas: %v\n", dply.Replicas)
	fmt.Printf("Labels: %v\n", dply.Label)
	fmt.Printf("Containers: %v\n", dply.Container)

}

func int32Ptr(i int32) *int32 { return &i }

func main() {
	var kubeconfig *string
	//if the user has a homedir, try to read the kubeconfig from there first
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	ParseFile("dply.toml")

	// build the client config from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	lblSelectorMap := make(map[string]string)
	for i, _ := range dply.Label {
		lblSelectorMap[dply.Label[i].Key] = dply.Label[i].Value
	}
	ports := make([]apiv1.ContainerPort, 0)
	containers := make([]apiv1.Container, 0)
	for i, _ := range dply.Container {
		for j, _ := range dply.Container[i].Port {
			port := apiv1.ContainerPort{
				Name:          dply.Container[i].Port[j]["name"].(string),
				ContainerPort: int32(dply.Container[i].Port[j]["portnum"].(int64)),
				Protocol:      apiv1.ProtocolTCP,
			}
			ports = append(ports, port)
		}
		container := apiv1.Container{
			Name:  dply.Container[i].Name,
			Image: dply.Container[i].Image + ":" + dply.Container[i].Tag,
			Ports: ports,
		}
		containers = append(containers, container)
	}
	objectMeta := metav1.ObjectMeta{
		Name: dply.Name,
	}
	dplySpec := appsv1.DeploymentSpec{
		Replicas: int32Ptr(dply.Replicas),
		Selector: &metav1.LabelSelector{
			MatchLabels: lblSelectorMap,
		},
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: lblSelectorMap,
			},
			Spec: apiv1.PodSpec{
				Containers: containers,
			},
		},
	}

	// create a deployment clinet to deploy to the default namespace
	dplyClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	// create a deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: objectMeta,
		Spec:       dplySpec,
	}


	result, err := dplyClient.Create(deployment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created deployment %v\n", result.GetObjectMeta().GetName())
}
