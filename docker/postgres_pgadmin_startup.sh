#Create a directory to store the cert and key files and navigate to this folder
mkdir ~/docker/postgres_data
mkdir ~/docker/postgres_ssl
cd ~/docker/postgres_ssl
#Create our key
openssl genrsa -des3 -out server.key 1024
#Remove the passphrase from our new key and set permissions
openssl rsa -in server.key -out server.key
chmod 600 server.key
#Create our self signed cert
openssl req -new -key server.key -days 3650 -out server.crt -x509 -subj '/C=US/ST=Utah/L=Draper/O=nav.com/CN=nav.com/emailAddress=bcole@nav.com'
#Create a docker instance of postgres with mounts and arguments to use our new cert and key
docker run --name postgres-ssl \
-e POSTGRES_PASSWORD=postgres \
-e POSTGRES_USER=postgres \
-v ~/docker/postgres_data:/var/lib/postgresql/data \
-v ~/docker/postgres_ssl/server.crt:/var/lib/postgresql/server.crt:ro \
-v ~/docker/postgres_ssl/server.key:/var/lib/postgresql/server.key:ro \
-p 5432:5432 \
-d postgres:latest \
-c ssl=on \
-c ssl_cert_file=/var/lib/postgresql/server.crt \
-c ssl_key_file=/var/lib/postgresql/server.key

#create pgadmin container
mkdir -p ~/docker/pgadmin
mkdir -p ~/docker/pgadmin_config
touch ~/docker/pgadmin_config/servers.json
docker run --name pgadmin \
-e PGADMIN_DEFAULT_EMAIL=you@nav.com \
-e PGADMIN_DEFAULT_PASSWORD=postgres \
-v /Users/bcole/docker/pgadmin/config_servers.json:/servers.json \
-v /Users/bcole/docker/pgadmin:/var/lib/pgadmin \
--link postgres \
-d dpage/pgadmin4