import { inject, Injectable, signal } from '@angular/core';
import { Game } from './model/game.class';
import { v4 as uuid } from 'uuid';
import { User } from './model/user.class';
import { RestService } from './rest.service';

@Injectable({
  providedIn: 'root'
})
export class StateService {
  private restService = inject(RestService);

  uid: string = this.getUid();

  game = signal<Game>();
  users = signal<Record<string, User>>({})

  private getUid(): string {
    try {
      const uid = localStorage.getItem('uid');

      if (uid) {
        return uid;
      } else {
        return this.initUser();
      }
    } catch (error) {
      return this.initUser();
    }
  }

  join(host: boolean) {
    this.addUser(host ? User.createHost(this.uid) : User.create(this.uid))
  }

  private initUser() {
    const uid = uuid();

    localStorage.setItem('uid', uid);

    return uid;
  }

  createNewGame() {
    this.game.set(Game.new(uuid()))

    this.saveGame();
  }

  private saveGame() {
    this.restService.saveGame(this.game());
  }

  addUser(user: User = User.create(uuid())) {
    this.users.update((users: Record<string, User>) => ({
      ...users,
      [user.id]: user,
    }));

    return user;
  }

  createNewUser(nick: string) {
    const user = User.create(uuid());

    user.nick = nick;

    return user;
  }
}
