# Dexfile Reference
each *yardpack* project can have multiple *Dexfile*s. these dexfiles follow offical Dockerfile syntax. but Yard parser can process these dexfiles as an Object Oriented Dockerfile. Dexfile stands for Dockerfile extendable file.

```Dockerfile
# syntax=docker/dockerfile:1

# PUBLIC prefix refers to access modifier. i.e. any yardpack an access this meta arg from an [[import]]
ARG PUBLIC_META_ARG
# PRIVATE prefix refers to access modifier. i.e. not accessible outside yardpack
ARG PRIVATE_META_ARG
# DEFAULT access modifier. i.e. any yardpack with same registry can access this meta arg from an [[import]]
ARG META_ARG
# CONST meta args cannot be modified once assigned
ARG CONST_META_ARG

FROM SCRATCH AS node-detect-phase
# DETECT is reserved ARG, pointing to comma separated build-stage names
ARG DETECT = "node-builder, fake-builder"
# FEATURE is a reserved ARG, indicates this detect phase is a feature, that must be included in yardfile to participate
ARG FEATURE = false
# OPTIONAL is a reserved ARG, indicates this stage is an optional stage, if any error is
# detected any dependent stages along with the current stage are ignored from build process
ARG OPTIONAL = true
# CONST is a reserved ARG, indicating the stage as a const, no other stage can extend current stage
ARG CONST = false
RUN ./detect.sh

FROM nodejs:latest as node-builder
RUN ["npm", "run", "build"]

FROM nodejs:18 as fake-builder

# npm imported from yardfile
FROM npm:stage AS a-stage-named-stage-from-npm-yardpack

# run stage
FROM alpine
COPY --from=node-builder ./dist .
ENTRYPOINT ["node", "server.js"]
```
