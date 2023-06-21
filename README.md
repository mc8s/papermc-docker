# Papermc-Docker

This repository contains scripts to build Docker images for PaperMC.

It gets the versions and JAR's from [papermc.io](https://papermc.io)
using [papermcdl](https://github.com/jonas-be/papermcdl).

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