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

## Xunit

There is no official xsd for xunit, we are using the on defined by xunit-plugin from jenkins-ci, accessible on [github](https://github.com/jenkinsci/xunit-plugin/blob/master/src/main/resources/org/jenkinsci/plugins/xunit/types/model/xsd/junit-10.xsd)
