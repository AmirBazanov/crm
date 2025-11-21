import {
  IsEmail,
  IsJWT,
  IsNotEmpty,
  MinLength,
  ValidateNested,
} from 'class-validator';
import { UserDto } from './user.dto';
import { Type } from 'class-transformer';

export class AuthRegisterDto {
  @IsEmail()
  email: string;
  @IsNotEmpty()
  @MinLength(3)
  username: string;
  @IsNotEmpty()
  password: string;
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
