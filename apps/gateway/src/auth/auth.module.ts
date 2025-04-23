import {Module} from '@nestjs/common';
import {AuthController} from './auth.controller';
import {AuthService} from './auth.service';
import {ClientsModule, Transport} from "@nestjs/microservices";
import {join} from "path";

@Module({
  imports: [ClientsModule.register([{
    name: "AUTH_PACKAGE",
    transport: Transport.GRPC,
    options:{
      package: "auth",
      protoPath: join(__dirname, '..', '..', '..', 'proto', 'auth.proto'),
      url: "0.0.0.0:5001"
    }
  }])],
  controllers: [AuthController],
  providers: [AuthService]
})
export class AuthModule {}
