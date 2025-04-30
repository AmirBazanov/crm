import { utilities as nestWinstonModuleUtilities } from 'nest-winston';
import * as winston from 'winston';

export function createWinstonLogger(serviceName: string) {
  return winston.createLogger({
    level: 'info',
    format: winston.format.combine(
      winston.format.timestamp(),
      winston.format.ms(),
      winston.format.printf(({ level, message, timestamp }) => {
        return `${timestamp} [${level.toUpperCase()}] ${message}`;
      }),
    ),
    transports: [
      new winston.transports.File({
        filename: `logs/${serviceName}.log`,
      }),
      new winston.transports.Console({
        format: nestWinstonModuleUtilities.format.nestLike(serviceName, {
          prettyPrint: true,
        }),
      }),
    ],
  });
}
