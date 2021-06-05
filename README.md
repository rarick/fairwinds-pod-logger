# Fairwinds Pod Logger

## Installation

### Helm Chart
```
git clone https://github.com/rarick/fairwinds-pod-logger.git
cd fairwinds-pod-logger/chart
helm install <deployment-name> .
```

### Build from Source
```
git clone https://github.com/rarick/fairwinds-pod-logger.git
cd fairwinds-pod-logger/src
go install .
```

## Configuration

### Application
Application configuration will be added in the future

### Helm Chart Values
| Key              | Default                     | Description       |
|------------------|-----------------------------|-------------------|
| image.name       | rarick/fairwinds-pod-logger | Image name        |
| image.tag        | master                      | Image tag         |
| image.pullPolicy | Always                      | Image pull policy |

## Improvements

### Application
- Add configuration options via flags and/or env variables 
- Error handling, error messages
- Add interface to allow customization and extensibility of logging messages
- Allow custom annotation key / timestamp format
- Unit tests, I'm sure there's a way to test this without installing into a cluster
- Update context / timeout
- Ability to filter which namespaces are desired with an annotation, or to list from namespaces by annotation
  - This could be accomplished by using `corev1.Namespace(namespaceName)` to get a
    [NamespaceInterface](https://pkg.go.dev/k8s.io/client-go@v0.21.1/kubernetes/typed/core/v1#NamespaceInterface),
    and checking the metadata on it. Update the `podAdded` call to return if the desired annotation is not found.
    The same approach could be used with the PodInterface.
- Leader election
  - I did not look into leader election much due to time constraints, but it appears you'll want to use the
    [client-go leader-election tools](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection)

### Docker image
- Docker image user creation/improvements there
- Make container use dumb-init, print out a startup message, etc

### Helm chart
- Usage of labels
- Countless customization options

###
