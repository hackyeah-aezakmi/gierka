import { Route } from '@angular/router';
import {LandingComponent} from "@gierka/landing";
import { JoinComponent } from '@gierka/join';
import { RoomComponent } from '@gierka/room';
import { GameComponent } from '@gierka/game';

export const appRoutes: Route[] = [
    {
        path: '',
        component: LandingComponent
    },
    {
        path: 'join/:id',
        component: JoinComponent
    },
    {
        path: 'room',
        component: RoomComponent
    },
    {
        path: 'game/:id',
        component: GameComponent,
    },
];
