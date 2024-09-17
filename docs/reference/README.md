## REFERENCE

each Buildpack can have multiple Dexfiles but there should be only one Yardfile.
This will eliminate the need of 
1. component buildpacks 
2. composite buildpacks
3. inline buildpacks
4. extentions
5. lifecycle
   
while replacing all the above features to Yardfiles and Dexfiles with the help of buildkit.

```shell
Project Directory
├── Yardfile
├── Dexfile1
├── Dexfile2
├── ...
├── Subdirectory1
│   ├── Dexfile1
│   ├── Dexfile2
│   ├── Dexfile3
│   ├── ...
├── Subdirectory2
│   ├── Dexfile1
│   ├── Dexfile2
│   ├── Dexfile3
│   ├── ...
└── ...
```
