import { Component } from '@angular/core';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-lobby',
  imports: [NavigationBar, RouterLink],
  templateUrl: './lobby.html',
  styleUrl: './lobby.css'
})
export class Lobby {

}
