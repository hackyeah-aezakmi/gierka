import { effect, inject, Injectable, signal } from '@angular/core';
import { Game } from './model/game.class';
import { v4 as uuid } from 'uuid';
import { User } from './model/user.class';
import { RestService } from './rest.service';
import { Router } from '@angular/router';
import { MiniGame } from './model/mini-game.enum';

@Injectable({
  providedIn: 'root'
})
export class StateService {
  private restService = inject(RestService);
  private router = inject(Router);

  uid: string = this.getUid();

  game = signal<Game | null>(null);
  users = signal<Record<string, User>>({});

  saveEffect = effect(() => {
    console.log('save game');
    this.saveGame(this.game());

    if (this.game()?.started) {
      this.router.navigate([`/game/${this.game()?.id}`]);
    }
  });

  get user(): User {
    return this.users()[this.uid];
  }

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

  private initUser() {
    const uid = uuid();

    localStorage.setItem('uid', uid);

    return uid;
  }

  createNewGame() {
    this.game.set(Game.new(uuid()));

    this.saveGame(this.game());
    this.router.navigate([`/join/${this.game()?.id}`]);
  }

  private saveGame(game: Game | null = this.game()) {
    if (game) {
      this.restService.newGame(game);

      localStorage.setItem('data-game', JSON.stringify(game));
    }
  }

  private updateUser(data?: Partial<User>, user = this.users()[this.uid]) {
    if (data) {
      user = {
        ...user,
        ...data
      };
    }
    if (user) {
      this.restService.updateUser(user, this.uid, this.game()!.id);

      localStorage.setItem('data-users', JSON.stringify(this.users()));
    }
  }

  private updateGame(game: Partial<Game>) {
    if (game) {
      game = {
        ...this.game(),
        ...game
      };

      this.restService.updateGame(game as Game);

      localStorage.setItem('data-game', JSON.stringify(game));

      // todo remove if WS
      this.game.set(game as Game);
    }
  }

  addUser(user: User) {
    this.users.update((users: Record<string, User>) => ({
      ...users,
      [user.id]: user
    }));

    this.restService.newUser(user, this.game()!.id);

    localStorage.setItem('data-users', JSON.stringify(this.users()));

    return user;
  }

  createNewUser(nick: string) {
    return User.create(this.uid, nick, this.isUsersEmpty());
  }

  createFakeUser() {
    return User.create(uuid(), 'fake-' + Object.keys(this.users()).length + 1);
  }

  private isUsersEmpty() {
    return Object.keys(this.users()).length === 0;
  }

  start() {
    this.game.update((game: Game | null) => ({
      ...game!,
      started: true
    }));
  }

  checkGameState() {
    if (this.users()[this.uid] && this.game()?.started) {
      this.router.navigate([`/game/${this.game()?.id}`]);
    }
  }

  voteForMiniGame(miniGame: MiniGame) {
    this.updateUser({
      selectedMiniGame: miniGame
    });
  }

  resetVoteForMiniGame() {
    this.updateUser({
      selectedMiniGame: undefined
    });
  }

  async endVoting() {
    const counter = (Object.values(this.users()) as Array<User>).reduce((counter: Array<number>, user: User) => {
      if (user.selectedMiniGame && counter[user.selectedMiniGame] != undefined) {
        counter[user.selectedMiniGame] = counter[user.selectedMiniGame] ? (counter[user.selectedMiniGame] + 1) : 1;
      }
      return counter;
    }, []);
    let max = 0;
    const miniGame = counter.reduce((index, count, currentIndex) => {
      if (count > counter[index]) {
        return currentIndex;
      }
      return index;
    }, 0);

    this.updateGame({
      loading: true,
    });

    const config = await this.getMiniGameConfig(miniGame);

    this.updateGame({
      miniGame: miniGame,
      config: config,
      loading: false,
    });
  }

  private async getMiniGameConfig(miniGame: MiniGame) {
    switch (miniGame) {
      case MiniGame.QUIZ:
        return Promise.resolve({
          question: 'Który król panujący w czasie Wielkiej Rewolucji Francuskiej został zgilotynowany na Placu Rewolucji w Paryżu?',
          answers: [
            'Ludwik XVI',
            'Ludwik XIV',
            'Napoleon Bonaparte',
            'Karol X'
          ],
          correctIndex: 0,
        });
      case MiniGame.COLORS:
        return Promise.resolve({});
    }
  }

  loadGame(id?: any): boolean {
    try {
      const game = JSON.parse(localStorage.getItem('data-game')!);
      const users = JSON.parse(localStorage.getItem('data-users')!);

      if (!game || !users) {
        return false;
      }

      this.users.set(users);
      this.game.set(game);

      return true;
    } catch {
      return false;
    }
  }
}
