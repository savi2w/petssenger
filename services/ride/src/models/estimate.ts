import { InferType } from "yup";

import { GetFeesByCity } from "./pricing/interfaces";
import { getPricingFees, getDynamicFees } from "./pricing";

import env from "../config/env";
import redis from "../config/redis";
import { rideSchema } from "../controllers/estimate";
import round from "../utils/round";

export type Ride = InferType<typeof rideSchema>;

export interface Estimate {
  pricing: number;
  ride: Ride;
}

const setLastEstimate = async (
  uuid: string,
  estimate: Estimate
): Promise<Estimate> => {
  const stringify = JSON.stringify(estimate);
  await redis.set(uuid, stringify, "EX", env.estimateExpirationTime);
  return estimate;
};

export const getEstimate = async (
  uuid: string,
  ride: Ride
): Promise<Estimate> => {
  const req = new GetFeesByCity();
  req.setCity(ride.city);

  const fees = await getPricingFees(req);
  const dynamic = await getDynamicFees(req);

  const pricing = round(
    fees.base +
      (fees.minute * ride.time + fees.distance * ride.distance * dynamic) +
      fees.service
  );

  const estimate = {
    pricing,
    ride
  };

  return setLastEstimate(uuid, estimate);
};

export const getLastEstimate = async (uuid: string): Promise<Estimate> => {
  const stringify = await redis.get(uuid);
  if (typeof stringify !== "string") {
    throw new Error("NoEstimate");
  }

  const estimate: Estimate = JSON.parse(stringify);

  return estimate;
};

export const deleteLastEstimate = (uuid: string): Promise<number> =>
  redis.del(uuid);
