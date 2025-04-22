import { NestFactory } from '@nestjs/core';
import { GatewayModule } from './gateway.module';
import {ValidationPipe} from "@nestjs/common";

async function bootstrap() {
  const app = await NestFactory.create(GatewayModule);
  await app.listen(process.env.port ?? 3000);
  app.useGlobalPipes(new ValidationPipe())
}
bootstrap();
