docker exec -it petssenger_pricing-api_1 go run ./migrations/. init
# docker exec -it petssenger_pricing-api_1 go run ./migrations/. down
docker exec -it petssenger_pricing-api_1 go run ./migrations/. up

# docker exec -it petssenger_ride-api_1 yarn migrations:revert
docker exec -it petssenger_ride-api_1 yarn migrations:run

docker exec -it petssenger_user-api_1 go run ./migrations/. init
# docker exec -it petssenger_user-api_1 go run ./migrations/. down
docker exec -it petssenger_user-api_1 go run ./migrations/. up
