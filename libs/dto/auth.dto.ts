import {IsEmail, IsHash, IsJWT, IsNotEmpty, MinLength, ValidateNested} from "class-validator";
import {UserGetDto} from "./user.dto";
import {Type} from "class-transformer";

export class AuthRegisterDto{
    @IsEmail()
    email: string
    @IsNotEmpty()
    @MinLength(3)
    username: string
    @IsNotEmpty()
    password: string
}

export class AuthDtoResp{
    @IsJWT()
    accessToken: string
    @IsJWT()
    refreshToken: string
    @ValidateNested()
    @Type(()=>UserGetDto)
    user: UserGetDto
}
export class AuthLoginDto{
    @IsEmail()
    email: string
    @IsNotEmpty()
    password: string
}