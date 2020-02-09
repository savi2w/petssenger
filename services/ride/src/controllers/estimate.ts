import { Context, Next } from "koa";
import * as yup from "yup";

import { getEstimatePricing } from "../models/pricing";

const rideSchema = yup.object({
  city: yup.string().required(),
  distance: yup
    .number()
    .integer()
    .required(),
  time: yup
    .number()
    .integer()
    .required()
});

export type Ride = yup.InferType<typeof rideSchema>;

const estimate = async (ctx: Context, next: Next): Promise<void> => {
  let ride: Ride;
  try {
    ride = await rideSchema.validate(ctx.request.body);
  } catch (err) {
    ctx.status = 400;
    ctx.body = {
      message: err.message,
      payload: null
    };

    return next();
  }

  let pricing;
  try {
    pricing = await getEstimatePricing(ride);
  } catch (err) {
    ctx.status = 500;
    ctx.body = {
      message: err.message,
      payload: null
    };

    return next();
  }

  ctx.body = {
    message: null,
    payload: {
      pricing
    }
  };

  return next();
};

export default estimate;
