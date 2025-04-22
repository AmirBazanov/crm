import { Controller, Get } from '@nestjs/common';
import { GraphqlApiService } from './graphql-api.service';

@Controller()
export class GraphqlApiController {
  constructor(private readonly graphqlApiService: GraphqlApiService) {}

  @Get()
  getHello(): string {
    return this.graphqlApiService.getHello();
  }
}
