Docker Networking
---------------------------

In a terminal:
------------------

```docker login```<br>
```docker run -d —name login nginx:latest```<br>
```docker inspect login```(Ipaddr of login: 172.17.0.3)<br>
```docker exec -it login /bin/bash```<br>

```apt update```<br>
```apt-get install iputils-ping -y```<br>
```ping -V```<br>

```ping logout_ip```<br>
This is pingable since the network is bridge.<br>


```ping finance_ip```<br>
This is not pingable since the network is secure-network<br>


In seperate terminal:
------------------------

```docker login```<br>
```docker run -d —name logout nginx:latest```<br>
```docker inspect logout``` (Ipaddr of logout: 172.17.0.4)<br>
verify Networks:bridge<br>
```docker exec -it logout /bin/bash```<br>


In another Terminal:
------------------------

```docker login```<br>
```docker network ls```<br>
```docker network create secure-network```<br>
```docker network ls```<br>

```docker run -d —name finance —network=secure-network nginx:latest```<br>
```docker ps```<br>
```docker inspect finance```(Ipaddr of finance: 172.18.0.2)<br>
verify Networks:secure-network<br>
```docker exec -it finance /bin/bash```<br>


In seperate terminal:
--------------------------

```docker login```<br>
```docker run -d —name host-demo —network=host nginx:latest```<br>
```docker inspect host-demo```<br>
verify Networks: host<br>
Ipaddress: “” —>will be ipaddr of host<br>

