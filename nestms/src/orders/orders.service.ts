import { Inject, Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';
import { AccountStorageService } from '../accounts/account-storage/account-storage.service';
import { EmptyResultError } from 'sequelize';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { UpdateOptions } from 'sequelize';

@Injectable()
export class OrdersService {
  constructor(
    @InjectModel(Order) private orderModel: typeof Order,
    private accountStorageService: AccountStorageService,
    @Inject('KAFKA_PRODUCER') private kafkaProducer: Producer,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    const order = await this.orderModel.create({
      ...createOrderDto,
      account_id: this.accountStorageService.account.id,
    });

    await this.kafkaProducer.send({
      topic: 'transactions',
      messages: [{ key: 'transaction', value: JSON.stringify(order) }],
    });

    return order;
  }

  async findAll() {
    return this.orderModel.findAll({
      where: { account_id: this.accountStorageService.account.id },
    });
  }

  async findOne(id: string) {
    return this.orderModel.findOne({
      where: {
        id,
        account_id: this.accountStorageService.account.id,
      },
      rejectOnEmpty: new EmptyResultError(
        `Account not found with id <${id}> for this account`,
      ),
    });
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const updateOptions: UpdateOptions = {
      where: {
        id,
      },
    };

    if (this.accountStorageService?.account?.id) {
      updateOptions.where['account_id'] = this.accountStorageService.account.id;
    }

    const updates = await this.orderModel.update(updateOrderDto, updateOptions);

    return !!updates.length;
  }

  async remove(id: string) {
    console.log(id, this.accountStorageService.account.id);
    return this.orderModel.destroy({
      where: { id, account_id: this.accountStorageService.account.id },
    });
  }
}
