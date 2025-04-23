import {Module} from '@nestjs/common';
import {AuthService} from './auth.service';
import {ConfigModule} from "@nestjs/config";
import {PrismaModule} from "../prisma/prisma.module";
import {AuthController} from './auth.controller';

@Module({
  imports: [ConfigModule.forRoot({
    envFilePath: 'auth.env',
    isGlobal: true
  }), PrismaModule,],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
