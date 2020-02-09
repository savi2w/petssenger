import grpc from "grpc";
import bluebird from "bluebird";

import env from "../config/env";
import { Ride } from "../controllers/estimate";
import round from "../utils/round";

import {
  PricingClient,
  PricingClientAsync,
  GetFeesByCity
} from "./pricing.async";

const cli = bluebird.promisifyAll(
  new PricingClient(env.pricing, grpc.credentials.createInsecure())
) as PricingClientAsync;

const getDynamicFees = async (req: GetFeesByCity): Promise<number> => {
  const res = await cli.getDynamicFeesByCityAsync(req);
  return res.getDynamic();
};

export const getEstimatePricing = async (ride: Ride): Promise<number> => {
  const req = new GetFeesByCity();
  req.setCity(ride.city);

  const res = await cli.getPricingFeesByCityAsync(req);
  const pricing = res.toObject();

  const dynamic = await getDynamicFees(req);

  // Test purpose
  await cli.increaseDynamicFeesByCityAsync(req);

  const estimate =
    pricing.base +
    (pricing.minute * ride.time + pricing.distance * ride.distance * dynamic) +
    pricing.service;

  return round(estimate);
};
