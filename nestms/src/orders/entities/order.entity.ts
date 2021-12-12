import { Account } from '../../accounts/entities/account.entity';
import {
  BelongsTo,
  Column,
  DataType,
  ForeignKey,
  Model,
  PrimaryKey,
  Table,
} from 'sequelize-typescript';

export enum OrderStatusEnum {
  Pending = 'pending',
  Approved = 'approved',
  Rejected = 'rejected',
}

@Table({
  tableName: 'orders',
  createdAt: 'created_at',
  updatedAt: 'updated_at',
})
export class Order extends Model {
  @PrimaryKey
  @Column({
    type: DataType.UUID,
    unique: true,
    defaultValue: DataType.UUIDV4,
    allowNull: false,
  })
  id: string;

  @Column({ allowNull: false, type: DataType.DECIMAL(10, 2) })
  amount: number;

  @Column({ allowNull: false })
  credit_card_number: string;

  @Column({ allowNull: false })
  credit_card_name: string;

  @Column({ allowNull: false })
  credit_card_expiration_month: number;

  @Column({ allowNull: false })
  credit_card_expiration_year: number;

  @Column({ allowNull: false, defaultValue: OrderStatusEnum.Pending })
  status: OrderStatusEnum;

  @ForeignKey(() => Account)
  @Column({
    type: DataType.UUID,
    defaultValue: DataType.UUIDV4,
  })
  account_id: string;

  @BelongsTo(() => Account)
  account: Account;
}
