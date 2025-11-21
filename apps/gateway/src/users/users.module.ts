import { Module } from '@nestjs/common';
import { UsersService } from './users.service';
import { UsersController } from './users.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { LoggerModule } from 'libs/logger/logger.module';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'USER_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'users.v3',
          protoPath: join(__dirname, '..', '..', 'proto', 'users', 'v3', 'users.proto'),
          url: '0.0.0.0:4000',
        },
      },
    ]),
    LoggerModule.forService('gateway/user'),
  ],
  controllers: [UsersController],
  providers: [UsersService],
})
export class UsersModule { }
