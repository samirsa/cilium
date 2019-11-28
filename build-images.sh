#!/usr/bin/env bash
docker rmi samircilium/cilium:samir
docker rmi samircilium/operator:samir
docker rmi cilium/docker-plugin:samir
DOCKER_IMAGE_TAG=samir make docker-image
docker image prune -f
docker tag $(docker images --format "{{json . }}"  | grep -w "\"cilium\/operator\""  | jq -r '.ID') samircilium/operator:samir
docker tag $(docker images --format "{{json . }}"  | grep -w "\"cilium\/cilium\""  | jq -r '.ID') samircilium/cilium:samir
docker rmi cilium/cilium:samir
docker rmi cilium/operator:samir
docker login -u samirsa -p kVLNtYR5r0Y0
docker push samircilium/cilium:samir
docker push samircilium/operator:samir
