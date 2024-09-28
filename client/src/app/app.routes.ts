import { Route } from '@angular/router';
import {LandingComponent} from "@gierka/landing";
import { JoinComponent } from '@gierka/join';
import { RoomComponent } from '@gierka/room';

export const appRoutes: Route[] = [
    {
        path: '',
        component: LandingComponent
    },
    {
        path: 'join',
        component: JoinComponent
    },
    {
        path: 'room',
        component: RoomComponent
    }
];
