import { Test, TestingModule } from '@nestjs/testing';
import { GraphqlApiController } from './graphql-api.controller';
import { GraphqlApiService } from './graphql-api.service';

describe('GraphqlApiController', () => {
  let graphqlApiController: GraphqlApiController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [GraphqlApiController],
      providers: [GraphqlApiService],
    }).compile();

    graphqlApiController = app.get<GraphqlApiController>(GraphqlApiController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(graphqlApiController.getHello()).toBe('Hello World!');
    });
  });
});
