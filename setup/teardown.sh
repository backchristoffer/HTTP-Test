#!/bin/bash

#Teardown http-test
oc delete namespace http-test && oc delete sc local-path && oc delete pv http-pv