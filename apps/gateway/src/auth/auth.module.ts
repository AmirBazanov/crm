import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { LoggerModule } from '../../../../libs/logger/logger.module';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'AUTH_PACKAGE',
        transport: Transport.GRPC,
        options: {
          package: 'auth.v2',
          protoPath: join(__dirname, '..', '..', 'proto', 'auth', 'v2', 'auth.proto'),
          url: '0.0.0.0:5001',
        },
      },
    ]),
    LoggerModule.forService('gateway/auth'),
  ],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule { }
