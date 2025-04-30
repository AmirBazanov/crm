import { HttpException, HttpStatus } from '@nestjs/common';
import { status as GrpcStatus } from '@grpc/grpc-js';
import { RpcException } from '@nestjs/microservices';

export function mapGrpcErrorToHttp(error: RpcException): HttpException {
  if (!error || typeof error !== 'object') {
    return new HttpException(
      'Unknown error occurred',
      HttpStatus.INTERNAL_SERVER_ERROR,
    );
  }

  const message: string = error.message || 'gRPC Error';
  switch (error.code) {
    case GrpcStatus.NOT_FOUND:
      return new HttpException(message, HttpStatus.NOT_FOUND);
    case GrpcStatus.INVALID_ARGUMENT:
      return new HttpException(message, HttpStatus.BAD_REQUEST);
    case GrpcStatus.UNAUTHENTICATED:
      return new HttpException(message, HttpStatus.UNAUTHORIZED);
    case GrpcStatus.PERMISSION_DENIED:
      return new HttpException(message, HttpStatus.FORBIDDEN);
    case GrpcStatus.ALREADY_EXISTS:
      return new HttpException(message, HttpStatus.CONFLICT);
    case GrpcStatus.FAILED_PRECONDITION:
      return new HttpException(message, HttpStatus.BAD_REQUEST);
    case GrpcStatus.UNAVAILABLE:
      return new HttpException(message, HttpStatus.SERVICE_UNAVAILABLE);
    default:
      return new HttpException(message, HttpStatus.INTERNAL_SERVER_ERROR);
  }
}
