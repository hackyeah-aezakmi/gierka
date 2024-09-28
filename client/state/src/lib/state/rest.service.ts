import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Game } from './model/game.class';

@Injectable({
  providedIn: 'root'
})
export class RestService {
  http = inject(HttpClient);

  saveGame(data: Game) {
    this.http.put(`/api/game/state/${data.id}`, data)
  }
}
