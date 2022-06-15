# dater

data management tool  

## Overview

This tools will provide some features for manipulate, analyze and digest data files (xunit, umb messages, ...).

## Scope

[x] Analyze results for overall state (PASS or FAILED)  
[ ] Ingest xunit style on elasticsearch  
  [ ] Adapt to json  
  [ ] Aggregate some metadata  
  [ ] Integrate elastic client  
  [ ] Define elastic schemas with nested objects  
[ ] Adapt xunit versionig  

## Schema generation

The solution relies on schemas to define the data it will handle, currently implemeting generate structs based on xsd schemas and json schemas (json or yaml representation).

The generators are embedded and can be checked at [xsd-generator](pkg/schemas/gen-xsd.go) and [json-generator](pkg/schemas/gen-json.go)  

These generators can be used on generate annotation, both type of generators (xsd and json-schema) share the interface:

`func(sourceFolder, targetFolder, packageName)`

A sample on how to include as annotation:

```go
package schemas

//go:generate go run gen-json.go fedora-ci fedora-ci fedoraci
//go:generate go run gen-xsd.go xunit xunit xunit
```

All go files at any subfolder on pkg schema are self generated (so not pushed to scm) based on schemas

Also for json-schema the schema can be defined with `yaml` in that case generator will create a `temporary`  
folder at `targetFoder` with fixed name `generated`

Any temporary / generated file will be deleted on `make clean`  

### Xunit

There is no official xsd for xunit, we are using the on defined by xunit-plugin from jenkins-ci,  
accessible on [github](https://github.com/jenkinsci/xunit-plugin/blob/master/src/main/resources/org/jenkinsci/plugins/xunit/types/model/xsd/junit-10.xsd)
