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

docker run -v ${pwd}:/app -v /app/node_modules -p 5000:5000 -d --name node-app node-app-image

Do we need our COPY . ./ command if we are using the -v bind mount flag to sync with the folder ?
Ans : Yes. as we use the bind mount mode just in Development.For Production we definitely need COPY . ./

But : to prevent the container's writing abilities to alter our source code -->
pass this flag -->
docker run -v ${pwd}:/app:ro -v /app/node_modules -p 5000:5000 -d --name node-app node-app-image

Output : 
```
root@f8fe10ac70b7:/app# touch urmom
touch: cannot touch 'urmom': Read-only file system
```

