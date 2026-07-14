import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { ConfigModule } from '@nestjs/config';
import { PrismaModule } from '../prisma/prisma.module';
import { AuthController } from './auth.controller';
import { APP_FILTER } from '@nestjs/core';
import { GrpcExceptionFilter } from '../../../libs/exeption-filters/rpc-exceptions.filter';
import { PrismaExceptionFilter } from '../../../libs/exeption-filters/prisma-exeptions.filter';
import { LoggerModule } from '../../../libs/logger/logger.module';
import { KafkaProducerService } from './kafka.producer.service';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { UsersService } from './user.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: 'auth.env',
      isGlobal: true,
    }),
    PrismaModule,
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
    LoggerModule.forService('auth'),
  ],
  controllers: [AuthController],
  providers: [
    AuthService,
    { provide: APP_FILTER, useClass: GrpcExceptionFilter },
    { provide: APP_FILTER, useClass: PrismaExceptionFilter },
    KafkaProducerService, UsersService
  ],
})
export class AuthModule {}
