import path from "path";
import { ConnectionOptions } from "typeorm";

const migrations = [path.join("dist", "migrations", "*.js")];
const ormConfig: ConnectionOptions = {
  type: "postgres",
  host: "ride-postgres",
  database: "ride",
  username: "postgres",
  password: "122ff0fb63174b0f8496ec3f30c64470",
  entities: [],
  migrations,
  synchronize: process.env.NODE_ENV !== "production"
};

export = ormConfig;
