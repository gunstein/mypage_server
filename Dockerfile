FROM nginx:stable-alpine
COPY gunstein_vatnar_no /usr/share/nginx/html
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
