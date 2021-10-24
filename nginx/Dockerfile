# Dockerfile-nginx
FROM nginx:alpine


# Nginx will listen on this port
EXPOSE 443
EXPOSE 80

# Remove the default config file that
# /etc/nginx/nginx.conf includes

RUN rm /etc/nginx/conf.d/default.conf
RUN mkdir -p /var/www/html/img

# We copy the requirements file in order to install
# Python dependencies
COPY app.conf /etc/nginx/conf.d
