import {Inject, Injectable, OnModuleInit} from '@nestjs/common';
import {AuthRegisterDto} from "../../../../libs/dto/auth.dto";
import {ClientGrpc} from "@nestjs/microservices";
import {AuthServiceClient} from "../../../../proto/generated/auth";



@Injectable()
export class AuthService implements OnModuleInit{
    private authService: AuthServiceClient
    constructor(@Inject('AUTH_PACKAGE') private client: ClientGrpc) {}
    onModuleInit(){
        this.authService = this.client.getService<AuthServiceClient>("AuthService")
    }
    async register(dto: AuthRegisterDto){
        this.authService.register(dto).subscribe((resp)=>{return resp})
    }
}
