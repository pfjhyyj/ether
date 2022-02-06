import {
  HttpException,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';
import { User } from 'src/user/entities/user.interface';
import { UserService } from 'src/user/service/user.service';
import { LoginEmailInput } from '../dto/login-email.input';
import { LoginUsernameInput } from '../dto/login-username.input';
import { RegisterInput } from '../dto/register.input';

@Injectable()
export class AuthService {
  constructor(
    private readonly jwtService: JwtService,
    private userService: UserService,
  ) {}

  generateJWT(user: User): Promise<string> {
    return this.jwtService.signAsync({ user });
  }

  hashPassword(password: string): Promise<string> {
    return bcrypt.hash(password, 13);
  }

  comparePasswords(newPassword: string, passwortHash: string): Promise<any> {
    return bcrypt.compare(newPassword, passwortHash);
  }

  async loginByEmail(loginInput: LoginEmailInput): Promise<string> {
    const user = await this.userService.findOnePasswordByEmail(
      loginInput.email,
    );
    const match: boolean = await this.comparePasswords(
      loginInput.password,
      user.password,
    );
    if (match) {
      const token = this.generateJWT(user);
      return token;
    } else {
      throw new UnauthorizedException();
    }
  }

  async loginByUsername(loginInput: LoginUsernameInput): Promise<string> {
    const user = await this.userService.findOnePasswordByUsername(
      loginInput.username,
    );
    const match: boolean = await this.comparePasswords(
      loginInput.password,
      user.password,
    );
    if (match) {
      const token = this.generateJWT(user);
      return token;
    } else {
      throw new UnauthorizedException();
    }
  }

  async registerByEmail(registerInput: RegisterInput): Promise<User> {
    const oldUser = await this.userService.findOnePasswordByEmail(
      registerInput.email,
    );
    if (oldUser.id) {
      throw new HttpException('Existed Email', 401);
    } else {
      const newUser: User = {
        ...registerInput,
        password: await this.hashPassword(registerInput.password),
      };
      return this.userService.create(newUser);
    }
  }
}
