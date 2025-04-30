import { Module } from '@nestjs/common';
import { PrismaService } from './prisma.service';
import { ConfigModule } from '@nestjs/config';
import { LoggerModule } from '../../../libs/logger/logger.module';

@Module({
  imports: [ConfigModule, LoggerModule.forService('auth')],
  providers: [PrismaService],
  exports: [PrismaService],
})
export class PrismaModule {}
