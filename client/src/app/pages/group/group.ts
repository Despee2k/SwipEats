import { Component } from '@angular/core';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';

@Component({
  selector: 'app-group',
  imports: [NavigationBar],
  templateUrl: './group.html',
  styleUrl: './group.css'
})
export class Group {
  activeModal: 'create' | 'join' | 'success' | null = 'create';
}
