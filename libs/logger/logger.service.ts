import { Injectable } from '@nestjs/common';
import * as winston from 'winston';

@Injectable()
export class WinstonLoggerService {
  constructor(
    private readonly logger: winston.Logger,
    private readonly context: string,
  ) {}

  log(message: string) {
    this.logger.info(`[${this.context}] ${message}`);
  }

  error(message: string, trace?: string) {
    this.logger.error(`[${this.context}] ${message} ${trace || ''}`);
  }

  warn(message: string) {
    this.logger.warn(`[${this.context}] ${message}`);
  }

  debug(message: string) {
    this.logger.debug(`[${this.context}] ${message}`);
  }
}
