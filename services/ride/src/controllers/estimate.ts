import { Context, Next } from "koa";

import { getEstimate } from "../models/estimate";
import { rideSchema } from "../models/ride";

const estimate = async (ctx: Context, next: Next): Promise<void> => {
  const uuid = ctx.request.get("X-User-ID");

  // Replace with an real check when the user's microservice is ready
  if (uuid !== "c68914e0-d085-4049-81eb-789322ce284c") {
    ctx.status = 401;
    ctx.body = {
      message: "Unauthorized",
      payload: null
    };

    return next();
  }

  let ride;
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

  let estimate;
  try {
    estimate = await getEstimate(uuid, ride);
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
    payload: estimate
  };

  return next();
};

export default estimate;
