export class User {
  host: boolean = false;
  id!: string;
  nick!: string;

  static createHost(id: string) {
    const user = new User();

    user.id = id;
    user.host = true;
    return user;
  }

  static create(id: string) {
    const user = new User();

    user.id = id;
    return user;
  }
}
