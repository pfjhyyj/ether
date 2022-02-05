import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateUserInput } from '../dto/create-user.input';
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

  create(createUserInput: CreateUserInput): Promise<User> {
    const newUser: User = {
      ...createUserInput,
    };
    return this.userRepository.save(newUser);
  }

  findAll(usersArgs: UsersArgs): Promise<UserEntity[]> {
    return this.userRepository.find(usersArgs);
  }

  findOne(id: number): Promise<UserEntity> {
    return this.userRepository.findOne({ id });
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
