import {
  AuthUserRequest,
  AuthUserResponse
} from "../../../../../protos/user_pb";

import { UserClient } from "../../../../../protos/user_grpc_pb";

export * from "../../../../../protos/user_pb";
export * from "../../../../../protos/user_grpc_pb";

// https://github.com/Microsoft/TypeScript/issues/8685#issuecomment-240201897
export interface UserClientAsync extends UserClient {
  authUserAsync(req: AuthUserRequest): Promise<AuthUserResponse>;
}
