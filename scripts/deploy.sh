# This script is designed to be run via SSH by GitHub Actions on the VPS
# This pulls the docker image of the repo and run it in the VPS
set -euo pipefail

APP_NAME=todo-api
IMAGE=${DOCKER_USERNAME}/todo-api:latest

echo "[+] Pulling image $IMAGE"
docker pull $IMAGE

echo "[+] Restarting container"
docker stop $APP_NAME || true
docker rm $APP_NAME || true
docker run -d \
  --name $APP_NAME \
  --env-file ~/services/todo-api/.env \
  -p 5001:5001 \
  --restart always \
  $IMAGE
