# Papermc-Docker

This repository contains scripts to build Docker images for PaperMC.

It gets the versions and JAR's from [papermc.io](https://papermc.io)
using [papermcdl](https://github.com/jonas-be/papermcdl).

## Docker Hub

The projects are available on the [Mc8s Docker Hub](https://hub.docker.com/u/mc8s) profile.
Each papermc project has its own repository on docker hub.
The versions are updated everytime the pipeline runs, and it finds a new build for this version.

### Available Projects

- [x] [Paper](https://hub.docker.com/r/mc8s/paper)
- [X] [Velocity](https://hub.docker.com/r/mc8s/velocity)
- [X] [Folia](https://hub.docker.com/r/mc8s/folia)
- [X] [Waterfall](https://hub.docker.com/r/mc8s/waterfall)
- [X] [Travertine](https://hub.docker.com/r/mc8s/travertine)

## Pipeline

The pipeline is configured to build hourly, at round about the 50th minute.

## Usage

You have to set the environment variable `PROJECT` or give the project over the flag `project`,
to set the project you want to build the images for.

````bash
papermc-docker --project=waterfall
````

or

````bash
export PROJECT=waterfall
papermc-docker
````