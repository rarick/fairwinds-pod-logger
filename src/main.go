package main

import (
  "context"
  "log"
  "os"
  "time"

  api "k8s.io/api/core/v1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/apimachinery/pkg/fields"
  "k8s.io/apimachinery/pkg/util/wait"
  corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
  "k8s.io/client-go/rest"
  "k8s.io/client-go/tools/cache"
  "k8s.io/client-go/util/retry"
)

func setPodAnnotations(pod *api.Pod) {
  ann := pod.ObjectMeta.Annotations
  if ann == nil {
    ann = make(map[string]string)
    pod.ObjectMeta.Annotations = ann
  }

  ann["fairwinds-timestamp"] = time.Now().String()
}

func podAdded(obj interface{}, kubeClient *corev1.CoreV1Client, timeAtStart *time.Time) {
  pod := obj.(*api.Pod)
  if pod.ObjectMeta.CreationTimestamp.Time.Before(*timeAtStart) {
    // Pod existed before this watcher started
    return
  }

  log.Printf("Pod created: %s/%s\n", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name)

  // TODO: Select a context
  err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
      newPod, getErr := kubeClient.Pods(pod.ObjectMeta.Namespace).Get(
        context.TODO(),
        pod.ObjectMeta.Name,
        metav1.GetOptions{},
      )

      if getErr != nil {
        panic(getErr.Error())
      }

      setPodAnnotations(newPod)

      _, updateErr := kubeClient.Pods(pod.ObjectMeta.Namespace).Update(
        context.TODO(),
        newPod,
        metav1.UpdateOptions{},
      )

      return updateErr
    },
  )

  if err != nil {
    panic(err.Error())
  }
}

func watchForPodsCreated(kubeClient *corev1.CoreV1Client, timeAtStart *time.Time) {
  // ListerWatcher for pods, watching for all fields
  // Required by cache.NewInformer
  listerWatcher := cache.NewListWatchFromClient(
    kubeClient.RESTClient(),
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
        podAdded(obj, kubeClient, timeAtStart)
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

  watchForPodsCreated(kubeClient, &timeAtStart)
}
