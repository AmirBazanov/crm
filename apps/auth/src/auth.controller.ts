import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import {
  AuthRegisterRequest,
  AuthServiceController,
} from '../../../proto/gen/ts/auth/v1/auth';
import { AuthService } from './auth.service';

@Controller()
// @ts-ignore
export class AuthController implements AuthServiceController {
  constructor(private readonly authService: AuthService) {}

  @GrpcMethod('AuthService', 'Register')
  async Register(data: AuthRegisterRequest) {
    return await this.authService.register(data);
  }
}
