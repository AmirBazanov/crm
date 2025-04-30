import { Prisma } from '@prisma/client';
import { RpcException } from '@nestjs/microservices';
import { PrismaExceptionFilter } from '../exeption-filters/prisma-exeptions.filter';

export function mapPrismaErrorToRpcError(
  error: Prisma.PrismaClientKnownRequestError,
): RpcException {
  const code = new PrismaExceptionFilter().getErrorCode(error.code);

  const message = new PrismaExceptionFilter().getErrorMessage(error);
  return new RpcException({
    message: message,
    code: code,
  });
}
