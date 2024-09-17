# Yardfile

Yardfile is a configuration file that Yardfile frontend uses to build docker images from source(s2i).
```toml
# syntax=ctryard/yardfile:1
[io.buildpacks]
  name = "nodejs"
  ref = "docker.io/ctryard/nodejs"
  tags = ["latest", "lts", "1"]
  homepage = "github.com/ctryard/nodejs"
  licence = "MIT"
  description = "a nodejs buildpack to build nodejs app"
  dockerfile = "."

[[feature]]
    name = "some-stage" # some-stage is registred as a feature. this feature is included iff it is included in [features.include] by any yardfile's [dependency]
    optional = true # build process has no effect even when the give stage("some-stage") fails.

[[dependency]]
  ref = "docker.io/library/image:latest"
  optional = true
  features.include = ["..."]
  features.require = ["..."] # optional stages in [docker.io/library/image:latest] becomes required.
  [[depends.on]]
    ref = "docker.io/library/image:latest"
    optional = true
    features.include = ["..."]
    features.require = ["..."]
    [[depends.on]]
      ref = "docker.io/library/alpine:latest"
      optional = false
      features.include = ["..."]
      features.require = ["..."]

# type=any
[[metadata]]
```

1. yardfile is a configuration file
2. it is used to alter the behaviour of run-images at each vertex of any dexfile.
3. yardfiles can define their dependencies in **dependency** table
4. each **dependency** table can have an array of **depends.on** table
5. there can be nested **depends.on** tables in a **depends.on** table
6. each yardfile can have multiple **dependency** tables
    - *ref* must refer to a *buildpack*
    - *optional* makes the buildpack optional to participate.
      If any build errors are thrown, the buildpack doesn't participate in build process
    - *features.include* includes a list of features that should participate in build process
    - *features.require* marks *optional* *features* to *required*
       - *features.require* features must be included in *features.include* in order to participate
7. **metadata** is optional metadata included in buildpack image's *Result*
