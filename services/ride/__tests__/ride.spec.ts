import request from "supertest";

import env from "../src/config/env";
import { Ride } from "../src/models/estimate";

interface Context {
  app: string;
  ride: Ride;
  user: string;
}

const ctx: Context = {
  app: "http://localhost:" + env.port,
  ride: {
    city: "SAO_PAULO",
    distance: 4,
    time: 10
  },
  user: "08842beb-a4fc-4cb2-9f87-d80f1a2d5045"
};

it("POST /ride/estimate", done => {
  request(ctx.app)
    .post("/ride/estimate")
    .set("X-User-ID", ctx.user)
    .send(ctx.ride)
    .expect("Content-Type", /json/)
    .expect(201)
    .end(err => {
      if (err) throw err;
      done();
    });
});

it("POST /ride/estimate with unauthorized user", done => {
  const random = "ff1b05a6-82fa-492b-bf9d-a70f7277ddfa";
  request(ctx.app)
    .post("/ride/estimate")
    .set("X-User-ID", random)
    .send(ctx.ride)
    .expect("Content-Type", /json/)
    .expect(401)
    .end(err => {
      if (err) throw err;
      done();
    });
});

it("POST /ride/estimate with an wrong body", done => {
  const random = {};
  request(ctx.app)
    .post("/ride/estimate")
    .set("X-User-ID", ctx.user)
    .send(random)
    .expect("Content-Type", /json/)
    .expect(400)
    .end(err => {
      if (err) throw err;
      done();
    });
});

it("POST /ride/perform", done => {
  request(ctx.app)
    .post("/ride/perform")
    .set("X-User-ID", ctx.user)
    .expect("Content-Type", /json/)
    .expect(201)
    .end(err => {
      if (err) throw err;
      done();
    });
});

it("POST /ride/perform with unauthorized user", done => {
  const random = "ff1b05a6-82fa-492b-bf9d-a70f7277ddfa";
  request(ctx.app)
    .post("/ride/perform")
    .set("X-User-ID", random)
    .send(ctx.ride)
    .expect("Content-Type", /json/)
    .expect(401)
    .end(err => {
      if (err) throw err;
      done();
    });
});

it("POST /ride/perform with an user without estimate", done => {
  request(ctx.app)
    .post("/ride/perform")
    .set("X-User-ID", ctx.user)
    .send(ctx.ride)
    .expect("Content-Type", /json/)
    .expect(500)
    .end(err => {
      if (err) throw err;
      done();
    });
});
