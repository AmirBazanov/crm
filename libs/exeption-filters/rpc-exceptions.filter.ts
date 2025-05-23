import { Catch, ArgumentsHost } from '@nestjs/common';
import { BaseRpcExceptionFilter, RpcException } from '@nestjs/microservices';

@Catch()
export class GrpcExceptionFilter extends BaseRpcExceptionFilter {
  catch(exception: unknown, host: ArgumentsHost) {
    if (exception instanceof RpcException) {
      return super.catch(exception, host);
    }

    return super.catch(new RpcException('Unexpected error'), host);
  }
}
