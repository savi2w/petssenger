import grpc from "grpc";
import bluebird from "bluebird";

import env from "../../config/env";

import {
  PricingClient,
  PricingClientAsync,
  GetFeesByCity,
  GetPricingFeesByCityResponse,
  Empty
} from "./interfaces";

const cli = bluebird.promisifyAll(
  new PricingClient(env.pricingAddr, grpc.credentials.createInsecure())
) as PricingClientAsync;

export const getPricingFees = async (
  req: GetFeesByCity
): Promise<GetPricingFeesByCityResponse.AsObject> => {
  const res = await cli.getPricingFeesByCityAsync(req);
  return res.toObject();
};

export const getDynamicFees = async (req: GetFeesByCity): Promise<number> => {
  const res = await cli.getDynamicFeesByCityAsync(req);
  return res.getDynamic();
};

export const increaseDynamicFees = (req: GetFeesByCity): Promise<Empty> =>
  cli.increaseDynamicFeesByCityAsync(req);
