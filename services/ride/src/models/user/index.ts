import grpc from "grpc";
import bluebird from "bluebird";

import env from "../../config/env";

import { UserClient, UserClientAsync, AuthUserRequest } from "./interfaces";

const cli = bluebird.promisifyAll(
  new UserClient(env.userAddr, grpc.credentials.createInsecure())
) as UserClientAsync;

export const authUser = async (uuid: string): Promise<boolean> => {
  const req = new AuthUserRequest();
  req.setUser(uuid);

  const res = await cli.authUserAsync(req);
  return res.getAuthed();
};
