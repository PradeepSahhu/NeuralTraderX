



container_name=$(docker ps --format "{{.Names}}" | head -n 1)
echo $container_name
docker exec -it "$container_name" psql -U pradeep -d stocks



