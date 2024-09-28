import { MiniGame } from './mini-game.enum';

export class User {
  host: boolean = false;
  id!: string;
  nick!: string;
  selectedMiniGame?: MiniGame;

  static create(id: string, nick: string, isHost = false) {
    const user = new User();

    user.id = id;
    user.nick = nick;
    user.host = isHost;
    return user;
  }
}
