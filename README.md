[https://gunstein.vatnar.no](https://gunstein.vatnar.no)

Edit HTML/CSS in `gunstein_vatnar_no/`.

Local build:
`sudo docker build -t mypage_server:local .`

Local test:
`sudo docker run --rm -p 8080:80 mypage_server:local`

No Docker Hub push is required for deploy.

Deploy recommendation (via `reverseproxy` repo + Traefik):
1. Push code changes to this repo (`mypage_server`).
2. On server: pull latest code in both repos.
3. In `reverseproxy/docker-compose.yml`, use local build for this service:

```yaml
mypage_server:
  build:
    context: /path/to/mypage_server
    dockerfile: Dockerfile
  image: mypage_server:local
  pull_policy: never
  container_name: "mypage_server"
  restart: always
  labels:
    - "traefik.enable=true"
    - "traefik.http.routers.mypage_server.rule=Host(`gunstein.vatnar.no`)"
    - "traefik.http.routers.mypage_server.entrypoints=websecure"
    - "traefik.http.routers.mypage_server.tls.certresolver=myresolver"
    - "traefik.http.routers.mypage_server.middlewares=test-compress"
    - "traefik.http.middlewares.test-compress.compress=true"
```

Then deploy on server:
`docker compose up -d --build mypage_server`

Security headers:
- CSP and related headers are set in this repo via Nginx config: `nginx/default.conf`.
