import { MigrationInterface, QueryRunner, Table } from "typeorm";

export class Bootstrap1581618123935 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query('CREATE EXTENSION IF NOT EXISTS "uuid-ossp";');
    await queryRunner.createTable(
      new Table({
        name: "perform",
        columns: [
          {
            name: "id",
            type: "uuid",
            isGenerated: true,
            generationStrategy: "uuid"
          },
          {
            name: "user",
            type: "varchar"
          },
          {
            name: "estimate",
            type: "jsonb"
          }
        ]
      })
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropTable("perform");
  }
}
