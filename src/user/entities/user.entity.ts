import { Field, Int, ObjectType } from '@nestjs/graphql';
import {
  BeforeInsert,
  Column,
  CreateDateColumn,
  Entity,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';

@ObjectType()
@Entity()
export class UserEntity {
  @Field(() => Int, { description: 'User ID' })
  @PrimaryGeneratedColumn()
  id: number;

  @Field({ description: 'Username, unique' })
  @Column({ type: 'varchar', length: 100, unique: true })
  username: string;

  @Field({ nullable: true, description: 'User First name' })
  @Column({ type: 'varchar', length: 100, nullable: true })
  firstName: string;

  @Field({ nullable: true, description: 'User Last name' })
  @Column({ type: 'varchar', length: 100, nullable: true })
  lastName: string;

  @Field({ description: 'User Email address' })
  @Column({ type: 'varchar', length: 100, unique: true })
  email: string;

  @Field({ nullable: true, description: 'User password, meaningless to show' })
  @Column({ type: 'varchar', length: 100, select: false })
  password: string;

  @Field({ nullable: true, description: 'User Mobile number' })
  @Column({ type: 'varchar', length: 100, nullable: true })
  mobile: string;

  @Field({ nullable: true, description: 'User Tel number' })
  @Column({ type: 'varchar', length: 100, nullable: true })
  tel: string;

  @Field({ nullable: true, description: 'User Language code' })
  @Column({ type: 'varchar', length: 50, nullable: true })
  lang: string;

  @Field({ nullable: true, description: 'User Register date' })
  @CreateDateColumn()
  created: Date;

  @Field({ nullable: true, description: 'User update date' })
  @UpdateDateColumn()
  updated: Date;

  @BeforeInsert()
  emailToLowerCase() {
    this.email = this.email.toLowerCase();
  }
}
