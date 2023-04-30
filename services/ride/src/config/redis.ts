import Redis from "ioredis";

const redis = new Redis({
  host: "ride-redis"
});

export default redis;
