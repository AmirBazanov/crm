import { Injectable } from '@nestjs/common';

@Injectable()
export class GraphqlApiService {
  getHello(): string {
    return 'Hello World!';
  }
}
