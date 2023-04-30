import { Entity, PrimaryGeneratedColumn, Column, getConnection } from "typeorm";

import { Estimate } from "./estimate";
import { increaseDynamicFees } from "./pricing";
import { GetFeesByCity } from "./pricing/interfaces";

@Entity()
export class Perform {
  @PrimaryGeneratedColumn("uuid")
  id!: string;

  @Column("uuid")
  user!: string;

  @Column("jsonb")
  estimate!: Estimate;

  @Column("timestamp")
  performed_at!: string;
}

export const performRide = async (user: string, estimate: Estimate) => {
  await getConnection()
    .createQueryBuilder()
    .insert()
    .into(Perform)
    .values({ user, estimate })
    .execute();

  const req = new GetFeesByCity();
  req.setCity(estimate.ride.city);

  return increaseDynamicFees(req);
};
