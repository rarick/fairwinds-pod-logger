# Fairwinds Pod Logger

# TODO:
- Create configuration options
- Create installation instructions
- Create config instructions
- Create Dockerfile & image
  - Create configuration options

# Improvements

## Application
- Add configuration options via flags and/or env variables 
- Error handling, error messages
- Add interface to allow customization and extensibility of logging messages
- Allow custom annotation key / timestamp format
- Unit tests, I'm sure there's a way to test this without installing into a cluster
- Update context / timeout

## Docker image
- Docker image user creation/improvements there
- Make container use dumb-init, print out a startup message, etc

## Helm chart
- Unique naming of resources
- Labels

## Processes
- Iterations are slow
  - Have to build image, push, and then install
  - There must be a faster way to test kubernetes-dependent code
