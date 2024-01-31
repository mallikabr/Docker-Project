Volumes
---------------------------
Run the docker container by executing below commands: <br>
```Docker login```<br>
```cd Docker-Project/examples/first-docker-file/```<br>
```docker build -t mallikabr/my-first-docker-image:latest .```<br>
```docker images```<br>
```docker run -it mallikabr/my-first-docker-image```<br>

Create and manage volumes
-----------------------------

Create a volume:<br>
```docker volume create mallika```<br><br>
List the volumes:<br>
```docker volume ls```<br><br>
Start a container with a volume:<br>
```docker run -d —mount source=mallika,target=/app nginx:latest```<br><br>
List the running dockers:<br>
```docker ps```<br><br>
Verify that docker created the volume and it mounted correctly.<br>
```docker inspect CONTAINER_ID```<br>
check the "Mounts” section.<br><br>

To remove a volume:<br>
```docker volume rm mallika```<br>
