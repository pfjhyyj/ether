import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
// import { CreateUserInput } from '../dto/create-user.input';
import { UpdateUserInput } from '../dto/update-user.input';
import { UsersArgs } from '../dto/users.args';
import { UserEntity } from '../entities/user.entity';
import { User } from '../entities/user.interface';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(UserEntity)
    private readonly userRepository: Repository<UserEntity>,
  ) {}

  create(createUserInput: User): Promise<User> {
    return this.userRepository.save(createUserInput);
  }

  findAll(usersArgs: UsersArgs): Promise<UserEntity[]> {
    return this.userRepository.find(usersArgs);
  }

  findOne(id: number): Promise<UserEntity> {
    return this.userRepository.findOne({ id });
  }

  findOnePasswordByUsername(username: string): Promise<UserEntity> {
    return this.userRepository.findOne(
      { username },
      { select: ['id', 'username', 'password', 'email'] },
    );
  }

  findOnePasswordByEmail(email: string): Promise<UserEntity> {
    return this.userRepository.findOne(
      { email },
      { select: ['id', 'username', 'password', 'email'] },
    );
  }

  update(id: number, updateUserInput: UpdateUserInput): Promise<UserEntity> {
    const updatedUser: User = {
      ...updateUserInput,
    };
    return this.userRepository.update(id, updatedUser).then(() => {
      return this.findOne(id);
    });
  }

  remove(id: number): Promise<any> {
    return this.userRepository.delete(id);
  }
}
