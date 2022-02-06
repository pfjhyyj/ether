import { Body, Controller, Post } from '@nestjs/common';
import { User } from 'src/user/entities/user.interface';
import { LoginEmailInput } from '../dto/login-email.input';
import { RegisterInput } from '../dto/register.input';
import { AuthService } from '../service/auth.service';

@Controller('auth')
export class AuthController {
  constructor(private authService: AuthService) {}

  @Post('login')
  loginByEmail(@Body() loginInput: LoginEmailInput): Promise<string> {
    return this.authService.loginByEmail({
      email: loginInput.email,
      password: loginInput.password,
    });
  }

  @Post('register')
  registerByEmail(@Body() registerInput: RegisterInput): Promise<User> {
    return this.authService.registerByEmail(registerInput);
  }
}
