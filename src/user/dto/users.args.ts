import { Field, ArgsType, Int } from '@nestjs/graphql';
import { Max, Min } from 'class-validator';

@ArgsType()
export class UsersArgs {
  @Field({ description: 'Username, unique' })
  username?: string;

  @Field({ nullable: true, description: 'User First name' })
  firstName?: string;

  @Field({ nullable: true, description: 'User Last name' })
  lastName?: string;

  @Field({ description: 'User Email address' })
  email?: string;

  @Field({ nullable: true, description: 'User Mobile number' })
  mobile?: string;

  @Field({ nullable: true, description: 'User Tel number' })
  tel?: string;

  @Field(() => Int)
  @Min(0)
  skip = 0;

  @Field(() => Int)
  @Min(1)
  @Max(50)
  take = 10;
}
