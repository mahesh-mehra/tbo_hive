# docker build
docker build --platform=linux/amd64 -t tbo_backend .

# docker run command
docker run --rm --platform=linux/amd64 tbo_backend

# docker run command on specific port
docker run --rm --platform=linux/amd64 -p 3700:3700 tbo_backend
