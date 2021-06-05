package main

import (
  "log"
  "os"
  "time"

  api "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/fields"
  "k8s.io/apimachinery/pkg/util/wait"
  "k8s.io/client-go/rest"
  "k8s.io/client-go/tools/cache"
  corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func podAdded(obj interface{}, timeAtStart *time.Time) {
  pod := obj.(*api.Pod)
  if pod.ObjectMeta.CreationTimestamp.Time.Before(*timeAtStart) {
    // Pod existed before our watcher started
    return
  }

  log.Println("Pod created: " + pod.ObjectMeta.Name)
}

func watchForPodsCreated(client cache.Getter, timeAtStart *time.Time) {
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
      AddFunc: func (obj interface{}) {
        podAdded(obj, timeAtStart)
      },
    },
  )

  eventController.Run(wait.NeverStop)
}

func main() {
  timeAtStart := time.Now()
  log.SetOutput(os.Stdout)

  config, err := rest.InClusterConfig()

  if err != nil {
    panic(err.Error())
  }

  kubeClient, err := corev1.NewForConfig(config)

  if err != nil {
    panic(err.Error())
  }

  watchForPodsCreated(kubeClient.RESTClient(), &timeAtStart)
}
