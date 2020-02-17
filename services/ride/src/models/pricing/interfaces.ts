import {
  GetFeesByCity,
  GetPricingFeesByCityResponse,
  GetDynamicFeesByCityResponse,
  Empty
} from "../../../../../protos/pricing_pb";

import { PricingClient } from "../../../../../protos/pricing_grpc_pb";

export * from "../../../../../protos/pricing_pb";
export * from "../../../../../protos/pricing_grpc_pb";

// https://github.com/Microsoft/TypeScript/issues/8685#issuecomment-240201897
export interface PricingClientAsync extends PricingClient {
  getPricingFeesByCityAsync(
    req: GetFeesByCity
  ): Promise<GetPricingFeesByCityResponse>;

  getDynamicFeesByCityAsync(
    req: GetFeesByCity
  ): Promise<GetDynamicFeesByCityResponse>;

  increaseDynamicFeesByCityAsync(req: GetFeesByCity): Promise<Empty>;
}
