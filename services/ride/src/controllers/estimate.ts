import { Context, Next } from "koa";
import * as yup from "yup";

import { authUser } from "../models/user";
import { getEstimate } from "../models/estimate";

export const rideSchema = yup.object({
  city: yup.string().required(),
  distance: yup
    .number()
    .positive()
    .required(),
  time: yup
    .number()
    .positive()
    .required()
});

const estimate = async (ctx: Context, next: Next): Promise<void> => {
  const uuid = ctx.request.get("X-User-ID");

  try {
    await authUser(uuid);
  } catch (err) {
    ctx.status = 401;
    ctx.body = {
      message: err.message,
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

  ctx.status = 201;
  ctx.body = {
    message: "Estimated",
    payload: estimate
  };

  return next();
};

export default estimate;
