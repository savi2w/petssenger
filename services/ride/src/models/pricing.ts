import grpc from "grpc";
import bluebird from "bluebird";

import env from "../config/env";
import { Ride } from "../controllers/estimate";
import round from "../utils/round";

import messages from "../../../../protos/pricing_pb";
import { PricingClient } from "../../../../protos/pricing_grpc_pb";

interface PricingClientAsync extends PricingClient {
  getPricingFeesByCityAsync(
    req: messages.GetFeesByCity
  ): Promise<messages.GetPricingFeesByCityResponse>;

  getDynamicFeesByCityAsync(
    req: messages.GetFeesByCity
  ): Promise<messages.GetDynamicFeesByCityResponse>;
}

const cli = bluebird.promisifyAll(
  new PricingClient(env.pricing, grpc.credentials.createInsecure())
) as PricingClientAsync;

export const getEstimatePricing = async (ride: Ride): Promise<number> => {
  const req = new messages.GetFeesByCity();
  req.setCity(ride.city);

  const pri = await cli.getPricingFeesByCityAsync(req);
  const pricing = pri.toObject();

  const dyn = await cli.getDynamicFeesByCityAsync(req);
  const dynamic = dyn.getDynamic();

  const estimate =
    pricing.base +
    (pricing.minute * ride.time + pricing.distance * ride.distance * dynamic) +
    pricing.service;

  return round(estimate);
};
