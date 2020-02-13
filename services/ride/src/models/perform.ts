import { Entity, PrimaryGeneratedColumn, Column, getRepository } from "typeorm";

import { Estimate } from "./estimate";
import { increaseDynamicFees } from "./pricing";
import { GetFeesByCity } from "./pricing/interfaces";

@Entity()
export class Perform {
  @PrimaryGeneratedColumn("uuid")
  id!: string;

  @Column("varchar")
  user!: string;

  @Column("jsonb")
  estimate!: Estimate;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const performRide = async (
  uuid: string,
  estimate: Estimate
): Promise<void> => {
  const req = new GetFeesByCity();
  req.setCity(estimate.ride.city);

  await increaseDynamicFees(req);
  await getRepository(Perform).save({ uuid, estimate });
};
