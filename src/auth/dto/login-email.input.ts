import { IsEmail, IsNotEmpty, IsString, Max } from 'class-validator';

export class LoginEmailInput {
  @IsEmail()
  @IsNotEmpty()
  @Max(100)
  email: string;

  @IsString()
  @IsNotEmpty()
  @Max(100)
  password: string;
}
