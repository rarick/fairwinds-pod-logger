# Fairwinds Pod Logger

# TODO:
- Create go package
  - Make a reader for the kubernetes API
  - Make a printer
  - Look into Golang loggers
  - Create configuration options
- Create helm chart
- Create installation instructions
- Create config instructions
- Create Dockerfile & image
  - Create configuration options

# Improvements
- Unique naming of resources
- Countless helm chart customization improvements
- Docker image user creation/improvements there
- Trim down contents of container
  - dockerignore, go.sum, etc
- Make container use dumb-init and print out a startup message
- Separate build/run images
