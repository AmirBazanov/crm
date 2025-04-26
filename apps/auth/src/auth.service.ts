import { Injectable} from '@nestjs/common';
import {PrismaService} from "../prisma/prisma.service";
import * as bcrypt from 'bcrypt';
import {AuthRegisterDto} from "../../../libs/dto/auth.dto";
import {mapPrismaErrorToRpcError} from "../../../libs/exeptions-mapper/prisma-to-grpc.mapper";
import {ConfigService} from "@nestjs/config";
import {WinstonLoggerService} from "../../../libs/logger/logger.service";

@Injectable()
export class AuthService {
  constructor(private prisma: PrismaService, private readonly logger: WinstonLoggerService,private readonly configService: ConfigService) {
  }

  async register(dto: AuthRegisterDto) {
    try {
      return await this.prisma.userCredential.create({data: {email: dto.email, password: await this.hashPassword(dto.password)}})
    }catch (err){
      this.logger.error(err.message,err.trace)
      throw mapPrismaErrorToRpcError(err)
    }
  }

  private async hashPassword(password: string): Promise<string> {
    let salt_rounds = this.configService.get<number>("SALT_ROUNDS")
    if (!salt_rounds){
      salt_rounds = 10
    }
    return bcrypt.hash(password,salt_rounds);
  }

  // Проверка пароля при логине
  async comparePasswords(plainPassword: string, hashedPassword: string): Promise<boolean> {
    return bcrypt.compare(plainPassword, hashedPassword); // Сравнение паролей
  }

}
