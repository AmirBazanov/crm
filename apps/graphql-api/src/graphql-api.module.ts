import { Module } from '@nestjs/common';
import { GraphqlApiController } from './graphql-api.controller';
import { GraphqlApiService } from './graphql-api.service';

@Module({
  imports: [],
  controllers: [GraphqlApiController],
  providers: [GraphqlApiService],
})
export class GraphqlApiModule {}
