change api version and recreate followd this command

docker-compose create --force-recreate api && docker-compose start api && docker-compose create --force-recreate api2 && docker-compose start api2
