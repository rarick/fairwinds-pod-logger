package main

import (
  "log"
  "os"
  "time"

  // TODO: Check if this these are right
  api "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/fields"
  "k8s.io/apimachinery/pkg/util/wait"
  "k8s.io/client-go/rest"
  "k8s.io/client-go/tools/cache"
  corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func podAddCallback(obj interface{}) {
  pod := obj.(*api.Pod)
  log.Println("Pod created: " + pod.ObjectMeta.Name)
}

func watchForPodsCreated(client cache.Getter) {
  // ListerWatcher for pods, watching for all fields
  // Required by cache.NewInformer
  listerWatcher := cache.NewListWatchFromClient(
    client,
    "pods",
    api.NamespaceAll,
    fields.Everything(),
  )

  resyncPeriod := 30 * time.Second

  _, eventController := cache.NewInformer(
    listerWatcher,
    &api.Pod{},
    resyncPeriod,
    cache.ResourceEventHandlerFuncs{
      AddFunc: podAddCallback,
    },
  )

  eventController.Run(wait.NeverStop)
}

func main() {
  log.SetOutput(os.Stdout)

  // TODO: Update how authentication is done
  config, err := rest.InClusterConfig()

  if err != nil {
    // TODO: Determine how errors should be handled
    panic(err.Error())
  }

  kubeClient, err := corev1.NewForConfig(config)

  if err != nil {
    // TODO: Determine how errors should be handled
    panic(err.Error())
  }

  watchForPodsCreated(kubeClient.RESTClient())
}
