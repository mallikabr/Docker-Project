Installation of Docker
---------------------------

create an ubuntu EC2 instance on AWS.<br>
Open Terminal, login into the instance<br> 
```ssh -i "key name" ubuntu@ec2instace_publicIp```<br>

and run the below commands to install docker<br>
```sudo apt update```<br>
```sudo apt install docker.io -y```<br>

Start Docker and Grant Access
---------------------------------
Ensure docker daemon is up and running <br>
```docker run hello-world```<br>

if the output is:<br>
```docker: permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/containers/create": dial unix /var/run/docker.sock: connect: permission denied.```
```See 'docker run --help'.```<br>

that means <br>
1. docker daemon is not running<br>
2. your user does not have access to run the docker commands<br>

Start docker daemon
---------------------------
Verify if docker daemon is started and active <br>
```sudo systemctl status docker```<br>

if you notice docker daemon is not running, then you can start the daemon <br>
```sudo systemctl start docker``` <br>

Grant access to your user to run the docker commands
-------------------------------------------------------
```sudo usermod -aG docker ubuntu```<br>

<strong>NOTE:</strong> logout and login back for the reflected changes<br>


Docker is installed and running
-------------------------------------
Now run the command again<br>
```docker run hello-world```<br>

output should be like :<br>
```……………..```<br>
```………………```<br>
```Hello from Docker!```<br>
```This message shows that your installation appears to be working correctly.```<br>
```………………..```<br>
```……………..```<br>

Clone the repository and move to example folder
---------------------------------------------------
```git clone https://github.com/mallikabr/Docker-Project.git```<br>
```cd Docker-Project/examples/first-docker-file/```<br>

Login to Docker [https://hub.docker.com/]
-------------------------------------------
```docker login```<br>

Output:<br>
```Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.```<br>
```Username: mallikabr```<br>
```Password: ```<br>
```WARNING! Your password will be stored unencrypted in /home/ubuntu/.docker/config.json.```<br>
```Configure a credential helper to remove this warning. See```<br>
```https://docs.docker.com/engine/reference/commandline/login/#credentials-store```<br>

```Login Succeeded```<br>

Build your first Docker image
-----------------------------------

change the username accordingly <br>
```docker build -t mallikabr/my-first-docker-image:latest .```<br>

Output:<br>

```Sending build context to Docker daemon  10.24kB```<br>
```Step 1/6 : FROM ubuntu:latest```<br>
 ```---> e34e831650c1```<br>
```Step 2/6 : WORKDIR /app```<br>
 ```---> Using cache```<br>
 ```---> b3e9c8a20487```<br>
```Step 3/6 : COPY . /app```<br>
``` ---> Using cache```<br>
``` ---> 5f6fa3dc8494```<br>
```Step 4/6 : RUN apt-get update && apt-get install -y python3 python3-pip```<br>
``` ---> Using cache```<br>
 ```---> b44cf74d2489```<br>
```Step 5/6 : ENV NAME World```<br>
 ```---> Using cache```<br>
 ```---> b454cad62dde```<br>
```Step 6/6 : CMD ["python3", "app.py"]```<br>
``` ---> Using cache```<br>
``` ---> d53368281db1```<br>
```Successfully built d53368281db1```<br>
```Successfully tagged mallikabr/my-first-docker-image:latest```<br>


Verify Docker Image is created
-------------------------------------
```docker images```<br>

Output:<br>
```REPOSITORY                         TAG       IMAGE ID       CREATED         SIZE```<br>
```mallikabr6/my-first-docker-image   latest    d53368281db1   6 minutes ago    473MB```<br>
```mallikabr/my-first-docker-image    latest    d53368281db1   6 minutes ago    473MB```<br>
```<none>                             <none>    bc5afc4a324c   19 minutes ago   473MB```<br>
```ubuntu                             latest    e34e831650c1   13 days ago      77.9MB```<br>
```hello-world                        latest    d2c94e258dcb   8 months ago     13.3kB```<br>


Run First Docker Container
--------------------------------

```docker run -it mallikabr/my-first-docker-image```<br>

Output:<br>
```Hello world```<br>

Push the image to DockerHub and share it with the world
-----------------------------------------------------------
```docker push mallikabr/my-first-docker-image```<br>

Output:<br>
```Using default tag: latest```<br>
```The push refers to repository [docker.io/mallikabr/my-first-docker-image]```<br>
```9793bbaf736e: Pushed ```<br>
```6d2c6fa6d905: Pushed ```<br>
```4c6ded832e9c: Pushed ```<br>
```8e87ff28f1b5: Pushed ```<br>
```latest: digest: sha256:ae6c71cc82a4a322fd9772c2821bfd015a6c2aabf85ba90eaa504d6390f3a048 size: 1155```<br>










