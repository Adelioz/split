# split

## Without compose and db
```
docker build -t split:latest .
```
from /split
```
docker run --rm -p 65000:65000 split
```

`--rm` - delete container after termination

`-p 65000:65000` forwarding from 65000 on local to 65000 in container

## With compose and mongodb
```
docker compose up --build -d
```

`-d` for detach from containers - run containers in background

## Fun to know

`docker ps` - running containers
`docker image ls` - images list
`docker compose up --build -d [service]` - run all containers in `docker-compose.yaml`, `--build` - build container on every up, `[service]` - name of specific service to run
`docker compose down [service]` - stop all containers, `service` - name of container (working also for `up`)
`docker compose logs -f [service]` - list all logs in [service], `-f` - follow, observe online all logs
`docker run --rm -it --network host mongo:7.0.1-rc0 mongosh -u test_user -p p4ssw0rd mongodb://localhost:27017/admin` - connect to the database directly