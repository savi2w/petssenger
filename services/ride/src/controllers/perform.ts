import { Context, Next } from "koa";

import { authUser } from "../models/user";
import { getLastEstimate, deleteLastEstimate } from "../models/estimate";
import { performRide } from "../models/perform";

const perform = async (ctx: Context, next: Next): Promise<void> => {
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
    message: "Performed",
    payload: estimate
  };

  return next();
};

export default perform;
