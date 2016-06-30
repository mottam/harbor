# Harbor
## Description

Harbor is a wrapper for running commands and downloading file dependencies (currently only from AWS S3) before Docker image building.

Using Harbor should be simple and Harbor should help to stop usage of customized scripts run before a `docker build`.

## Objectives
At the time, Harbor main objectives are:

+ Manage and download configuration files that don't belong in code repositories.
 + Manage per environment configuration such as: downloading different files for dev, test or production environments.
+ Execute shell commands before a `docker build` run (such as running some `ant` or `maven` build).
+ Execute `docker build`, `docker tag` and `docker push` to repository

## Usage

Harbor takes a YAML configuration file with the following structure.

```yaml
 imagetag: <tag to be used on 'docker build'>
 tags:
   - <YAML array of custom tags to create and push into registry>
 downloadpath: <local root path to download files into>
 s3:
   bucket: <base bucket to download files from>
   basepath: <inside the bucket the root path for files to be downloaded>
   region: <[optional] region of the bucket, default us-east-1>
 files:
   - s3path: <path to file in S3 after [s3.bucket]/[s3.basepath]>
     filename: <local path + name of the file, will be downloaded into [downloadpath]/[localname]>
     permission: <[optional] file permissions, default 0644>
 commands:
   - <YAML array containing shell commands (currently /bin/bash) to be run before 'docker build'>
```
 
You can use `${<KEY>}` as a placeholder in harbor.yml to be replaced by the value passed in a -e flag

By default, it looks up a file named `harbor.yml` in the current directory, but you can specify another path. 

```
Usage:
  harbor -h | --help
  harbor --version
  harbor --list-variables [--project-path <path>] [--config <name>]
  harbor [-e KEY=VALUE]... [options]
  harbor [options]

Options:
  -h, --help                    Show this screen.
  -v, --version                 Show version.
  --config <name>               Path to config file. By default, Harbor looks up for 'harbor.yml' in the current directory, or in the project path (when --project-path is passed).
  --project-path <path>         Project source files path.
  --list-variables              Parses harbor.yml and prints out every ${KEY} found.
  -e KEY=VALUE                  Replaces every ${KEY} in harbor.yml with VALUE
  --debug                       Dry-run and print command executions.
  --no-download                 Prevents downloading files from S3.
  --no-commands                 Prevents commands in harbor.yml from being run.
  --no-docker                   Do not run 'docker build', 'docker tag' and 'docker push' commands.
  --no-docker-push              Do not run 'docker push' after building and tagging images.
  --docker-opts=<DOCKER_OPTS>   Will be appended to the base docker command (ex.: 'docker <docker-opts> command').
  --no-latest-tag               Do not build image tagging as 'latest',
                                will use the first tag in 'tags' list from harbor.yml or
                                generate a timestamped tag if no 'tags' list exists.`
```

### Templating in harbor.yml
You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag

## Building (Linux/MAC)
- Install Go >= 1.5
- Clone this repository

`git clone https://github.com/elo7/harbor.git`
- Run the following commands
```
cd harbor
export GOPATH=$(pwd)
cd harbor/src/github.com/elo7/harbor
```
- Set _GOOS_ and _GOARCH_ variables according to your [plataform](https://golang.org/doc/install/source#environment) and run the following command
```
env GOOS=<operating_system> GOARCH=<architecture> go build -v -o <output_filename>
