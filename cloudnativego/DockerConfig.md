  
  Use Docker as a lightweight OS virutalization technology
  See the differences between images and containers
  Docker workflow
  Dockerfile to containerize
  Run Docker image
  Use Docker Compose to build and run images
  
  
  # Docker Configuration

 ### 1. [Create Dockerfile](Dockerfile)
  
  509  docker tag cloudnativego:1.0.0 jaynejacobs/cloudnativego:1.0.0	
  510  docker images	
  512  docker login	
  513  docker push jaynejacobs/cloudnativego:1.0.0	
  514  ls	
  515  docker push jaynejacobs/cloudnativego:1.0.0	
  516  cd $HOME/Library/Containers/com.docker.docker/Data/ cd Docker.qcow2/com.docker.driver.amd64-linux	
  525  docker images --help	
  526* docker push jaynejacobs/cloudnativego:1.0.0	
  527  docker login	username
  528  docker tag cloudnativego:1.0.0 jaynejacobs/cloudnativego:1.0.0	
  529  docker push jaynejacobs/cloudnativego:1.0.0	This needs your docker hub username

  run docker container locally
  533  docker run -it 8080:8080 cloudnativego
  534  docker run -it -p 8080:8080 jaynejacobs/cloudnativego:1.0.0
  535  docker run -it -e "PORT=9090" -p 9090:9090 cloudnativego:1.0.0
  536  docker ps

##  Run as a background process and give it a name to start an stop with
  537  docker run --name cloudnativego -d -p 8080:8080 cloudnativego:1.0.0  

    docker stop cloudnativego   

    docker run --name cloudnativego --cpu-quota 50000 --memory 64m --memory-swappiness 0 -d -p 8080:8080 jaynejacobs/cloudnativego:1.0.0


## Refine the Docker file 
  502  GOOS=linux GOARCH=amd64 go build

  506  docker build -t cloudnativego:1.0.1-alpine .
  507  docker login
  510  docker tag cloudnativego:1.0.1-alpine jaynejacobs/cloudnativego:1.0.1-alpine
  511  docker push jaynejacobs/cloudnativego:1.0.1-alpine
## After creating docker-compose file
  526  docker-compose build
  527  docker images
  528  docker-compose up -d
  529  docker ps
  530  docker-compose kill
docker run -d  jaynejacobs/gin-web:1.0.1