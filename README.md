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
docker exec -it node-app bash

# To observe changes in your files :
Step 1 : 
- build the image again

Step 2 :(easy)

docker run -p 5000:5000 -d --name node-app node-app-image -v C:\Users\KIIT\Documents\Codes\devops-node\devops-node\:/app


or 

docker run -p 5000:5000 -d --name node-app node-app-image -v %cd%:/app

- we uninstall node_modules and run the docker container, but it fails to load 
REASON ? :
- when we use the -v flag , it Syncs the local dir with container dir , and hence the node -module installed as per the docker image gets uninstalledby syncing with local dir