import {
  IsEmail,
  IsJWT,
  IsNotEmpty,
  MinLength,
  ValidateNested,
  IsEnum,
  IsOptional,
} from 'class-validator';
import { Type } from 'class-transformer';
import { UserDto, Country } from './user.dto';

export class AuthRegisterDto {
  @IsEmail()
  email: string;

  @IsNotEmpty()
  password: string;

  @IsNotEmpty()
  firstname: string;

  @IsNotEmpty()
  lastname: string;

  @IsNotEmpty()
  nickname: string;

  @IsEnum(Country)
  country: Country;
}

export class AuthDtoResp {
  @IsJWT()
  accessToken: string;

  @IsJWT()
  refreshToken: string;

  @ValidateNested()
  @Type(() => UserDto)
  user: UserDto;
}

export class AuthLoginDto {
  @IsEmail()
  email: string;

  @IsNotEmpty()
  password: string;
}

export class AuthRefreshDto {
  @IsJWT()
  refreshToken: string;
}

export class AuthRefreshDtoResp {
  @IsJWT()
  accessToken: string;

  @IsJWT()
  refreshToken: string;
}

export class AuthLogoutDto {
  @IsJWT()
  refreshToken: string;
}

export class AuthLogoutDtoResp {
  @IsNotEmpty()
  message: string;
}
