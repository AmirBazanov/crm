import { Module, DynamicModule } from '@nestjs/common';
import { LOGGER_CONTEXT } from './logger.constants';
import { WinstonLoggerService } from './logger.service';
import { createWinstonLogger } from './logger.util';

@Module({})
export class LoggerModule {
  static forService(serviceName: string): DynamicModule {
    return {
      module: LoggerModule,
      providers: [
        {
          provide: LOGGER_CONTEXT,
          useValue: serviceName,
        },
        {
          provide: WinstonLoggerService,
          useFactory: () => {
            return new WinstonLoggerService(
              createWinstonLogger(serviceName),
              serviceName,
            );
          },
        },
      ],
      exports: [WinstonLoggerService],
    };
  }
}
