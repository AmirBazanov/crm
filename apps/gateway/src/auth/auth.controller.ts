import { Body, Controller, HttpException, Post, UseFilters, ValidationPipe } from '@nestjs/common';
import {
  AuthDtoResp,
  AuthLoginDto,
  AuthLogoutDto,
  AuthRefreshDto,
  AuthRefreshDtoResp,
  AuthRegisterDto,
} from '../../../../libs/dto/auth.dto';
import { AuthService } from './auth.service';
import { HttpExceptionFilter } from '../../../../libs/exeption-filters/http-exeptions.filter';
import { WinstonLoggerService } from '../../../../libs/logger/logger.service';

@Controller('auth')
@UseFilters(HttpExceptionFilter)
export class AuthController {
  constructor(private readonly authService: AuthService, private readonly logger: WinstonLoggerService) { }
  @Post('login')
  async login(@Body() body: AuthLoginDto) {
    return await this.authService.login(body);
  }

  @Post('register')
  async register(@Body() body: AuthRegisterDto) {
    return await this.authService.register(body);
  }

  @Post('refresh')
  async refresh(@Body() body: AuthRefreshDto) {
    return AuthRefreshDtoResp;
  }

  @Post('logout')
  async logout(@Body() body: AuthLogoutDto) {
    return AuthRefreshDtoResp;
  }

}
