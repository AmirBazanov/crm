import { Controller } from '@nestjs/common';
import {GrpcMethod} from "@nestjs/microservices";
import {
    AuthRegisterRequest,
    AuthServiceController,
} from "../../../proto/generated/auth";


@Controller()
// @ts-ignore
export class AuthController implements AuthServiceController{

    @GrpcMethod('AuthService', 'Register')
    async Register(data: AuthRegisterRequest){
        console.log("GOT IT!")
    }
}
