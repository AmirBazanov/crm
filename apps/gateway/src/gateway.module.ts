import { Module } from '@nestjs/common';
import { AuthModule } from './auth/auth.module';
import { APP_FILTER } from '@nestjs/core';
import { GrpcExceptionFilter } from '../../../libs/exeption-filters/rpc-exceptions.filter';
import { UsersModule } from './users/users.module';

@Module({
  imports: [AuthModule, UsersModule],
  providers: [{ provide: APP_FILTER, useClass: GrpcExceptionFilter }],
})
export class GatewayModule { }
