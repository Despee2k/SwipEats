import { Routes } from '@angular/router';
import { Landing } from './pages/landing/landing';
import { Login } from './pages/login/login';
import { Signup } from './pages/signup/signup'
import { Lobby } from './pages/lobby/lobby';
import { Group } from './pages/group/group';
import { GroupInterface } from './pages/group-interface/group-interface';

export const routes: Routes = [
    {
        path: '',
        component: Landing
    },
    {
        path: 'login',
        component: Login
    },
    {
        path: 'signup',
        component: Signup
    },
    {
        path: 'lobby',
        component: Lobby
    },
    {
        path: 'group',
        component: Group
    },
    {
        path: 'group/:groupCode',
        component: GroupInterface
    }
];
