import { Field, ArgsType } from '@nestjs/graphql';

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
}
