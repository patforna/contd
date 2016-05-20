# can't mount volumes using docker-machine =/ copy the files over to the host machine
docker-machine scp ./logstash.conf droplet:/logstash.conf
docker-machine scp ./nginx.conf droplet:/nginx.conf

# create data volume (should happen using docker compose)
# and copy Hello.java into it (for now...)

# docker create -v /input -v /output -v /logs --name data tianon/true
docker cp Hello.java data:/input
# docker-machine scp -r ./Hello.java droplet:Hello.java
