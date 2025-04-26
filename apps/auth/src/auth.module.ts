import {Module} from '@nestjs/common';
import {AuthService} from './auth.service';
import {ConfigModule} from "@nestjs/config";
import {PrismaModule} from "../prisma/prisma.module";
import {AuthController} from './auth.controller';
import {APP_FILTER} from "@nestjs/core";
import {GrpcExceptionFilter} from "../../../libs/exeption-filters/rpc-exceptions.filter";
import {PrismaExceptionFilter} from "../../../libs/exeption-filters/prisma-exeptions.filter";
import {LoggerModule} from "../../../libs/logger/logger.module";

@Module({
  imports: [ConfigModule.forRoot({
    envFilePath: 'auth.env',
    isGlobal: true
  }), PrismaModule, LoggerModule.forService("auth")],
  controllers: [AuthController],
  providers: [AuthService, {provide: APP_FILTER, useClass: GrpcExceptionFilter},{ provide: APP_FILTER, useClass: PrismaExceptionFilter }],
})
export class AuthModule {}
