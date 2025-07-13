import { Component, inject} from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-group-interface',
  imports: [CommonModule],
  templateUrl: './group-interface.html',
  styleUrl: './group-interface.css'
})
export class GroupInterface {
  groupCode: string;
  constructor(private router: Router) {
    const urlParts = this.router.url.split('/');
    this.groupCode = urlParts[urlParts.length - 1]; // gets the last part of the URL which is the fucking code
  }
}