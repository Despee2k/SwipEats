import { Routes } from '@angular/router';
import { Landing } from './pages/landing/landing';
import { Login } from './pages/login/login';
import { Signup } from './pages/signup/signup'

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
    }
];
