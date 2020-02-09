import grpc from "grpc";
import bluebird from "bluebird";

import env from "../config/env";
import { Ride } from "../controllers/estimate";

import {
  GetPricingFeesByCityRequest,
  GetPricingFeesByCityResponse
} from "../../../../protos/pricing_pb";
import { PricingClient } from "../../../../protos/pricing_grpc_pb";

interface PricingClientAsync extends PricingClient {
  getPricingFeesByCityAsync(
    req: GetPricingFeesByCityRequest
  ): Promise<GetPricingFeesByCityResponse>;
}

const cli = bluebird.promisifyAll(
  new PricingClient(env.pricing, grpc.credentials.createInsecure())
) as PricingClientAsync;

export const getEstimatePricing = async (ride: Ride): Promise<number> => {
  const req = new GetPricingFeesByCityRequest();
  req.setCity(ride.city);

  const res = await cli.getPricingFeesByCityAsync(req);
  const pricing = res?.toObject();

  const estimate =
    pricing?.base +
    (pricing?.minute * ride.time + pricing?.distance * ride.distance * 100) +
    pricing?.service;

  if (isNaN(estimate)) {
    throw new TypeError('"estimate" must be a number');
  }

  return estimate;
};
