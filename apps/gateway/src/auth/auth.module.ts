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
          package: 'auth',
          protoPath: join(__dirname, '..', '..', 'proto', 'auth', 'v1','auth.proto'),
          url: '0.0.0.0:5001',
        },
      },
    ]),
    LoggerModule.forService('gateway'),
  ],
  controllers: [AuthController],
  providers: [AuthService, { provide: 'APP_NAME', useValue: 'gateway' }],
})
export class AuthModule {}
