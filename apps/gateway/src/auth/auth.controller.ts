import {Body, Controller, Post} from '@nestjs/common';
import {
    AuthDtoResp,
    AuthLoginDto, AuthLogoutDto,
    AuthRefreshDto,
    AuthRefreshDtoResp,
    AuthRegisterDto
} from "../../../../libs/dto/auth.dto";
import {AuthService} from "./auth.service";

@Controller('auth')
export class AuthController {
    constructor(private readonly authService: AuthService) {
    }
    @Post('login')
    async login(@Body() body: AuthLoginDto){

        return AuthDtoResp
    }

    @Post("register")
    async register(@Body() body: AuthRegisterDto){

        return await this.authService.register(body)
    }

    @Post("refresh")
    async refresh(@Body() body: AuthRefreshDto){

        return AuthRefreshDtoResp
    }

    @Post("logout")
    async logout(@Body() body: AuthLogoutDto){

        return AuthRefreshDtoResp
    }

}
