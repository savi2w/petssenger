import path from "path";
import { ConnectionOptions } from "typeorm";

import { Perform } from "../models/perform";

const migrations = [path.join("dist", "migrations", "*.js")];
const ormConfig: ConnectionOptions = {
  type: "postgres",
  host: "ride-postgres",
  database: "ride",
  username: "postgres",
  password: "122ff0fb63174b0f8496ec3f30c64470",
  entities: [Perform],
  migrations
};

export = ormConfig;
