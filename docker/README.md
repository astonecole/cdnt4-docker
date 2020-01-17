# Docker

## [Installing Portainer](https://portainer.readthedocs.io/en/stable/index.html)

### Linux

```sh
docker volume create portainer_data

docker run -d -p 9000:9000 -p 8000:8000 --name portainer --restart always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer
```

### Windows

```sh
docker volume create portainer_data

docker run -d -p 9000:9000 -p 8000:8000 --name portainer --restart always -v \\.\pipe\docker_engine:\\.\pipe\docker_engine -v portainer_data:/data portainer/portainer
```

## Docker Machine

### Créer une nouvelle machine

```sh
docker-machine create --driver virtualbox prod1
docker-machine create --driver virtualbox prod2
docker-machine create --driver virtualbox prod3
```

### Lister les machines disponibles

```sh
docker-machine ls
```

### Se connecter à une machine

```sh
docker-machine ssh prod1
```

## Docker Swarm

### Ajouter des machine au cluster

```sh
docker swarm init --advertise-addr=192.168.99.101
```

### Lister les noeuds géré par docker swarm

```sh
docker node ls
```

## Docker Service

### Créer un réseau

```sh
docker network create --driver overlay my_network
```

### Créer un service

```sh
docker service create --name aston_service -d --replicas 6 -p 8080:80 --network aston_swarm nginx:alpine 
```

### Mettre à jour un service

```sh
docker service update -d --update-delay 5s --update-parallelism 1 aston_service
```