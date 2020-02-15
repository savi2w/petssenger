import {
  AuthUserRequest,
  AuthUserResponse
} from "../../../../../protos/user_pb";

import { UserClient } from "../../../../../protos/user_grpc_pb";

export * from "../../../../../protos/user_pb";
export * from "../../../../../protos/user_grpc_pb";

export interface UserClientAsync extends UserClient {
  authUserAsync(req: AuthUserRequest): Promise<AuthUserResponse>;
}
