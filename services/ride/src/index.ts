import "reflect-metadata";
import { createConnection } from "typeorm";

import env from "./config/env";
import ormConfig from "./config/ormconfig";
import server from "./server";

createConnection(ormConfig).then(() => {
  server().listen(env.port);
});
