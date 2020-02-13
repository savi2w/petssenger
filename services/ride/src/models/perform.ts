import { Entity, PrimaryGeneratedColumn, Column, getRepository } from "typeorm";

import { Estimate } from "./estimate";

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
export const performRide = (uuid: string, estimate: Estimate): Promise<any> =>
  getRepository(Perform).save({ uuid, estimate });
