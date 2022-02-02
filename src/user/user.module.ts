import { Module } from '@nestjs/common';
import { UserService } from './service/user.service';
import { UserResolver } from './resolver/user.resolver';
import { TypeOrmModule } from '@nestjs/typeorm';
import { User } from './entities/user.entity';

@Module({
  imports: [
    TypeOrmModule.forFeature([User])
  ],
  providers: [UserResolver, UserService]
})
export class UserModule {}
