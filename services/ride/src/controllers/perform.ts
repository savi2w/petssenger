import { Context, Next } from "koa";

import { getLastEstimate, deleteLastEstimate } from "../models/estimate";
import { performRide } from "../models/perform";

const perform = async (ctx: Context, next: Next): Promise<void> => {
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

  let estimate;
  try {
    estimate = await getLastEstimate(uuid);
  } catch (err) {
    ctx.status = 500;
    ctx.body = {
      message: err.message,
      payload: null
    };

    return next();
  }

  try {
    await performRide(uuid, estimate);
  } catch (err) {
    ctx.status = 500;
    ctx.body = {
      message: err.message,
      payload: null
    };
  }

  try {
    await deleteLastEstimate(uuid);
  } catch (err) {
    ctx.status = 500;
    ctx.body = {
      message: err.message,
      payload: null
    };
  }

  ctx.status = 201;
  ctx.body = {
    message: null,
    payload: estimate
  };

  return next();
};

export default perform;
