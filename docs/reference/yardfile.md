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

[[annotations]]
  name = "io.buildpacks.name"
  value = "nodejs"

[[import]]
  as = "npm"
  ref = "docker.io/library/npm"

[[dependency]]
  ref = "docker.io/library/image:latest"
  optional = true
  features.include = ["..."]
  features.require = ["..."]
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
2. yardfiles can define their dependencies in **dependency** table
3. each **dependency** table can have an array of **depends.on** table
4. there can be nested **depends.on** tables in a **depends.on** table
5. each yardfile can have multiple **import** tables
6. each **import** table can have fields named *as* and *ref*
    - *ref* must reference another **yardpack**
    - *as* refers to the name by which a dockerfile can access a stage or meta_args
    - imports have no effect on build process of run image
    - imports are useful for reusability purpose only
7. each yardfile can have multiple **dependency** tables
    - *ref* must refer to a remote *yardpack*
    - *optional* makes the yardpack optional to participate.
      If any build errors are thrown, the yardpack doesn't participate in build process
    - *features.include* includes a list of features that should participate in build process
    - *features.require* marks *optional* *features* to *required*
       - *features.require* features must be included in *features.include* in order to participate
8. **annotations** are list of annotations that should be added to the image of the generated *yardpack*
    - *key* refers to the key of annotation. must follow annotation key convention
    - *value* refers to the value of annotation. must be of type string
9. **metadata** is optional metadata included in yardpack image's *Result*
