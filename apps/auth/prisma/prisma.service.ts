import { Injectable } from '@nestjs/common';
import {PrismaClient} from "@prisma/client";
import {ConfigService} from "@nestjs/config";

@Injectable()
export class Prisma extends PrismaClient{
    constructor(private configService: ConfigService) {
        super({
            datasources: {
                db :{
                    url: configService.get("DATABASE_URL")
                },
            },
        });
    }
}
