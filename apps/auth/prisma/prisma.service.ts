import {Injectable, OnModuleInit} from '@nestjs/common';
import {PrismaClient} from "@prisma/client";
import {ConfigService} from "@nestjs/config";
import {PRISMA_ERROR_MESSAGES, PRISMA_MESSAGES} from "../../../libs/constants/error-messages";
import {WinstonLoggerService} from "../../../libs/logger/logger.service";

@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit{
    constructor(private configService: ConfigService, private readonly logger: WinstonLoggerService) {
        super({
            datasources: {
                db :{
                    url: configService.get("DATABASE_URL")
                },
            },
        });
    }
    onModuleInit(): any {
        this.checkDatabaseConnection().then()
    }
    private async checkDatabaseConnection() {
        try {
            // Попробуем выполнить простой запрос для проверки соединения
            await this.$queryRaw`SELECT 1`;
            this.logger.log(PRISMA_MESSAGES.connectionSuccess);
        } catch (error) {
            this.logger.error(PRISMA_ERROR_MESSAGES.CONNECTION_ERROR, error.trace);
            throw new Error(PRISMA_ERROR_MESSAGES.CONNECTION_ERROR);
        }
    }
}
