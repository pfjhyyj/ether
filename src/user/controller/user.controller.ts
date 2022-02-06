import {
  Body,
  Controller,
  Delete,
  Param,
  Put,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/guard/jwt-guard';
import { UpdateUserInput } from '../dto/update-user.input';
import { User } from '../entities/user.interface';
import { UserService } from '../service/user.service';

@Controller('user')
@UseGuards(JwtAuthGuard)
export class UserController {
  constructor(private userService: UserService) {}

  @Put(':id')
  updateOne(
    @Param('id') id: string,
    @Body() user: UpdateUserInput,
  ): Promise<User> {
    return this.userService.update(Number(id), user);
  }

  @Delete(':id')
  deleteOne(@Param('id') id: string): Promise<any> {
    return this.userService.remove(Number(id));
  }
}
