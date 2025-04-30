import { IsEmail, IsString, IsUUID, MinLength } from 'class-validator';

export class UserGetDto {
  @IsUUID()
  id: string;
  @IsEmail()
  email: string;
  @IsString()
  @MinLength(3)
  username: string;
}
