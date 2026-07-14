import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { AuthLoginDto, AuthRegisterDto } from '../../../../libs/dto/auth.dto';
import { ClientGrpc } from '@nestjs/microservices';
import { AuthServiceClient } from '../../../../proto/gen/ts/auth/v1/auth';
import { lastValueFrom } from 'rxjs';
import { mapGrpcErrorToHttp } from '../../../../libs/exeptions-mapper/grpc-to-http.mapper';
import { WinstonLoggerService } from '../../../../libs/logger/logger.service';

@Injectable()
export class AuthService implements OnModuleInit {
  private authService: AuthServiceClient;
  constructor(
    @Inject('AUTH_PACKAGE') private client: ClientGrpc,
    private readonly logger: WinstonLoggerService,
  ) { }
  onModuleInit() {
    this.authService = this.client.getService<AuthServiceClient>('AuthService');
  }
  async register(dto: AuthRegisterDto) {
    const op = "auth.service.register "
    try {
      return await lastValueFrom(this.authService.register(dto));
    } catch (err) {
      this.logger.error(op + err.message, err.stack);
      throw mapGrpcErrorToHttp(err);
    }
  }

  async login(dto: AuthLoginDto) {
    try {
      return await lastValueFrom(this.authService.login(dto));
    } catch (err) {
      this.logger.error(err.message, err.stack);
      throw mapGrpcErrorToHttp(err);
    }
  }
}
