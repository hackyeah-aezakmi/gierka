import { inject, Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Game } from './model/game.class';
import { User } from './model/user.class';
import { environment } from '../../../../src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RestService {
  http = inject(HttpClient);

  apiUrl = environment.apiUrl;

  newGame(data: Game) {
    this.http.put(`${this.apiUrl}/api/game/state/${data.id}`, data, {
    });
  }

  updateGame(game: Game) {
    const headers = new HttpHeaders({
      'x-game-id': game.id,
    });
    this.http.patch(`${this.apiUrl}/api/game/state`, game, { headers });
  }

  updateUser(user: User, userId: string, gameId: string) {
    const headers = new HttpHeaders({
      'x-user-id': userId,
      'x-game-id': gameId,
    });
    this.http.patch(`${this.apiUrl}/api/user/state`, user, { headers });
  }

  newUser(user: User, gameId: string) {
    this.http.put(`${this.apiUrl}/api/user/state`, {
      id: user.id,
      gameId,
      data: user,
    });
  }
}
