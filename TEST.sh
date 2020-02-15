# Utility script to test services

# Pricing
docker exec -it petssenger_pricing-api_1 go test

# Ride
docker exec -it petssenger_ride-api_1 yarn test
