import { BadRequestException, HttpException, Injectable } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import * as bcrypt from 'bcrypt';
import { AuthDtoResp, AuthLoginDto, AuthRegisterDto } from '../../../libs/dto/auth.dto';
import { mapPrismaErrorToRpcError } from '../../../libs/exeptions-mapper/prisma-to-grpc.mapper';
import { ConfigService } from '@nestjs/config';
import { WinstonLoggerService } from '../../../libs/logger/logger.service';
import { KafkaProducerService } from './kafka.producer.service';
import { async } from 'rxjs';
import { status as GrpcStatus } from '@grpc/grpc-js';
import { UsersService } from './user.service';
import { userService } from '../../web/src/api/userService';
import { error } from 'winston';
import { PrismaClientKnownRequestError } from '@prisma/client/runtime/library';
import { RpcException } from '@nestjs/microservices';
import { AuthRegisterRequest } from '../../../proto/gen/ts/auth/v2/auth';


@Injectable()
export class AuthService {
  constructor(
    private prisma: PrismaService,
    private readonly logger: WinstonLoggerService,
    private readonly configService: ConfigService,
    private readonly kafkaService: KafkaProducerService,
    private readonly userService: UsersService
  ) { }

  async register(dto: AuthRegisterRequest) {
    const op = "auth.service.register ";
    const nicknameTaken = await this.isNicknameTaken(dto.nickname);
    const password = await this.hashPassword(dto.password)
    dto.password = password
    if (nicknameTaken) {
      this.logger.error(op+"nickname taken")
      throw new RpcException({
        code: GrpcStatus.ALREADY_EXISTS,
        message: "nickname already exists",
      });
    }

    let userCredential: { id: any; email?: string; password?: string; createdAt?: Date; updatedAt?: Date; };
    try {
      userCredential = await this.prisma.userCredential.create({
        data: {
          email: dto.email,
          password: password,
        },
      });
    } catch (error) {
      if (error instanceof PrismaClientKnownRequestError && error.code === "P2002") {
        const field = error.meta?.target?.[0] ?? "field";
        this.logger.error(op + ": duplicate " + field);

        throw new RpcException({
          code: GrpcStatus.ALREADY_EXISTS,
          message: `${field} already exists`,
        });
      }
      throw error;
    }

    await this.kafkaService.send("user-crud", userCredential.id, "create-user", dto);

    return userCredential;
  }

  // TODO: test throwing error (should be http type)
  async login(dto: AuthLoginDto) {
    const user = await this.prisma.userCredential.findUnique({
      where: {
        email: dto.email,
      },
    });
    if (!user) {
      throw new Error('User not found');
    }
    const isPasswordValid = await this.comparePasswords(
      dto.password,
      user.password,
    );
    if (!isPasswordValid) {
      throw new Error('Invalid password');
    }
    return user;
  }

  private async hashPassword(password: string): Promise<string> {
    let salt_rounds = this.configService.get<number>('SALT_ROUNDS');
    if (!salt_rounds) {
      salt_rounds = 10;
    }
    return bcrypt.hash(password, salt_rounds);
  }

  async comparePasswords(
    plainPassword: string,
    hashedPassword: string,
  ): Promise<boolean> {
    return bcrypt.compare(plainPassword, hashedPassword);
  }
  async isNicknameTaken(nickname: string): Promise<boolean> {
    try {
      const result = await this.userService.findByNickname(nickname);
      return !!result.user;
    } catch (err) {
      if (err.code === GrpcStatus.NOT_FOUND) {
        return false;
      }
      throw err;
    }
  }
}


