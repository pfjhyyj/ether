import { IsEmail, IsNotEmpty, IsString, Max } from 'class-validator';

export class UpdateUserInput {
  @IsString()
  @IsNotEmpty()
  @Max(100)
  username: string;

  @IsString()
  @Max(100)
  firstName?: string;

  @IsString()
  @Max(100)
  lastName?: string;

  @IsEmail()
  @IsNotEmpty()
  @Max(100)
  email: string;

  @IsString()
  @IsNotEmpty()
  @Max(100)
  password?: string;

  @IsString()
  @Max(100)
  mobile?: string;

  @IsString()
  @Max(100)
  tel?: string;

  @IsString()
  @Max(100)
  lang?: string;
}
