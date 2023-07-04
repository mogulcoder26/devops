# devops-node

docker build .
docker image ls 
docker image rm <image id >
docker build -t node-app-image
docker image ls 
docker run -d --name node-app  node-app-image
docker ps
docker rm node-app -f
docker run -p 5000:5000 -d --name node-app2 node-app-image