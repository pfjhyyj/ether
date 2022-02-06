import { IsNotEmpty, IsString, Max } from 'class-validator';

export class LoginUsernameInput {
  @IsString()
  @IsNotEmpty()
  @Max(100)
  username: string;

  @IsString()
  @IsNotEmpty()
  @Max(100)
  password: string;
}
