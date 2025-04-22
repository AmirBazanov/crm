import { Module } from '@nestjs/common';
import { Prisma } from './prisma.service';
import {ConfigModule} from "@nestjs/config";

@Module({
  imports: [ConfigModule],
  providers: [Prisma]
})
export class PrismaModule {}
