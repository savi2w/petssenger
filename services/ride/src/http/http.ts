import Koa from "koa";
import Router from "koa-router";
import bodyParser from "koa-bodyparser";

import estimate from "../controllers/estimate";
import perform from "../controllers/perform";

const http = (): Koa => {
  const app = new Koa();

  app.use(bodyParser());

  const router = new Router();

  router.post("/ride/estimate", estimate);
  router.post("/ride/perform", perform);

  app.use(router.routes());
  app.use(router.allowedMethods());

  return app;
};

export default http;
