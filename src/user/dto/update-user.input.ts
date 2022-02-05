import { CreateUserInput } from './create-user.input';
import { InputType, Field, Int, PartialType } from '@nestjs/graphql';

@InputType()
export class UpdateUserInput extends PartialType(CreateUserInput) {
  @Field(() => Int, { description: 'User ID' })
  id: number;

  @Field({ description: 'Username, unique' })
  username: string;

  @Field({ nullable: true, description: 'User First name' })
  firstName?: string;

  @Field({ nullable: true, description: 'User Last name' })
  lastName?: string;

  @Field({ description: 'User Email address' })
  email: string;

  @Field({ nullable: true, description: 'User password, meaningless to show' })
  password?: string;

  @Field({ nullable: true, description: 'User Mobile number' })
  mobile?: string;

  @Field({ nullable: true, description: 'User Tel number' })
  tel?: string;

  @Field({ nullable: true, description: 'User Language code' })
  lang?: string;
}
