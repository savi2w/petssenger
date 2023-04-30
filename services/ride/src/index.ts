import "reflect-metadata";
import { createConnection } from "typeorm";

import env from "./config/env";
import ormConfig from "./config/ormconfig";
import { http } from "./http";

createConnection(ormConfig).then(() => {
  http().listen(env.port);
});
