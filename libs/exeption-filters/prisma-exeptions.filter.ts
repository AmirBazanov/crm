import { Catch, ArgumentsHost, ExceptionFilter } from '@nestjs/common';

import { RpcException } from '@nestjs/microservices';
import { PrismaClientKnownRequestError } from '@prisma/client/runtime/library';
import {PRISMA_ERROR_MESSAGES} from "../constants/error-messages";

@Catch(PrismaClientKnownRequestError)
export class PrismaExceptionFilter implements ExceptionFilter {
    catch(exception: PrismaClientKnownRequestError, host: ArgumentsHost) {
        const statusCode = this.getErrorCode(exception.code);
        const message = this.getErrorMessage(exception);

        return new RpcException({
            statusCode,
            message,
            errorCode: exception.code,
        });
    }

    public getErrorCode(code: string): number {
        switch (code) {
            case 'P2002':
                return 400; // Unique constraint violation
            case 'P2003':
                return 400; // Foreign key constraint violation
            case 'P2025':
                return 404; // Record not found
            case 'P2004':
                return 500; // Internal server error
            case 'P2005':
                return 500; // Invalid data provided
            // Add cases for other Prisma error codes as needed
            default:
                return 500; // Default server error for unhandled errors
        }
    }

    public getErrorMessage(exception: PrismaClientKnownRequestError): string {
        if (!exception.meta){
            exception.meta = {"target": "Unknown"}
        }
        switch (exception.code) {
            case 'P2002':
                return PRISMA_ERROR_MESSAGES.UNIQUE_CONSTRAINT_VIOLATION;
            case 'P2003':
                return  PRISMA_ERROR_MESSAGES.FOREIGN_KEY_CONSTRAINT_VIOLATION;
            case 'P2025':
                return  PRISMA_ERROR_MESSAGES.RECORD_NOT_FOUND;
            case 'P2004':
                return PRISMA_ERROR_MESSAGES.INTERNAL_SERVER_ERROR;
            case 'P2005':
                return PRISMA_ERROR_MESSAGES.INVALID_DATA_PROVIDED;
            default:
                return `${PRISMA_ERROR_MESSAGES.UNEXPECTED_ERROR}${exception.message}`;
        }
    }
}
