#!/bin/bash

nextReleaseVersion="$1"
export RELEASE=${nextReleaseVersion}
export RELEASE=$RELEASE; make build
