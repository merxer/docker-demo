docker login
git clone http://github.com/merxer/docker-demo.git
docker build -t merxer/docker-demo:0.0.2 .
docker images
docker run -d -p 80:8888 merxer/docker-demo:0.0.2

curl http://127.0.0.1
