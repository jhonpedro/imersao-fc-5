import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { CreateAccountDto } from './dto/create-account.dto';
import { UpdateAccountDto } from './dto/update-account.dto';
import { Account } from './entities/account.entity';
import { EmptyResultError } from 'sequelize';

@Injectable()
export class AccountsService {
  constructor(@InjectModel(Account) private accountModel: typeof Account) {}

  create(createAccountDto: CreateAccountDto) {
    return this.accountModel.create(createAccountDto);
  }

  async findAll() {
    return this.accountModel.findAll();
  }

  async findBy({ column, value }: { column: 'id' | 'token'; value: string }) {
    return this.accountModel.findOne({
      where: {
        [column]: value,
      },
      rejectOnEmpty: new EmptyResultError(
        `Account not found in <${column}> with value <${value}>`,
      ),
    });
  }

  async update(id: string, updateAccountDto: UpdateAccountDto) {
    const updates = await this.accountModel.update(updateAccountDto, {
      where: {
        id,
      },
    });

    return !!updates.length;
  }

  async remove(id: string) {
    return this.accountModel.destroy({ where: { id } });
  }
}
