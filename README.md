
[https://gunstein.vatnar.no](https://gunstein.vatnar.no)

Edit html and css

Build:
sudo docker build -t mypage_server .
sudo docker tag mypage_server:latest gunstein/mypage_server:latest

Test local before push:
sudo docker run -p 8080:80 gunstein/mypage_server

Push 
sudo docker push gunstein/mypage_server:latest

From docker-compose.yml:
  mypage_server:
    image: gunstein/mypage_server:latest
    container_name: "mypage_server"
    restart: always
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.mypage_server.rule=Host(`gunstein.vatnar.no`)"
      - "traefik.http.routers.mypage_server.entrypoints=websecure"
      - "traefik.http.routers.mypage_server.tls.certresolver=myresolver"
      # use compression
      - "traefik.http.routers.mypage_server.middlewares=test-compress"
      - "traefik.http.middlewares.test-compress.compress=true"
