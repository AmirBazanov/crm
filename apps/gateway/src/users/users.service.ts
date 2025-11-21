import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { CreateUserDto } from 'libs/dto/user.dto';
import { mapGrpcErrorToHttp } from 'libs/exeptions-mapper/grpc-to-http.mapper';
import { WinstonLoggerService } from 'libs/logger/logger.service';
import { UserServiceClient } from 'proto/gen/ts/users/v3/users';

import { lastValueFrom } from 'rxjs';

@Injectable()
export class UsersService implements OnModuleInit {
  private userService: UserServiceClient
  constructor(@Inject('USER_PACKAGE') private client: ClientGrpc,
    private readonly logger: WinstonLoggerService) { }
  onModuleInit() {
    this.userService = this.client.getService<UserServiceClient>('UserService')
  }

  async create(dto: CreateUserDto) {
    try {
      return await lastValueFrom(this.userService.createUser(dto))
    } catch (error) {
      this.logger.error(error.message, error.stack)
      throw mapGrpcErrorToHttp(error)
    }
  }

  async findOne(id: number) {
    try {
      return await lastValueFrom(this.userService.getUser({ id }))
    } catch (error) {
      this.logger.error(error.message, error.stack)
      throw mapGrpcErrorToHttp(error)
    }
  }
}
