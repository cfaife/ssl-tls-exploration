#!/bin/bash




echo "calling the application Build process ...."

cd ./application
./build.sh
cd ..

echo "stopping all running containers...."
declare container_array=$(docker ps -q)
for container in ${container_array[@]}
do 
    echo "stopping container $container"
    docker container stop $container
done




echo "cleaning the container processes...."
declare container_array=$(docker ps -aq)
for container in ${container_array[@]}
do 
    echo "removing container $container"
    docker container rm $container
done


echo "cleaning container images...."
images_array=$(docker image ls -q)
for image in ${images_array[@]}
do
    if [ $image != "nginx" ]; then 
        echo "removing docker image $image"
        docker image rm  $image 
    fi
done

docker-compose down

docker-compose up -d 


