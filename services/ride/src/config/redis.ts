import Redis from "ioredis";

import env from "./env";

const redis = new Redis({
  host: env.redisAddr
});

export default redis;
