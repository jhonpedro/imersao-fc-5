import {
  Column,
  DataType,
  Model,
  PrimaryKey,
  Table,
} from 'sequelize-typescript';
import { randomBytes } from 'crypto';

@Table({
  tableName: 'accounts',
  createdAt: 'created_at',
  updatedAt: 'updated_at',
})
export class Account extends Model {
  @PrimaryKey
  @Column({
    type: DataType.UUID,
    unique: true,
    defaultValue: DataType.UUIDV4,
    allowNull: false,
  })
  id: string;

  @Column({ allowNull: false })
  name: string;

  @Column({
    allowNull: false,
    defaultValue: () =>
      randomBytes(40)
        .toString('base64')
        .replace(/[^(1-9a-zA-Z)]/gm, ''),
  })
  token: string;
}
