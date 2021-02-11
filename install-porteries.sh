#!/bin/bash

helm upgrade --install portieris --create-namespace --namespace portieris ./portieris --set IBMContainerService=false
