import { Module } from '@nestjs/common';
import { AuthService } from './service/auth.service';
import { AuthController } from './controller/auth.controller';
import { UserModule } from 'src/user/user.module';
import { JwtAuthGuard } from './guard/jwt-guard';
import { JwtStrategy } from './guard/jwt-strategy';
import { GqlAuthGuard } from './guard/gql-guard';

@Module({
  imports: [UserModule],
  providers: [AuthService, JwtAuthGuard, JwtStrategy, GqlAuthGuard],
  controllers: [AuthController],
})
export class AuthModule {}
