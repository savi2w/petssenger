import { Ride } from "../controllers/estimate";

import grpc from "grpc";
import { GetPricingFeesByCityRequest } from "../../../../protos/pricing_pb";
import { PricingClient } from "../../../../protos/pricing_grpc_pb";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const getEstimatePricing = (ride: Ride): Promise<any> =>
  new Promise((resolve, reject) => {
    const cli = new PricingClient(
      "pricing-api:50051",
      grpc.credentials.createInsecure()
    );

    const req = new GetPricingFeesByCityRequest();
    req.setCity(ride.city);

    cli.getPricingFeesByCity(req, (err, res) => {
      if (err) {
        reject(err);
        return undefined;
      }

      resolve(res?.toObject());
    });
  });
