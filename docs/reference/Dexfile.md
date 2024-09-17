# Dexfile Reference
each *yardpack* project can have multiple *Dexfile*s. these dexfiles follow offical Dockerfile syntax. but Yard parser can process these dexfiles as an Object Oriented Dockerfile. Dexfile stands for Dockerfile extendable file.

Every Dexfile must and should have one run-image stage. example: a `FROM` instruction without `AS`

```Dockerfile
# syntax=ctryard/yardfile:1

IMPORT FROM paketobuildpacks/node-engine AS node-engine

# PUBLIC META_ARGS are accessible through IMPORTs
PUBLIC ARG META_ARG = "I am accessible through imports"
# META_ARGS that are not public are by default are [PRIVATE] and are not accessible through IMPORTs
ARG META_ARG = "I am not accessible through imports"

FROM SCRATCH AS app-img
COPY . /workspace

FROM app-img AS detect-phase
WORKDIR ./workspace
# detect.sh should write file(s) to pass data to other stages
RUN ./detect.sh

FROM nodejs:alpine as node-builder
COPY --from=app-img ./workspace .
RUN ["npm", "run", "build"]

FROM nodejs:18 as test-builder
COPY --from=app-img ./workspace .
RUN ["npm", "run", "test"]

# run image
FROM alpine
COPY --from=node-builder ./dist .
ENTRYPOINT ["node", "server.js"]
```

## Documentation

The following New Instructions are added in addition to current Dockerfile instructions


### FROM
FROM instruction is used to create a New Stage. Same as Dockerfile
#### `--platform`
Specifies the platform to build on.

#### Example
```Dockerfile
FROM --platform=linux/amd64 alpine AS build-img

FROM docker-image://docker.io/library/alpine:latest

FROM local:///directory as local-dir

FROM git://github.com/buildpacks/pack AS pack

FROM oci-layout:///directory AS oci-layout

FROM https://registry.buildpacks.io/some-ref-to-image AS image
```

### MERGE
MERGE instruction is used to create a New Stage by merging 2 or more Stages into a Signle stage.
#### `--platform`
specifies the platform to build on.

#### Example
```Dockerfile
MERGE stage1 stage2 AS merge-stage-example-1
# same as
# FROM stage1 AS merge-stage-example-1
# COPY --from=stage --link / /

MERGE stage1 MERGE stage2 stage3 AS merge-stage-example-2

MERGE stage1 FROM stage2 AS merge-stage-example-3

MERGE stage 1 DIFF stage2 stage3 AS merge-stage-example-3
```
for more infermation refer official (docs)[https://github.com/moby/buildkit/blob/master/docs/dev/merge-diff.md#mergeop]

### DIFF
DIFF instruction can be used for rebasing images. The intuition is that it returns a stage whose contents are the difference between lower stage and upper stage

#### Example
```Dockerfile
DIFF lower upper AS diff-stage

DIFF alpine node:alpine AS node
```

### LOCAL
LOCAL instruction can be used to create a new Stage,same as `FROM` with `local://`. 
#### Example
```Dockerfile
LOCAL ./directory AS dir

LOCAL --oci --platform=linux/amd64 ./directory AS oci-dir
```

### GIT
GIT is used to clone git repos as a new stage

#### Example
```Dockerfile
GIT CLONE github.com/buildpacks/pack.git AS pack
```

### HTTP
create A new stage from the given url resource. same as `FROM` with `https://` or `http://`etc...,

#### Example
```Dockerfile
HTTP https://some-link.to/aws/object-storage/image.tar AS image
```

### IMPORT
Import other dexfiles into cuurent dexfile

#### Example
```Dockerfile
IMPORT docker.io/paketobuildpacks/some-image AS some-image

FROM alpine
# some-image has a meta arg `hello`
ARG hello = some-image::hello

FROM some-image::some-stage AS some-stage
```

### PUBLIC
used to expose meta args and stages via imports

#### Example
```Dockerfile
# assume this dexfile is imported as [some-image]

# accessible in other dexfiles as [some-image::meta_arg]
PUBLIC ARG meta_arg = "meta arg value"

# cannot be imported
ARG private_arg = "i am not exposed"

# can be imported as [some-image::alpine]
PUBLIC FROM alpine AS alpine

# cannot be imported
FROM alpine AS alpine-private
```

### IF
used for conditional execute Instructions

#### Example
```Dockerfile
IMPORT docker.io/paketobuildpacks/buildpack AS bp

# EQ, NEQ, CONTAINS, AND, OR etc.., can be used
IF bp::meta-arg EQ alpine THEN
  FROM node:alpine
ELSE
  FROM node:latest
FI

IF STAT_FILE ./some-file FROM some-stage THEN
  RUN echo "detect phase passed, the file exists"
FI

IF READ_FILE ./some-file FROM some-image AS some-arg THEN
  RUN echo "content of ./some-file: $some-arg"
FI
```

### FOR
to loop over certain values

#### Example
```Dockerfile
IMPORT dockerfile.io/paketobuildpacks/bp AS bp

FOR item IN bp::some-array \
  RUN echo $item \
DONE

READ_DIR ./some-dir AS files
FOR filename IN files
  IF filename IS FILE THEN
    READ_FILE filename AS file-data
    RUN echo "content of file: $file-data"
  ELSE
    RUN echo "$filename is dir"
  FI
DONE
```

### UNSET
used to unset args and envs

#### Example

```Dockerfile
ARG some-arg = "some-value"
# prints ["some-value"]
RUN echo $some-arg

UNSET some-arg

# prints [""]
# no more arg defined. defaults to empty string if not defined  [ARG] or [ENV]s accessed
RUN echo $some-arg
```
