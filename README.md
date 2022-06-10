# dater

xunit tool  

## Overview

This tools will provide some features for manipulate, analyze and digest xunit files.

## Scope

* Analyze results for overall state (PASS or FAILED)
* Ingest xunit style on elasticsearch
  * Adapt to json
  * Aggregate some metadata
  * Integrate elastic client
  * Define elastic schemas with nested objects
* Ingest xunit to report portal
* Ingets xunit to polarion
* Adapt xunit versionig

## Schema generation

The solution relies on schemas to define the data it will handle, currently implemeting generate structs based on xsd schemas and json schemas (json or yaml representation).

The generators are embedded and can be checked at [xsd-generator](pkg/schemas/gen-xsd.go) and [json-generator](pkg/schemas/gen-json.go)  

These generators can be used on generate annotation:

```go
package schemas

//go:generate go run gen-xsd.go xunit xunit.xsd xunit xunit
```

### Xunit

There is no official xsd for xunit, we are using the on defined by xunit-plugin from jenkins-ci, accessible on [github](https://github.com/jenkinsci/xunit-plugin/blob/master/src/main/resources/org/jenkinsci/plugins/xunit/types/model/xsd/junit-10.xsd)
