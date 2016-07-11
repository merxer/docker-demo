vi Dockerfile
#---- inside Dockerfile ---#
FROM ubuntu:14.04
#----- end Dockerfile ---#
docker build -t  merxer/docker-demo:base .
docker images
docker login hub.docker.com
docker push  merxer/docker-demo
