FROM php:8.1-fpm-alpine
RUN apk add --update --no-cache nginx oniguruma-dev postgresql-dev libxml2-dev
RUN docker-php-ext-install \
    bcmath \
    ctype \
    fileinfo \
    mbstring \
    pdo_mysql \
    pdo_pgsql \
    xml
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer
WORKDIR /app
COPY . .
COPY ./storage/build/nginx.conf /etc/nginx/nginx.conf
RUN composer install --no-interaction --optimize-autoloader --no-dev
# Optimizing Configuration loading
RUN php artisan config:cache
# Optimizing Route loading
RUN php artisan route:cache
# Optimizing View loading
RUN php artisan view:cache
CMD ["nginx", "-g", "daemon off;"]
