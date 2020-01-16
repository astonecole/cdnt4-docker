# Docker container run

## Installation de MySQL

```sh
docker run -d --name mysqldb \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_DATABASE=mydb \
    --network=aston_network \
    -p 3306:3306 \
    mysql --default-authentication-plugin=mysql_native_password 
```

## Installation de NextCloud

```sh
docker run -d --name mycloud \
    --network=aston_network \
    -p 8080:80 \
    nextcloud
```

## Installation de Wordpress

```sh
docker run --name website -p 8080:80 -d --network=aston_network wordpress
```

## Installation de PHPMyAdmin

```sh
docker run --name myadmin --network=aston_network -d -e PMA_HOST=mysqldb -p "192.168.99.100:8000":80 phpmyadmin/phpmyadmin
```

```sh
docker run -d --name php_server \
    --network=aston_network \
    -v $(pwd)/www:/var/www/html \
    nanoninja/php-fpm
```

```sh
docker run -d --name nginx_server \
    -v $(pwd)/default.conf:/etc/nginx/conf.d/default.conf \
    -v $(pwd)/www:/var/www/html \
    -p 9090:80 \
    --network=aston_network \
    nginx
```

docker run -p 6060 nginx sh
docker cp 31ac9f45:/etc/nginx/conf.d/default.conf ./html
