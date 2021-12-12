import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { AccountStorageService } from './account-storage/account-storage.service';

@Injectable()
export class TokenGuard implements CanActivate {
  constructor(private accountStorage: AccountStorageService) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    if (context.getType() !== 'http') {
      return true;
    }

    const req = context.switchToHttp().getRequest();

    const token = req.headers?.['x-token'] as string | undefined;

    if (!token) {
      return false;
    }

    try {
      await this.accountStorage.setBy(token);

      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }
}
