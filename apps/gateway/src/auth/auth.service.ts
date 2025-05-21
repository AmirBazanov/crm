import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { AuthRegisterDto } from '../../../../libs/dto/auth.dto';
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
  ) {}
  onModuleInit() {
    this.authService = this.client.getService<AuthServiceClient>('AuthService');
  }
  async register(dto: AuthRegisterDto) {
    try {
      return await lastValueFrom(this.authService.register(dto));
    } catch (err) {
      this.logger.error(err.message, err.stack);
      throw mapGrpcErrorToHttp(err);
    }
  }
}
