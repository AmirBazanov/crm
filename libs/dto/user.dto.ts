import { IsArray, IsEmail, IsEnum, IsNumber, IsOptional, IsString, MinLength, ValidateNested } from 'class-validator';
import { Type } from 'class-transformer';

export enum Country {
  COUNTRY_UNSPECIFIED = 0,
  COUNTRY_EN = 1,
  COUNTRY_IT = 2,
  COUNTRY_FR = 3,
  COUNTRY_DE = 4,
  COUNTRY_RU = 5,
}

export class UserDto {
  @IsNumber()
  id: number;

  @IsString()
  @MinLength(3)
  firstname: string;

  @IsString()
  @MinLength(3)
  lastname: string;

  @IsString()
  @MinLength(3)
  nickname: string;

  @IsEmail()
  email: string;

  @IsEnum(Country)
  country: Country;

  @IsString()
  createdAt: string;

  @IsString()
  updatedAt: string;
}

export class CreateUserDto {
  @IsString()
  @MinLength(3)
  firstname: string;

  @IsString()
  @MinLength(3)
  lastname: string;

  @IsString()
  @MinLength(3)
  nickname: string;

  @IsEmail()
  email: string;

  @IsString()
  @MinLength(6)
  password: string;

  @IsEnum(Country)
  country: Country;
}

export class CreateUserDtoResp {
  @IsNumber()
  id: number;
}

export class UpdateUserDto {
  @IsNumber()
  id: number;

  @IsString()
  @MinLength(3)
  @IsOptional()
  firstname?: string;

  @IsString()
  @MinLength(3)
  @IsOptional()
  lastname?: string;

  @IsString()
  @MinLength(3)
  @IsOptional()
  nickname?: string;

  @IsEmail()
  @IsOptional()
  email?: string;

  @IsEnum(Country)
  @IsOptional()
  country?: Country;
}

export class UpdateUserDtoResp {
  @ValidateNested()
  @Type(() => UserDto)
  user: UserDto;
}

export class SearchUserDto {
  @IsEnum(Country)
  @IsOptional()
  country?: Country;

  @IsString()
  @MinLength(3)
  @IsOptional()
  nickname?: string;

  @IsEmail()
  @IsOptional()
  email?: string;

  @IsString()
  @MinLength(3)
  @IsOptional()
  firstname?: string;

  @IsString()
  @MinLength(3)
  @IsOptional()
  lastname?: string;
}

export class SearchUserDtoResp {
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => UserDto)
  users: UserDto[];
}

export class GetUserDto {
  @IsNumber()
  id: number;
}

export class GetUserDtoResp {
  @ValidateNested()
  @Type(() => UserDto)
  user: UserDto;
}

export class DeleteUserDto {
  @IsNumber()
  id: number;
}

export class DeleteUserDtoResp {
  @IsNumber()
  id: number;
}

