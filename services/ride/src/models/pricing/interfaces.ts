import {
  GetFeesByCity,
  GetPricingFeesByCityResponse,
  GetDynamicFeesByCityResponse,
  Empty
} from "../../protos/pricing_pb";

import { PricingClient } from "../../protos/pricing_grpc_pb";

export * from "../../protos/pricing_pb";
export * from "../../protos/pricing_grpc_pb";

export interface PricingClientAsync extends PricingClient {
  getPricingFeesByCityAsync(
    req: GetFeesByCity
  ): Promise<GetPricingFeesByCityResponse>;

  getDynamicFeesByCityAsync(
    req: GetFeesByCity
  ): Promise<GetDynamicFeesByCityResponse>;

  increaseDynamicFeesByCityAsync(req: GetFeesByCity): Promise<Empty>;
}
