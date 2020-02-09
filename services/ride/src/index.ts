import env from "./config/env";
import server from "./server";

Promise.resolve().then(() => {
  server().listen(env.port);
});
